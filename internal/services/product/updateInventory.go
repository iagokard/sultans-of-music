package product

import (
	"context"
	"errors"
	"fmt"

	"som/internal/database"
	dto "som/internal/dto/product"
	"som/internal/services"
)

var ErrNoInventoryRow = errors.New("inventory row not found or soft‑deleted")

func UpdateInventory(ctx context.Context, req dto.UpdateInventoryRequest) error {
	if req.NewStock < 0 {
		return errors.New("new stock cannot be negative")
	}

	tx, err := services.GetSQLDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	dbQueries := services.GetDatabaseQueries().WithTx(tx)

	inventory, err := dbQueries.GetInventoryByProductID(ctx, int32(req.ProductID))
	if err != nil {
		return fmt.Errorf("product %d: %w", req.ProductID, err)
	}
	currentStock := int(inventory.Stock)
	newStock := req.NewStock

	updateResult, err := dbQueries.SetInventoryStock(ctx, database.SetInventoryStockParams{
		NewStock:  int32(newStock),
		ProductID: int32(req.ProductID),
	})
	if err != nil {
		return fmt.Errorf("updating stock: %w", err)
	}
	rows, _ := updateResult.RowsAffected()
	if rows == 0 {
		return ErrNoInventoryRow
	}

	var amount int
	var logType string

	if newStock > currentStock {
		amount = newStock - currentStock
		logType = services.InventoryLogType.Entry
	} else if newStock < currentStock {
		amount = currentStock - newStock
		logType = services.InventoryLogType.Exit
	} else {
		return tx.Commit()
	}

	_, err = dbQueries.CreateInventoryLog(ctx, database.CreateInventoryLogParams{
		ProductID: int32(req.ProductID),
		Amount:    int32(amount),
		TypeName:  logType,
	})
	if err != nil {
		return fmt.Errorf("logging inventory change: %w", err)
	}

	return tx.Commit()
}

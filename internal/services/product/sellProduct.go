package product

import (
	"context"
	"errors"
	"fmt"

	"som/internal/database"
	dto "som/internal/dto/sale"
	"som/internal/services"
)

var (
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrEmptySale         = errors.New("sale must contain at least one item")
)

func SellProduct(ctx context.Context, req dto.SellProductRequest) (int, error) {
	if len(req.SaleItems) == 0 {
		return 0, ErrEmptySale
	}

	tx, err := services.GetSQLDB().BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()
	dbQueries := services.GetDatabaseQueries().WithTx(tx)

	saleResult, err := dbQueries.CreateSale(ctx)
	if err != nil {
		return 0, err
	}

	saleID, err := saleResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	for index, item := range req.SaleItems {
		inventory, err := dbQueries.GetInventoryByProductID(ctx, int32(item.ProductsID))
		if err != nil {
			return 0, fmt.Errorf("item %d (product %d): %w", index, item.ProductsID, err)
		}

		if inventory.Stock < int32(item.SaleAmount) {
			return 0, fmt.Errorf(
				"item %d (product %d): %w",
				index, item.ProductsID, ErrInsufficientStock,
			)
		}

		updateResult, err := dbQueries.UpdateInventoryStock(
			ctx,
			database.UpdateInventoryStockParams{
				Amount:    int32(item.SaleAmount),
				ProductID: int32(item.ProductsID),
			},
		)

		if err != nil {
			return 0, fmt.Errorf("item %d (product %d): %w", index, item.ProductsID, err)
		}

		rows, _ := updateResult.RowsAffected()
		if rows == 0 {
			return 0, fmt.Errorf(
				"item %d (product %d): %w",
				index, item.ProductsID, ErrInsufficientStock,
			)
		}

		_, err = dbQueries.CreateSaleItem(ctx, database.CreateSaleItemParams{
			SaleID:        int32(saleID),
			ProductID:     int32(item.ProductsID),
			ProductAmount: int32(item.SaleAmount),
		})
		if err != nil {
			return 0, fmt.Errorf("item %d (product %d): %w", index, item.ProductsID, err)
		}

		_, err = dbQueries.CreateInventoryLog(ctx, database.CreateInventoryLogParams{
			ProductID: int32(item.ProductsID),
			Amount:    int32(item.SaleAmount),
			TypeName:  services.InventoryLogType.Exit,
		})

		if err != nil {
			return 0, fmt.Errorf("item %d (product %d): %w", index, item.ProductsID, err)
		}
	}

	return int(saleID), tx.Commit()
}

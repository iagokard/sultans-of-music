package product

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"som/internal/database"
	dto "som/internal/dto/product"
	"som/internal/services"
)

func CreateProduct(ctx context.Context, req dto.CreateProductRequest) error {
	tx, err := services.GetSQLDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	dbQueries := services.GetDatabaseQueries().WithTx(tx)

	var artistID int32
	artist, err := dbQueries.GetArtistByAPIID(ctx, req.ArtistAPIID)
	if errors.Is(err, sql.ErrNoRows) {
		params := database.CreateArtistParams{
			Name:  req.ArtistName,
			ApiID: req.ArtistAPIID,
		}

		result, err := dbQueries.CreateArtist(ctx, params)
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		artistID = int32(id)
	} else if err != nil {
		return err
	} else {
		artistID = artist.ID
	}

	cover := sql.NullString{
		String: req.Cover,
		Valid:  true,
	}

	dateLayout := "2006-01-02"
	date, err := time.Parse(dateLayout, req.ReleaseDate)
	if err != nil {
		return err
	}

	releaseDate := sql.NullTime{
		Time:  date,
		Valid: true,
	}

	productParams := database.CreateProductParams{
		ApiID:       req.APIID,
		TypeID:      int32(req.TypeID),
		ReleaseDate: releaseDate,
		Title:       req.Title,
		ArtistID:    artistID,
		Price:       int32(req.Price),
		Cover:       cover,
	}

	result, err := dbQueries.CreateProduct(ctx, productParams)
	if err != nil {
		return err
	}

	productID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	inventoryParams := database.CreateInventoryItemParams{
		Stock:     int32(req.Stock),
		ProductID: int32(productID),
	}

	_, err = dbQueries.CreateInventoryItem(ctx, inventoryParams)
	if err != nil {
		return err
	}

	logParams := database.CreateInventoryLogParams{
		ProductID: int32(productID),
		Amount:    int32(req.Stock),
		TypeName:  services.InventoryLogType.Entry,
	}

	_, err = dbQueries.CreateInventoryLog(ctx, logParams)
	if err != nil {
		return err
	}

	return tx.Commit()
}

package product

import (
	"context"
	"database/sql"
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

	artist, err := dbQueries.GetArtistByName(ctx, req.Artist)
	if err != nil {
		return err
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
		ArtistID:    artist.ID,
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

	return tx.Commit()
}

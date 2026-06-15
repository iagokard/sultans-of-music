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

func UpdateProduct(ctx context.Context, req dto.UpdateProductRequest) error {
	tx, err := services.GetSQLDB().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	dbQueries := services.GetDatabaseQueries().WithTx(tx)

	product, err := dbQueries.GetProductByID(ctx, int32(req.ProductID))
	if err != nil {
		return err
	}

	currentArtist, err := dbQueries.GetArtistByID(ctx, product.ArtistID)
	if err != nil {
		return err
	}

	apiID := product.ApiID
	if req.APIID != "" {
		apiID = req.APIID
	}

	typeID := product.TypeID
	if req.TypeID != 0 {
		typeID = int32(req.TypeID)
	}

	title := product.Title
	if req.Title != "" {
		title = req.Title
	}

	price := product.Price
	if req.Price != 0 {
		price = int32(req.Price)
	}

	cover := product.Cover
	if req.Cover != "" {
		cover = sql.NullString{String: req.Cover, Valid: true}
	}

	releaseDate := product.ReleaseDate
	if req.ReleaseDate != "" {
		date, err := time.Parse("2006-01-02", req.ReleaseDate)
		if err != nil {
			return err
		}
		releaseDate = sql.NullTime{Time: date, Valid: true}
	}

	artistID := product.ArtistID
	if req.ArtistAPIID != "" && req.ArtistAPIID != currentArtist.ApiID {
		newArtist, err := dbQueries.GetArtistByAPIID(ctx, req.ArtistAPIID)
		if errors.Is(err, sql.ErrNoRows) {
			if req.ArtistName == "" {
				return errors.New("artist name required when creating a new artist")
			}
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
			artistID = newArtist.ID
		}
	}

	params := database.UpdateProductParams{
		ApiID:       apiID,
		TypeID:      typeID,
		ReleaseDate: releaseDate,
		Title:       title,
		ArtistID:    artistID,
		Price:       price,
		Cover:       cover,
		ProductID:   int32(req.ProductID),
	}

	result, err := dbQueries.UpdateProduct(ctx, params)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no product updated (possibly soft‑deleted)")
	}

	return tx.Commit()
}

package product

import (
	"context"

	dto "som/internal/dto/product"
	"som/internal/services"
)

func GetProductByID(
	ctx context.Context,
	productID int,
) (dto.GetProductResponse, error) {
	dbQueries := services.GetDatabaseQueries()
	var productResponse dto.GetProductResponse

	product, err := dbQueries.GetProductByID(ctx, int32(productID))
	if err != nil {
		return productResponse, err
	}

	artist, err := dbQueries.GetArtistByID(ctx, product.ArtistID)
	if err != nil {
		return productResponse, err
	}

	productType, err := dbQueries.GetProductTypeByID(ctx, product.TypeID)
	if err != nil {
		return productResponse, err
	}

	inventory, err := dbQueries.GetInventoryByProductID(ctx, product.ID)
	if err != nil {
		return productResponse, err
	}

	productResponse = dto.GetProductResponse{
		ID:          int(product.ID),
		APIID:       product.ApiID,
		TypeID:      int(product.TypeID),
		Type:        productType.TypeName,
		Title:       product.Title,
		ArtistID:    int(product.ArtistID),
		Artist:      artist.Name,
		ReleaseDate: product.ReleaseDate.Time.String(),
		Price:       int(product.Price),
		Cover:       product.Cover.String,
		Stock:       int(inventory.Stock),
	}

	return productResponse, nil
}

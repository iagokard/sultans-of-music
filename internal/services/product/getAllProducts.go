package product

import (
	"context"

	dto "som/internal/dto/product"
	"som/internal/services"
)

func GetAllProducts(
	ctx context.Context,
) (dto.GetProductListResponse, error) {
	dbQueries := services.GetDatabaseQueries()
	var listResponse dto.GetProductListResponse

	productList, err := dbQueries.GetAllProducts(ctx)
	if err != nil {
		return listResponse, err
	}

	listResponse = dto.GetProductListResponse{
		ProductList: []dto.GetProductResponse{},
	}

	for _, product := range productList {
		artist, err := dbQueries.GetArtistByID(ctx, product.ArtistID)
		if err != nil {
			return listResponse, err
		}

		productType, err := dbQueries.GetProductTypeByID(ctx, product.TypeID)
		if err != nil {
			return listResponse, err
		}

		inventory, err := dbQueries.GetInventoryByProductID(ctx, product.ID)
		if err != nil {
			return listResponse, err
		}

		res := dto.GetProductResponse{
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

		listResponse.ProductList = append(listResponse.ProductList, res)
	}

	return listResponse, nil
}

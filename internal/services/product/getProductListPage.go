package product

import (
	"context"

	"som/internal/database"
	dto "som/internal/dto/product"
	"som/internal/services"
)

func GetProductListPage(
	ctx context.Context,
	req dto.GetProductListPageRequest,
) (dto.GetProductListPageResponse, error) {
	dbQueries := services.GetDatabaseQueries()
	var listResponse dto.GetProductListPageResponse

	limit := max(req.Page, 1)
	offset := (limit - 1) * limit

	params := database.GetProductsListPageParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	productList, err := dbQueries.GetProductsListPage(ctx, params)
	if err != nil {
		return listResponse, err
	}

	listResponse = dto.GetProductListPageResponse{
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

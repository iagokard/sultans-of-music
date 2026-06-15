package sale

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	dto "som/internal/dto/sale"
	"som/internal/services"
)

func GetSaleByID(
	ctx context.Context,
	saleID int,
) (dto.GetSaleResponse, error) {
	dbQueries := services.GetDatabaseQueries()
	var response dto.GetSaleResponse

	sale, err := dbQueries.GetSaleByID(ctx, int32(saleID))
	if errors.Is(err, sql.ErrNoRows) {
		return response, fmt.Errorf("sale not found")
	}
	if err != nil {
		return response, err
	}

	items, err := dbQueries.GetSaleItemsBySaleID(ctx, int32(saleID))
	if err != nil {
		return response, err
	}

	response.ID = int(sale.ID)
	response.DateTime = sale.DateTime.Time
	response.Items = make([]dto.SaleItemInfo, 0, len(items))

	for _, item := range items {
		response.Items = append(response.Items, dto.SaleItemInfo{
			ProductID:    int(item.ProductID),
			ProductTitle: item.ProductTitle,
			ArtistID:     int(item.ArtistID),
			ArtistName:   item.ArtistName,
			TypeID:       int(item.TypeID),
			TypeName:     item.ProductType,
			Price:        int(item.ProductPrice),
			Cover:        item.ProductCover.String,
			Amount:       int(item.ProductAmount),
			Stock:        int(item.Stock),
		})
	}

	return response, nil
}

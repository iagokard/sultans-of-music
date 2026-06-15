package sale

import (
	"context"

	dto "som/internal/dto/sale"
	"som/internal/services"
)

func GetAllSales(
	ctx context.Context,
) (dto.GetSaleListResponse, error) {
	dbQueries := services.GetDatabaseQueries()
	var listResponse dto.GetSaleListResponse

	sales, err := dbQueries.GetAllSales(ctx)
	if err != nil {
		return listResponse, err
	}

	for _, sale := range sales {
		var response dto.GetSaleResponse
		items, err := dbQueries.GetSaleItemsBySaleID(ctx, sale.ID)
		if err != nil {
			return listResponse, err
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

		listResponse.SaleList = append(listResponse.SaleList, response)
	}

	return listResponse, nil
}

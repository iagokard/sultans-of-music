package dto

import "time"

type SaleItemInfo struct {
	ProductID    int    `json:"product-id"`
	ProductTitle string `json:"product-title"`
	ArtistID     int    `json:"artist-id"`
	ArtistName   string `json:"artist"`
	TypeID       int    `json:"type-id"`
	TypeName     string `json:"type"`
	Price        int    `json:"price"`
	Cover        string `json:"cover"`
	Amount       int    `json:"amount"`
	Stock        int    `json:"stock"`
}

type GetSaleResponse struct {
	ID       int            `json:"id"`
	DateTime time.Time      `json:"date-time"`
	Items    []SaleItemInfo `json:"items"`
}

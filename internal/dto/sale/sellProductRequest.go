package dto

type SellProductRequest struct {
	SaleItems []SaleRequestInfo `json:"sale-items" binding:"required"`
}

type SaleRequestInfo struct {
	ProductsID int `json:"product-id" binding:"required"`
	SaleAmount int `json:"sale-amount" binding:"required"`
}

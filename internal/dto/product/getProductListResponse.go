package dto

type GetProductListResponse struct {
	ProductList []GetProductResponse `json:"product-list" binding:"required"`
}

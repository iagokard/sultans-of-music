package dto

type GetProductListPageResponse struct {
	ProductList []GetProductResponse `json:"product-list" binding:"required"`
}

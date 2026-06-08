package dto

type GetProductListPageRequest struct {
	Amount int `json:"amount" binding:"required"`
	Page   int `json:"page" binding:"required"`
}

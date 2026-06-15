package dto

type UpdateInventoryRequest struct {
	ProductID int `json:"product-id"`
	NewStock  int `json:"new-stock"`
}

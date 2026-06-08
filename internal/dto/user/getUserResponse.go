package dto

type GetUserResponse struct {
	ID    int    `json:"id" binding:"required"`
	Email string `json:"email" binding:"required"`
}

package dto

type LoginUserResponse struct {
	JWTToken string `json:"jwt-token" binding:"required"`
}

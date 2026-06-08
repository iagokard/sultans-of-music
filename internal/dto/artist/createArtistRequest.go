package dto

type CreateArtistRequest struct {
	Name  string `json:"name" binding:"required"`
	APIID string `json:"api-id" binding:"required"`
}

package dto

type CreateProductRequest struct {
	APIID       string `json:"api-id" binding:"required"`
	TypeID      int    `json:"type-id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	ArtistAPIID string `json:"artist-api-id" binding:"required"`
	ArtistName  string `json:"artist-name" binding:"required"`
	ReleaseDate string `json:"release_date" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Cover       string `json:"cover" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
}

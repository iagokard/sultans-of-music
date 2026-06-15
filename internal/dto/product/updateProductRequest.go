package dto

type UpdateProductRequest struct {
	ProductID   int    `json:"product-id"`
	APIID       string `json:"api-id"`
	TypeID      int    `json:"type-id"`
	Title       string `json:"title"`
	ArtistAPIID string `json:"artist-api-id"`
	ArtistName  string `json:"artist-name"`
	ReleaseDate string `json:"release_date"`
	Price       int    `json:"price"`
	Cover       string `json:"cover"`
}

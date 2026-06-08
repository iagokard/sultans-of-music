package dto

type GetProductResponse struct {
	ID          int    `json:"id" binding:"required"`
	APIID       string `json:"api-id" binding:"required"`
	TypeID      int    `json:"type-id" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Title       string `json:"title" binding:"required"`
	ArtistID    int    `json:"artist-id" binding:"required"`
	Artist      string `json:"artist" binding:"required"`
	ReleaseDate string `json:"release_date" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Cover       string `json:"cover" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
}

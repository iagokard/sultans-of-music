package artist

import (
	"context"

	"som/internal/database"
	dto "som/internal/dto/artist"
	"som/internal/services"
)

func CreateArtist(ctx context.Context, req dto.CreateArtistRequest) error {
	dbQueries := services.GetDatabaseQueries()

	params := database.CreateArtistParams{
		Name:  req.Name,
		ApiID: req.APIID,
	}

	_, err := dbQueries.CreateArtist(ctx, params)
	return err
}

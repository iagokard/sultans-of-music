package user

import (
	"context"

	dto "som/internal/dto/user"
	"som/internal/services"
)

func GetUserByEmail(
	ctx context.Context,
	userEmail string,
) (dto.GetUserResponse, error) {
	dbQueries := services.GetDatabaseQueries()
	var userResponse dto.GetUserResponse

	user, err := dbQueries.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return userResponse, err
	}

	userResponse = dto.GetUserResponse{
		ID:    int(user.ID),
		Email: user.Email,
	}

	return userResponse, nil
}

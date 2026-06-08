package user

import (
	"context"

	dto "som/internal/dto/user"
	"som/internal/services"
)

func GetUserByID(
	ctx context.Context,
	userID int,
) (dto.GetUserResponse, error) {
	dbQueries := services.GetDatabaseQueries()
	var userResponse dto.GetUserResponse

	user, err := dbQueries.GetUserByID(ctx, int32(userID))
	if err != nil {
		return userResponse, err
	}

	userResponse = dto.GetUserResponse{
		ID:    int(user.ID),
		Email: user.Email,
	}

	return userResponse, nil
}

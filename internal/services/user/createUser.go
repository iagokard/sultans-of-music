package user

import (
	"context"

	"som/internal/database"
	dto "som/internal/dto/user"
	"som/internal/services"
	"som/internal/util"
)

func CreateUser(ctx context.Context, req dto.CreateUserRequest) error {
	dbQueries := services.GetDatabaseQueries()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}

	params := database.CreateUserParams{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	_, err = dbQueries.CreateUser(ctx, params)

	return err
}

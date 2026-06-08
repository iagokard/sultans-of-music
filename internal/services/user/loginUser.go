package user

import (
	"context"
	"strconv"

	"som/internal/config"
	dto "som/internal/dto/user"
	"som/internal/services"
	"som/internal/util"
)

func LoginUser(
	ctx context.Context,
	req dto.LoginUserRequest,
) (dto.LoginUserResponse, error) {
	dbQueries := services.GetDatabaseQueries()

	var loginResponse dto.LoginUserResponse
	user, err := dbQueries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return loginResponse, err
	}

	err = util.CheckPassword(user.PasswordHash, req.Password)
	if err != nil {
		return loginResponse, err
	}

	cfg := config.Load()

	userID := strconv.Itoa(int(user.ID))
	token, err := util.GenerateJWT(userID, cfg.JWTSecret)
	if err != nil {
		return loginResponse, err
	}

	loginResponse = dto.LoginUserResponse{
		JWTToken: token,
	}

	return loginResponse, nil
}

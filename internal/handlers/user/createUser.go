package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "som/internal/dto/user"
	"som/internal/handlers"
	services "som/internal/services/user"
	"som/internal/util"
)

// CreateUser godoc
// @summary      Register a new user.
// @tags         user
// @accept       json
// @produce      json
// @security     BearerAuth
// @param        user_info body dto.CreateUserRequest true "User information"
// @success      201 "User created successfully"
// @failure      400 {object} util.ErrorResponse "Invalid request body"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      409 {object} util.ErrorResponse "User already exists"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /user/create [post]
func CreateUser(ginContext *gin.Context) {
	var req dto.CreateUserRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "invalid request",
			Details: err.Error(),
		})
		return
	}

	err := services.CreateUser(ginContext.Request.Context(), req)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.Status(http.StatusCreated)
}

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "som/internal/dto/user"
	"som/internal/handlers"
	services "som/internal/services/user"
	"som/internal/util"
)

// LoginUser godoc
// @summary      Authenticate a user.
// @tags         user
// @accept       json
// @produce      json
// @param        user_login body dto.LoginUserRequest true "User login credentials"
// @success      200 {object} dto.LoginUserResponse "Login successful"
// @failure      400 {object} util.ErrorResponse "Invalid request body"
// @failure      401 {object} util.ErrorResponse "Invalid credentials"
// @failure      404 {object} util.ErrorResponse "User not found"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /user/login [post]
func LoginUser(ginContext *gin.Context) {
	var req dto.LoginUserRequest

	if err := ginContext.ShouldBindJSON(&req); err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "invalid request",
			Details: err.Error(),
		})
		return
	}

	loginResponse, err := services.LoginUser(ginContext.Request.Context(), req)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.JSON(http.StatusOK, loginResponse)
}

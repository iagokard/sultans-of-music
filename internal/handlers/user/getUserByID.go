package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"som/internal/handlers"
	services "som/internal/services/user"
	"som/internal/util"
)

// GetUserByID godoc
// @summary      Get user by ID.
// @tags         user
// @produce      json
// @security     BearerAuth
// @param        id path int true "User ID"
// @success      200 {object} dto.GetUserResponse "User found"
// @failure      400 {object} util.ErrorResponse "Invalid ID supplied"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      404 {object} util.ErrorResponse "User not found"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /user/{id} [get]
func GetUserByID(ginContext *gin.Context) {
	idParam := ginContext.Param("id")

	userID, err := strconv.Atoi(idParam)
	if err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "invalid request",
			Details: "ID should be a valid integer",
		})
		return
	}

	user, err := services.GetUserByID(ginContext.Request.Context(), userID)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.JSON(http.StatusOK, user)
}

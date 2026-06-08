package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "som/internal/dto/artist"
	"som/internal/handlers"
	"som/internal/services/artist"
	"som/internal/util"
)

// CreateArtist godoc
// @summary      Registers a new artist.
// @tags         artist
// @accept       json
// @produce      json
// @security     BearerAuth
// @param        artist_info body dto.CreateArtistRequest true "Artist information"
// @success      201 "Artist created successfully"
// @failure      400 {object} util.ErrorResponse "Invalid request body"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      409 {object} util.ErrorResponse "Artist already exists"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /artist/create [post]
func CreateArtist(ginContext *gin.Context) {
	var req dto.CreateArtistRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "requisição inválida",
			Details: err.Error(),
		})
		return
	}

	err := artist.CreateArtist(ginContext.Request.Context(), req)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.Status(http.StatusCreated)
}

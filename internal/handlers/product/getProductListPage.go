package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "som/internal/dto/product"
	"som/internal/handlers"
	services "som/internal/services/product"
	"som/internal/util"
)

// GetProductListPage godoc
// @summary      Get a paginated list of products.
// @tags         product
// @accept       json
// @produce      json
// @security     BearerAuth
// @param        page_request body dto.GetProductListPageRequest true "Pagination parameters (page number, limit)"
// @success      200 {object} dto.GetProductListPageResponse "Page of products"
// @failure      400 {object} util.ErrorResponse "Invalid request body"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /product/list [post]
func GetProductListPage(ginContext *gin.Context) {
	var req dto.GetProductListPageRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "invalid request",
			Details: err.Error(),
		})
		return
	}

	productList, err := services.GetProductListPage(ginContext.Request.Context(), req)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.JSON(http.StatusOK, productList)
}

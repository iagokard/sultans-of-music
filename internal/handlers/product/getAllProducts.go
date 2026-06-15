package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"som/internal/handlers"
	services "som/internal/services/product"
)

// GetAllProducts godoc
// @summary      Get all products.
// @tags         product
// @produce      json
// @security     BearerAuth
// @success      200 {object} dto.GetProductListResponse "List of all products"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /product/all [get]
func GetAllProducts(ginContext *gin.Context) {
	productList, err := services.GetAllProducts(ginContext.Request.Context())
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.JSON(http.StatusOK, productList)
}

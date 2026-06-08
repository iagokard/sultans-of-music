package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"som/internal/handlers"
	services "som/internal/services/product"
	"som/internal/util"
)

// GetProductByID godoc
// @summary      Get product by ID.
// @tags         product
// @produce      json
// @security     BearerAuth
// @param        id path int true "Product ID"
// @success      200 {object} dto.GetProductResponse "Product found"
// @failure      400 {object} util.ErrorResponse "Invalid ID supplied"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      404 {object} util.ErrorResponse "Product not found"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /product/{id} [get]
func GetProductByID(ginContext *gin.Context) {
	idParam := ginContext.Param("id")

	productID, err := strconv.Atoi(idParam)
	if err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "requisição inválida",
			Details: "ID deve ser um número inteiro",
		})
		return
	}

	product, err := services.GetProductByID(ginContext.Request.Context(), productID)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.JSON(http.StatusOK, product)
}

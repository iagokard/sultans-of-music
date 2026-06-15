package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "som/internal/dto/product"
	"som/internal/handlers"
	services "som/internal/services/product"
	"som/internal/util"
)

// UpdateProduct godoc
// @summary      Update an existing product.
// @tags         product
// @accept       json
// @produce      json
// @security     BearerAuth
// @param        update_info body dto.UpdateProductRequest true "Product fields to update (all optional)"
// @success      200 "Product updated successfully"
// @failure      400 {object} util.ErrorResponse "Invalid request body"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      404 {object} util.ErrorResponse "Product not found"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /product/update [put]
func UpdateProduct(ginContext *gin.Context) {
	var req dto.UpdateProductRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "requisição inválida",
			Details: err.Error(),
		})
		return
	}

	err := services.UpdateProduct(ginContext.Request.Context(), req)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.Status(http.StatusOK)
}

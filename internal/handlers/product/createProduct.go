package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "som/internal/dto/product"
	"som/internal/handlers"
	services "som/internal/services/product"
	"som/internal/util"
)

// CreateProduct godoc
// @summary      Creates a new product.
// @tags         product
// @accept       json
// @produce      json
// @security     BearerAuth
// @param        product_info body dto.CreateProductRequest true "Product information"
// @success      201 "Product created successfully"
// @failure      400 {object} util.ErrorResponse "Invalid request body"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      409 {object} util.ErrorResponse "Product already exists"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /product/create [post]
func CreateProduct(ginContext *gin.Context) {
	var req dto.CreateProductRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "requisição inválida",
			Details: err.Error(),
		})
		return
	}

	err := services.CreateProduct(ginContext.Request.Context(), req)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.Status(http.StatusCreated)
}

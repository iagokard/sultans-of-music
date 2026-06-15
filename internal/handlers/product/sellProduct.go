package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "som/internal/dto/sale"
	"som/internal/handlers"
	services "som/internal/services/product"
	"som/internal/util"
)

// SellProduct godoc
// @summary      Sell a product.
// @tags         product
// @accept       json
// @produce      json
// @security     BearerAuth
// @param        sale_info body dto.SellProductRequest true "Product ID and quantity to sell"
// @success      201 {object} dto.SellProductResponse "Sale created successfully"
// @failure      400 {object} util.ErrorResponse "Invalid request body"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      409 {object} util.ErrorResponse "Insufficient stock or conflict"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /product/sell [post]
func SellProduct(ginContext *gin.Context) {
	var req dto.SellProductRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "requisição inválida",
			Details: err.Error(),
		})
		return
	}

	saleID, err := services.SellProduct(ginContext.Request.Context(), req)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.JSON(http.StatusCreated, dto.SellProductResponse{
		SaleID: saleID,
	})
}

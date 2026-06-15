package product

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "som/internal/dto/product"
	"som/internal/handlers"
	services "som/internal/services/product"
	"som/internal/util"
)

// UpdateInventory godoc
// @summary      Update a product's stock to a new absolute value.
// @tags         inventory
// @accept       json
// @produce      json
// @security     BearerAuth
// @param        inventory_update body dto.UpdateInventoryRequest true "Product ID and new stock value"
// @success      200 "Inventory updated successfully"
// @failure      400 {object} util.ErrorResponse "Invalid request body or negative stock"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      404 {object} util.ErrorResponse "Inventory row not found"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /inventory/update [put]
func UpdateInventory(ginContext *gin.Context) {
	var req dto.UpdateInventoryRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "requisição inválida",
			Details: err.Error(),
		})
		return
	}

	err := services.UpdateInventory(ginContext.Request.Context(), req)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.Status(http.StatusOK)
}

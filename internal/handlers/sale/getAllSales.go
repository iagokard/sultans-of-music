package sale

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"som/internal/handlers"
	services "som/internal/services/sale"
)

// GetAllSales godoc
// @summary      Get all sales.
// @tags         sale
// @produce      json
// @security     BearerAuth
// @success      200 {array} dto.GetSaleResponse "List of all sales"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /sale/all [get]
func GetAllSales(ginContext *gin.Context) {
	sales, err := services.GetAllSales(ginContext.Request.Context())
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.JSON(http.StatusOK, sales)
}

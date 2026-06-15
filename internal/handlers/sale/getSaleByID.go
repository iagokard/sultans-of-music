package sale

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"som/internal/handlers"
	services "som/internal/services/sale"
	"som/internal/util"
)

// GetSaleByID godoc
// @summary      Get sale by ID.
// @tags         sale
// @produce      json
// @security     BearerAuth
// @param        id path int true "Sale ID"
// @success      200 {object} dto.GetSaleResponse "Sale found"
// @failure      400 {object} util.ErrorResponse "Invalid ID supplied"
// @failure      401 "Unauthorized – missing or invalid token"
// @failure      404 {object} util.ErrorResponse "Sale not found"
// @failure      500 {object} util.ErrorResponse "Internal server error"
// @router       /sale/{id} [get]
func GetSaleByID(ginContext *gin.Context) {
	idParam := ginContext.Param("id")

	saleID, err := strconv.Atoi(idParam)
	if err != nil {
		ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
			Error:   "requisição inválida",
			Details: "ID deve ser um número inteiro",
		})
		return
	}

	sale, err := services.GetSaleByID(ginContext.Request.Context(), saleID)
	if err != nil {
		handlers.HandleDBError(ginContext, err)
		return
	}

	ginContext.JSON(http.StatusOK, sale)
}

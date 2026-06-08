package handlers

import (
	"errors"
	"som/internal/util"

	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func HandleDBError(ginContext *gin.Context, err error) {
	if errors.Is(err, sql.ErrNoRows) {
		ginContext.AbortWithStatusJSON(http.StatusNotFound, util.ErrorResponse{
			Error:   "item não encontrado",
			Details: err.Error(),
		})
		return
	}

	const DuplicateEntry = 1062
	const ForeignKeyViolation = 1452
	const NullConstraintViolation = 1048

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case DuplicateEntry:
			ginContext.AbortWithStatusJSON(http.StatusConflict, util.ErrorResponse{
				Error:   err.Error(),
				Details: mysqlErr.Message,
			})
			return
		case ForeignKeyViolation:
			ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
				Error:   "Invalid reference",
				Details: err.Error(),
			})
			return
		case NullConstraintViolation:
			ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
				Error:   "Required field missing",
				Details: err.Error(),
			})
			return
		}
	}

	ginContext.AbortWithStatusJSON(http.StatusBadRequest, util.ErrorResponse{
		Error:   "Internal server error",
		Details: err.Error(),
	})
}

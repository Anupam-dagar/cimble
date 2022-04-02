package utilities

import (
	"cimble/constants"
	"cimble/models"
	"encoding/hex"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ByteToString(inputBytes []byte) string {
	return hex.EncodeToString(inputBytes)
}

func ErrorCodeFromError(err error) int {
	switch constants.ErrorMessage(err.Error()) {
	case constants.Unauthorised:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func GetRequestBody(ctx *gin.Context, bodyModel interface{}) (err error) {
	err = ctx.ShouldBindJSON(bodyModel)
	if err != nil {
		ResponseWithErrorCode(ctx, http.StatusBadRequest, err)
		return
	}

	return
}

func GeneratePage(totalCount int64, offset int64, limit int64) models.Pagination {
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	currentPage := int((offset / limit) + 1)
	return models.Pagination{
		TotalPages:  totalPages,
		CurrentPage: currentPage,
	}
}

func GetOffsetAndLimit(ctx *gin.Context) (offset int64, limit int64) {
	offsetQuery := ctx.Query("offset")
	limitQuery := ctx.Query("limit")

	offset, _ = strconv.ParseInt(offsetQuery, 10, 64)
	limit, err := strconv.ParseInt(limitQuery, 10, 64)
	if err != nil {
		limit = 10
	}

	return offset, limit
}

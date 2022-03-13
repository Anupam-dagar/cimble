package utilities

import (
	"cimble/constants"
	"encoding/hex"
	"net/http"

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

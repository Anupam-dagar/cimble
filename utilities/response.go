package utilities

import (
	"cimble/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ResponseWithSuccess(ctx *gin.Context, code int, jsonObj interface{}) {
	ctx.JSON(code, jsonObj)
}

func ResponseWithError(ctx *gin.Context, err error) {
	fmt.Println(err.Error())
	code := ErrorCodeFromError(err)
	errorJson := models.ErrorResponse{
		Code:    fmt.Sprint(code),
		Message: err.Error(),
	}
	ctx.AbortWithStatusJSON(code, errorJson)
}

func ResponseWithErrorCode(ctx *gin.Context, code int, err error) {
	fmt.Println(err.Error())
	errorJson := models.ErrorResponse{
		Code:    fmt.Sprint(code),
		Message: err.Error(),
	}
	ctx.AbortWithStatusJSON(code, errorJson)
}

package utilities

import (
	"cimble/constants"
	"encoding/hex"
	"net/http"
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

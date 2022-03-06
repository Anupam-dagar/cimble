package utilities

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ParseJwt(token *jwt.Token) (interface{}, error) {
	if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
		return nil, fmt.Errorf("invalid token %v", token.Header["alg"])
	}

	jwtSecret := GetEnvironmentVariableString("JWT_SECRET")

	return []byte(jwtSecret), nil
}

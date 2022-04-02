package utilities

import (
	"cimble/constants"
	"crypto"
)

func GenerateSha512Hash(inputStr string, securityType constants.SecurityType) []byte {
	hasher := crypto.SHA512.New()

	var salt string
	switch securityType {
	case constants.PASSWORD:
		salt = GetEnvironmentVariableString("PASSWORD_SALT")
	case constants.ApiKey:
		salt = GetEnvironmentVariableString("API_KEY_SALT")
	}

	inputStr = salt + inputStr

	hasher.Write([]byte(inputStr))

	return hasher.Sum(nil)
}

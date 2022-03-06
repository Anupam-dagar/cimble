package utilities

import "crypto"

func GenerateSha512Hash(inputStr string) []byte {
	hasher := crypto.SHA512.New()
	hasher.Write([]byte(inputStr))

	return hasher.Sum(nil)
}

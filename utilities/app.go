package utilities

import "encoding/hex"

func ByteToString(inputBytes []byte) string {
	return hex.EncodeToString(inputBytes)
}

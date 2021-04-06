package utils

import (
	"encoding/base64"
)

func Base64Decode(src string) []byte {
	b, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		panic(err)
	}
	return b
}

func Base64Encode(src []byte) string {
	s := base64.StdEncoding.EncodeToString(src)
	return s
}

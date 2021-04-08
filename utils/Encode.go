package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
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

func Md5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

package encodeutil

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

//将byte转换16进制的字符串
func HexEncode(src []byte) string {
	return hex.EncodeToString(src)
}

//将16进制的字符串转换byte
func HexDecode(s string) []byte {
	h, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return h
}
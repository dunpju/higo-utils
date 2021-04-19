package utils

import "strings"

type SliceString []string

func (this SliceString) String(sep ...string) string {
	s := ""
	if len(sep) > 0 {
		s = sep[0]
	}
	return strings.Join(this, s)
}

//byte切片倒序
func ByteReverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

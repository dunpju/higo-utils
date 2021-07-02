package utils

import (
	"strconv"
	"unicode"
)

// recover错误，转string
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func Int64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func IntString(i int) string {
	return strconv.Itoa(i)
}

func Int64String(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
package utils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

//转换为字符串切片
func ConvStrSlice(payload interface{}) (ret []string) {
	if s, ok := payload.([]string); ok {
		for _, v := range s {
			ret = append(ret, v)
		}
	} else if i, ok := payload.([]int); ok {
		for _, v := range i {
			ret = append(ret, IntString(v))
		}
	} else if i64, ok := payload.([]int64); ok {
		for _, v := range i64 {
			ret = append(ret, Int64String(v))
		}
	} else {
		panic(fmt.Errorf("Unsupported types, Only support string or int or int64"))
	}
	return
}

//结构体转json
func ToJson(struc interface{}) string {
	mjson, err := json.Marshal(struc)
	if err != nil {
		panic(err)
	}
	return string(mjson)
}

//json key 转下划线
func JsonKeyToCase(str string) string {
	re, err := regexp.Compile(`("[a-zA-Z]*?":)`)
	if nil != err {
		panic(err)
	}
	pregReplace := re.ReplaceAllFunc([]byte(str), func(bytes []byte) []byte {
		return []byte(`"` + CamelToCase(strings.Replace(string(bytes), `"`, "", 1)))
	})
	return string(pregReplace)
}

//json key 转小驼峰
func JsonKeyToLcCamel(str string) string {
	re, err := regexp.Compile(`("[_a-zA-Z]*?":)`)
	if nil != err {
		panic(err)
	}
	pregReplace := re.ReplaceAllFunc([]byte(str), func(bytes []byte) []byte {
		return []byte(`"` + Lcfirst(strings.Replace(CaseToCamel(string(bytes)), `"`, "", 1)))
	})
	return string(pregReplace)
}

//json key 转大驼峰
func JsonKeyToCamel(str string) string {
	re, err := regexp.Compile(`("[_a-zA-Z]*?":)`)
	if nil != err {
		panic(err)
	}
	pregReplace := re.ReplaceAllFunc([]byte(str), func(bytes []byte) []byte {
		return []byte(CaseToCamel(string(bytes)))
	})
	return string(pregReplace)
}

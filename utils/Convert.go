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
	} else if f, ok := payload.([]float32); ok {
		for _, v := range f {
			ret = append(ret, FloatString(v))
		}
	} else if f64, ok := payload.([]float64); ok {
		for _, v := range f64 {
			ret = append(ret, Float64String(v))
		}
	} else {
		panic(fmt.Errorf("Unsupported types, Only support string or int/int64 or float32/float64"))
	}
	return
}

func ConvString(payload interface{}) (ret string) {
	if s, ok := payload.(string); ok {
		ret = s
	} else if i, ok := payload.(int); ok {
		ret = IntString(i)
	} else if i64, ok := payload.(int64); ok {
		ret = Int64String(i64)
	} else if f, ok := payload.(float32); ok {
		ret = FloatString(f)
	} else if f64, ok := payload.(float64); ok {
		ret = Float64String(f64)
	} else {
		panic(fmt.Errorf("Unsupported types, Only support string or int/int64 or float32/float64"))
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

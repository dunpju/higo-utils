package utils

import "fmt"

//转换为字符串切片
func ConvStrSlice(s []interface{}) (ret []string) {
	for _, i := range s {
		if v, ok := i.(string); ok {
			ret = append(ret, v)
		} else if v, ok := i.(int); ok {
			ret = append(ret, IntString(v))
		} else if v, ok := i.(int64); ok {
			ret = append(ret, Int64String(v))
		} else {
			panic(fmt.Errorf("Unsupported types, Only support string or int or int64"))
		}
	}
	return
}

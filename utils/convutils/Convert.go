package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/dengpju/higo-utils/utils"
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
			ret = append(ret, utils.IntString(v))
		}
	} else if i64, ok := payload.([]int64); ok {
		for _, v := range i64 {
			ret = append(ret, utils.Int64String(v))
		}
	} else if f, ok := payload.([]float32); ok {
		for _, v := range f {
			ret = append(ret, utils.FloatString(v))
		}
	} else if f64, ok := payload.([]float64); ok {
		for _, v := range f64 {
			ret = append(ret, utils.Float64String(v))
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
		ret = utils.IntString(i)
	} else if i64, ok := payload.(int64); ok {
		ret = utils.Int64String(i64)
	} else if f, ok := payload.(float32); ok {
		ret = utils.FloatString(f)
	} else if f64, ok := payload.(float64); ok {
		ret = utils.Float64String(f64)
	} else if bytes, ok := payload.([]byte); ok {
		ret = string(bytes)
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
		return []byte(`"` + utils.CamelToCase(strings.Replace(string(bytes), `"`, "", 1)))
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
		return []byte(`"` + utils.Lcfirst(strings.Replace(utils.CaseToCamel(string(bytes)), `"`, "", 1)))
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
		return []byte(utils.CaseToCamel(string(bytes)))
	})
	return string(pregReplace)
}

//大端序整形转换成字节
func Int32ToBytesBigEndian(n int32) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, n)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

//大端序字节转换成整形
func BytesToInt32BigEndian(b []byte) (int32, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	err := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		return 0, err
	}
	return x, nil
}

//小端序整形转换成字节
func Int32ToBytesLittleEndian(n int32) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, n)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

//小端序字节转换成整形
func BytesToInt32LittleEndian(b []byte) (int32, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	err := binary.Read(bytesBuffer, binary.LittleEndian, &x)
	if err != nil {
		return 0, err
	}
	return x, nil
}

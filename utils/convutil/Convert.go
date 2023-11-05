package convutil

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/dunpju/higo-utils/utils/stringutil"
	"regexp"
	"strings"
)

type Convert struct {
}

//转换为字符串切片
func (this *Convert) Slice(payload interface{}) (ret []string) {
	return ConvStrSlice(payload)
}

func (this *Convert) String(payload interface{}) (ret string) {
	return ConvString(payload)
}

//结构体转json
func (this *Convert) Json(struc interface{}) string {
	return ToJson(struc)
}

//json key 转下划线
func (this *Convert) JsonKeyToCase(str string) string {
	return JsonKeyToCase(str)
}

//json key 转小驼峰
func (this *Convert) JsonKeyToLcCamel(str string) string {
	return JsonKeyToLcCamel(str)
}

//json key 转大驼峰
func (this *Convert) JsonKeyToCamel(str string) string {
	return JsonKeyToCamel(str)
}

//大端序整形转换成字节
func (this *Convert) Int32ToBytesBigEndian(n int32) ([]byte, error) {
	return Int32ToBytesBigEndian(n)
}

//大端序字节转换成整形
func (this *Convert) BytesToInt32BigEndian(bs []byte) (int32, error) {
	return BytesToInt32BigEndian(bs)
}

//小端序整形转换成字节
func (this *Convert) Int32ToBytesLittleEndian(n int32) ([]byte, error) {
	return Int32ToBytesLittleEndian(n)
}

//小端序字节转换成整形
func (this *Convert) BytesToInt32LittleEndian(bs []byte) (int32, error) {
	return BytesToInt32LittleEndian(bs)
}

//大端序整形转换成字节
func (this *Convert) Int64ToBytesBigEndian(n int64) ([]byte, error) {
	return Int64ToBytesBigEndian(n)
}

//大端序字节转换成整形
func (this *Convert) BytesToInt64BigEndian(bs []byte) (int64, error) {
	return BytesToInt64BigEndian(bs)
}

//小端序整形转换成字节
func (this *Convert) Int64ToBytesLittleEndian(n int64) ([]byte, error) {
	return Int64ToBytesLittleEndian(n)
}

//小端序字节转换成整形
func (this *Convert) BytesToInt64LittleEndian(bs []byte) (int64, error) {
	return BytesToInt64LittleEndian(bs)
}

func ConvStrSlice(payload interface{}) (ret []string) {
	if s, ok := payload.([]string); ok {
		ret = append(ret, s...)
	} else if i, ok := payload.([]int); ok {
		for _, v := range i {
			ret = append(ret, stringutil.IntString(v))
		}
	} else if i64, ok := payload.([]int64); ok {
		for _, v := range i64 {
			ret = append(ret, stringutil.Int64String(v))
		}
	} else if f, ok := payload.([]float32); ok {
		for _, v := range f {
			ret = append(ret, stringutil.FloatString(v))
		}
	} else if f64, ok := payload.([]float64); ok {
		for _, v := range f64 {
			ret = append(ret, stringutil.Float64String(v))
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
		ret = stringutil.IntString(i)
	} else if i64, ok := payload.(int64); ok {
		ret = stringutil.Int64String(i64)
	} else if f, ok := payload.(float32); ok {
		ret = stringutil.FloatString(f)
	} else if f64, ok := payload.(float64); ok {
		ret = stringutil.Float64String(f64)
	} else if bs, ok := payload.([]byte); ok {
		ret = string(bs)
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
		return []byte(`"` + stringutil.CamelToCase(strings.Replace(string(bytes), `"`, "", 1)))
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
		return []byte(`"` + stringutil.Lcfirst(strings.Replace(stringutil.CaseToCamel(string(bytes)), `"`, "", 1)))
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
		return []byte(stringutil.CaseToCamel(string(bytes)))
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

//大端序整形转换成字节
func Int64ToBytesBigEndian(n int64) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, n)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

//大端序字节转换成整形
func BytesToInt64BigEndian(b []byte) (int64, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	err := binary.Read(bytesBuffer, binary.BigEndian, &x)
	if err != nil {
		return 0, err
	}
	return x, nil
}

//小端序整形转换成字节
func Int64ToBytesLittleEndian(n int64) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, n)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

//小端序字节转换成整形
func BytesToInt64LittleEndian(b []byte) (int64, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	err := binary.Read(bytesBuffer, binary.LittleEndian, &x)
	if err != nil {
		return 0, err
	}
	return x, nil
}

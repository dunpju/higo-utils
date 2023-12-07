package stringutil

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
)

type Stringutil struct {
}

// recover错误，转string
func (this *Stringutil) Error(r interface{}) string {
	return ErrorToString(r)
}

func (this *Stringutil) Int(str string) int {
	return Int(str)
}

func (this *Stringutil) Int64(str string) int64 {
	return Int64(str)
}

func (this *Stringutil) IntString(i int) string {
	return IntString(i)
}

func (this *Stringutil) Int64String(i int64) string {
	return Int64String(i)
}

func (this *Stringutil) FloatString(f float32) string {
	return FloatString(f)
}

func (this *Stringutil) Float64String(f64 float64) string {
	return Float64String(f64)
}

func (this *Stringutil) Ucfirst(str string) string {
	return Ucfirst(str)
}

func (this *Stringutil) Lcfirst(str string) string {
	return Lcfirst(str)
}

func (this *Stringutil) CamelToCase(name string) string {
	return CamelToCase(name)
}

func (this *Stringutil) CaseToCamel(name string) string {
	return CaseToCamel(name)
}

func (this *Stringutil) Buffer() *Buffer {
	return NewBuffer()
}

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

func FloatString(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

func Float64String(f64 float64) string {
	return strconv.FormatFloat(f64, 'f', -1, 64)
}

// Ucfirst 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// Lcfirst 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// LCamelToCase 首字母小写驼峰
func LCamelToCase(name string) string {
	return Lcfirst(CamelToCase(name))
}

// UCamelToCase 首字母大写驼峰
func UCamelToCase(name string) string {
	return Ucfirst(CamelToCase(name))
}

// CamelToCase 驼峰式写法转为下划线写法
func CamelToCase(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// CaseToCamel 下划线写法转为驼峰写法
func CaseToCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// Buffer 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		_, err := b.Write(val)
		if err != nil {
			panic(err)
		}
	case rune:
		_, err := b.WriteRune(val)
		if err != nil {
			panic(err)
		}
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	_, err := b.WriteString(s)
	if err != nil {
		panic(err)
	}
	return b
}

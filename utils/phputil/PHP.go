package phputil

import (
	"encoding/json"
	"github.com/dengpju/higo-utils/utils/convutil"
	"github.com/dengpju/higo-utils/utils/dirutil"
	"github.com/dengpju/higo-utils/utils/ufuncutil"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
)

//path文件名
func Basename(path string, suffix ...string) string {
	suff := ""
	if len(suffix) > 0 {
		suff = suffix[0]
	}
	paths := strings.Split(path, dirutil.PathSeparator())
	name := ufuncutil.IfStringIndex(paths[len(paths)-1:], 0)
	if suff != "" {
		names := strings.Split(name, ".")
		name = ufuncutil.IfStringIndex(names, 0)
	}
	return name
}

func Dirname(path string) string {
	paths := strings.Split(path, dirutil.PathSeparator())
	paths = paths[:len(paths)-1]
	return strings.Join(paths, dirutil.PathSeparator())
}

func DirBasename(path string) string {
	paths := strings.Split(path, dirutil.PathSeparator())
	paths = paths[len(paths)-2 : len(paths)-1]
	return strings.Join(paths, dirutil.PathSeparator())
}

//目录path切片
func Dirslice(path string) []string {
	paths := strings.Split(path, dirutil.PathSeparator())
	return paths[:len(paths)-1]
}

// 目录是否存在
func DirExist(dirname string) bool {
	if _, err := os.Stat(dirname); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

//创建目录
func Mkdir(dirname string, perm ...os.FileMode) bool {
	if len(perm) > 0 {
		dirutil.SetModePerm(perm[0])
	}
	var dir []string
	for _, p := range Dirslice(dirname) {
		dir = append(dir, p)
		tmpPath := strings.Join(dir, dirutil.PathSeparator())
		if _, err := os.Stat(tmpPath); err != nil {
			if os.IsNotExist(err) {
				if tmpPath != "" {
					if err := os.Mkdir(tmpPath, dirutil.ModePerm()); err != nil {
						panic(err)
					}
				}
			} else {
				panic(err)
			}
		}
	}
	return true
}

//删除目录
func Rmdir(dirname string) bool {
	err := os.RemoveAll(dirname)
	if err != nil {
		panic(err)
	}
	return true
}

//清空目录
func Emdir(dirname string) bool {
	dir, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, d := range dir {
		err = os.RemoveAll(path.Join([]string{dirname, d.Name()}...))
		if err != nil {
			panic(err)
		}
	}
	return true
}

//删除文件
func Remove(filename string) bool {
	err := os.Remove(filename)
	if err != nil {
		panic(err)
	}
	return true
}

// 创建文件
func Mkfile(filename string) *os.File {
	// 目录不存在，并创建
	Mkdir(filename)
	//创建文件
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return f
}

//path切片 -> string
func Pathstring(paths []string) string {
	return strings.Join(paths, dirutil.PathSeparator())
}

func ToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	meta := make(map[string]interface{})
	if objValue.Kind() == reflect.Ptr {
		v := objValue.Elem()
		typeOfType := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.CanInterface() {
				jsonTag := typeOfType.Field(i).Tag.Get("json")
				if jsonTag != "" {
					meta[jsonTag] = field.Interface()
				} else {
					meta[typeOfType.Field(i).Name] = field.Interface()
				}
			}
		}
	} else {
		typeOfType := reflect.TypeOf(obj)
		for i := 0; i < typeOfType.NumField(); i++ {
			if objValue.Field(i).CanInterface() {
				jsonTag := typeOfType.Field(i).Tag.Get("json")
				if jsonTag != "" {
					meta[jsonTag] = objValue.Field(i).Interface()
				} else {
					meta[typeOfType.Field(i).Name] = objValue.Field(i).Interface()
				}
			}
		}
	}
	return meta
}

func JsonDecode(str string) (meta map[string]interface{}) {
	if err := json.Unmarshal([]byte(str), &meta); err != nil {
		panic(err)
	}
	return
}

func JsonEncode(meta interface{}) string {
	return convutil.ToJson(meta)
}

func Isset() {

}

func InArray() {

}

func ArrayFilter() {

}

func ArrayUnique() {

}

func ArrayColumn(array, column string) {

}

func ArrayCombine(obj interface{}, key, value string) map[string]interface{} {
	js, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	array := make([]map[string]interface{}, 0)
	err = json.Unmarshal(js, &array)
	if err != nil {
		panic(err)
	}
	ret := make(map[string]interface{})
	for _, m := range array {
		k, ok := m[key]
		if !ok {
			panic("There is no key")
		}
		v, ok := m[value]
		if !ok {
			panic("There is no value")
		}
		ret[convutil.ConvString(k)] = v
	}
	return ret
}

func Strpos() {

}

func Min(value interface{}) (min interface{}) {
	if i, ok := value.([]int); ok {
		m := i[0]
		for j := 1; j < len(i); j++ {
			if m > i[j] {
				m = i[j]
			}
		}
		min = m
	} else if i64, ok := value.([]int64); ok {
		m := i64[0]
		for j := 1; j < len(i64); j++ {
			if m > i64[j] {
				m = i64[j]
			}
		}
		min = m
	} else if f, ok := value.([]float32); ok {
		m := f[0]
		for j := 1; j < len(f); j++ {
			if m > f[j] {
				m = f[j]
			}
		}
		min = m
	} else if f64, ok := value.([]float64); ok {
		m := f64[0]
		for j := 1; j < len(f64); j++ {
			if m > f64[j] {
				m = f64[j]
			}
		}
		min = m
	}
	return
}

func Max(value interface{}) (max interface{}) {
	if i, ok := value.([]int); ok {
		m := i[0]
		for j := 1; j < len(i); j++ {
			if m < i[j] {
				m = i[j]
			}
		}
		max = m
	} else if i64, ok := value.([]int64); ok {
		m := i64[0]
		for j := 1; j < len(i64); j++ {
			if m < i64[j] {
				m = i64[j]
			}
		}
		max = m
	} else if f, ok := value.([]float32); ok {
		m := f[0]
		for j := 1; j < len(f); j++ {
			if m < f[j] {
				m = f[j]
			}
		}
		max = m
	} else if f64, ok := value.([]float64); ok {
		m := f64[0]
		for j := 1; j < len(f64); j++ {
			if m < f64[j] {
				m = f64[j]
			}
		}
		max = m
	}
	return
}
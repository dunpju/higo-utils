package phputil

import (
	"encoding/json"
	"github.com/dengpju/higo-utils/utils/convutil"
	"reflect"
)

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

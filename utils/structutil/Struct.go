package structutil

import (
	"fmt"
	"reflect"
)

type Struct struct {
}

func (this *Struct) IsEmpty(_struct interface{}) (bool, error) {
	return IsEmpty(_struct)
}

func IsEmpty(_struct interface{}) (bool, error) {
	v := reflect.ValueOf(_struct)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return false, fmt.Errorf("not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		if !reflect.DeepEqual(v.Field(i).Interface(), reflect.Zero(v.Field(i).Type()).Interface()) {
			return false, nil
		}
	}
	return true, nil
}

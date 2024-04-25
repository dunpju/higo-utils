package structutil

import "reflect"

type Struct struct {
}

func (this *Struct) IsEmpty(_struct interface{}) bool {
	return IsEmpty(_struct)
}

func IsEmpty(_struct interface{}) bool {
	return reflect.DeepEqual(_struct, reflect.Zero(reflect.TypeOf(_struct)).Interface())
}

package listutil

import (
	"fmt"
	"github.com/dunpju/higo-utils/utils/maputil"
	"reflect"
)

type List[T any] struct {
	list []T
}

func NewList[T any](list []T) *List[T] {
	return &List[T]{list: list}
}

func (this *List[T]) FindIndex(find string) int {
	for i, s := range this.list {
		if find == fmt.Sprintf("%v", s) {
			return i
		}
	}
	return -1
}

// Group 分组
// fieldName 字段名
// 返回值: map util.ArrayMap
// 示例: customerFollowRecordGroup := utils.List[*vo.FollowRecordGlueVO](followRecordGlueVOList).Group("CustomerFollowRecordId")
// followRecordVO.GlueInfoJson = customerFollowRecordGroup.Get(followRecordVO.Id).([]*vo.FollowRecordGlueVO)
func (this *List[T]) Group(fieldName string) *maputil.ArrayMap {
	array := maputil.Array()
	for _, s := range this.list {
		t := reflect.TypeOf(s)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		v := reflect.ValueOf(s)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Name == fieldName {
				val := v.Field(i).Interface()
				if array.Exist(fmt.Sprintf("%v", val)) {
					valList := array.Get(fmt.Sprintf("%v", val)).([]T)
					valList = append(valList, s)
					array.Put(fmt.Sprintf("%v", val), valList)
				} else {
					valList := make([]T, 0)
					valList = append(valList, s)
					array.Put(fmt.Sprintf("%v", val), valList)
				}
			}
		}
	}
	return array
}

// ColumnValueMap
// fieldName 字段名
// 返回值: map util.ArrayMap
// 示例: orderSalesIdMap := utils.List[*vo.SaleVO](orderSalesList).ColumnValueMap("OrderSalesId")
func (this *List[T]) ColumnValueMap(fieldName string) *maputil.ArrayMap {
	array := maputil.Array()
	for _, s := range this.list {
		t := reflect.TypeOf(s)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		v := reflect.ValueOf(s)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Name == fieldName {
				val := v.Field(i).Interface()
				array.Put(fmt.Sprintf("%v", val), val)
			}
		}
	}
	return array
}

func (this *List[T]) ToMap(fieldName string) *maputil.ArrayMap {
	array := maputil.Array()
	for _, s := range this.list {
		t := reflect.TypeOf(s)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		v := reflect.ValueOf(s)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Name == fieldName {
				val := v.Field(i).Interface()
				array.Put(fmt.Sprintf("%v", val), val)
			}
		}
	}
	return array
}

func (this *List[T]) Distinct() []T {
	distinctMap := make(map[interface{}]T)
	res := make([]T, 0)
	for _, s := range this.list {
		if _, ok := distinctMap[s]; !ok {
			distinctMap[s] = s
			res = append(res, s)
		}
	}
	return res
}

// ListToMap
// list 列表
// key 字段名
// 返回值: map util.ArrayMap
// 示例: sysDictStateMap = utils.ListToMap[*SysDict.Model](sysDictStateList, "DictValue")
func ListToMap[T any](list []T, key string) *maputil.ArrayMap {
	array := maputil.Array()
	for _, s := range list {
		t := reflect.TypeOf(s)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		v := reflect.ValueOf(s)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Name == key {
				val := v.Field(i).Interface()
				array.Put(fmt.Sprintf("%v", val), s)
			}
		}
	}
	return array
}

type ListColumn[T, E any] struct {
	list []T
}

// Column
// customerFollowRecordIds := utils.Column[*vo.FollowRecordVO, string](*res.GetItems().(*[]*vo.FollowRecordVO)).List("Id")
func Column[T any, E any](list []T) *ListColumn[T, E] {
	return &ListColumn[T, E]{list: list}
}

func (this *ListColumn[T, E]) List(fieldName string) []E {
	duplicate := make(map[interface{}]bool)
	res := make([]E, 0)
	for _, s := range this.list {
		t := reflect.TypeOf(s)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		v := reflect.ValueOf(s)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Name == fieldName {
				val := v.Field(i).Interface()
				if _, ok := duplicate[val]; !ok {
					duplicate[val] = true
					res = append(res, val.(E))
				}
			}
		}
	}
	return res
}

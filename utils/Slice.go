package utils

import (
	"strings"
)

type ISlice interface {
	Append(value interface{}) interface{}
	Insert(index int, value interface{}) interface{}
	Delete(index int) interface{}
	Remove(dist interface{}) interface{}
	Replace(src, dist interface{}) interface{}
	String() string
	Separator(sep ...interface{}) string
	Exist(value interface{}) bool
	Set(value interface{}) interface{}
	Value() interface{}
	ForEach(callable SliceCallable)
	Len() int
	Reverse()
	Empty()
}

type SliceString struct {
	value []string
}

func NewSliceString(src ...string) *SliceString {
	sl := &SliceString{}
	for _, s := range src {
		sl.Append(s)
	}
	return sl
}

func (this *SliceString) String() string {
	return "[" + strings.Join(this.value, " ") + "]"
}

func (this *SliceString) Separator(sep ...interface{}) string {
	s := ""
	if len(sep) > 0 {
		s = sep[0].(string)
	}
	return strings.Join(this.value, s)
}

func (this *SliceString) Append(value interface{}) interface{} {
	this.value = append(this.value, value.(string))
	return this
}

func (this *SliceString) Insert(index int, value interface{}) interface{} {
	var tmp []string
	tmp = append(tmp, this.value[index:]...)
	this.value = append(this.value[0:index], value.(string))
	this.value = append(this.value, tmp...)
	return this
}

func (this *SliceString) Remove(dist interface{}) interface{} {
	var tmp []string
	this.ForEach(func(index int, value interface{}) {
		if value.(string) != dist.(string) {
			tmp = append(tmp, value.(string))
		}
	})
	this.value = tmp
	return this
}

func (this *SliceString) Delete(index int) interface{} {
	this.value = append(this.value[:index], this.value[index+1:]...)
	return this
}

func (this *SliceString) Replace(src, dist interface{}) interface{} {
	this.ForEach(func(index int, value interface{}) {
		if value.(string) == src.(string) {
			this.value[index] = dist.(string)
		}
	})
	return this
}

func (this *SliceString) Empty() {
	this.value = this.value[0:0]
}

func (this *SliceString) Len() int {
	return len(this.value)
}

type SliceCallable func(index int, value interface{})

func (this *SliceString) ForEach(callable SliceCallable) {
	for key, value := range this.value {
		callable(key, value)
	}
}

func (this *SliceString) Reverse() {
	for i, j := 0, len(this.value)-1; i < j; i, j = i+1, j-1 {
		this.value[i], this.value[j] = this.value[j], this.value[i]
	}
}

func (this *SliceString) Exist(dist interface{}) bool {
	for _, value := range this.value {
		if value == dist.(string) {
			return true
		}
	}
	return false
}

func (this *SliceString) Set(value interface{}) interface{} {
	this.value = value.([]string)
	return this
}

func (this *SliceString) Value() interface{} {
	return this.value
}

//byte切片倒序
func ByteReverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

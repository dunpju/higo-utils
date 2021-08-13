package utils

import (
	"fmt"
	"strings"
	"sync"
)

type ISlice interface {
	Append(value interface{}) interface{}
	Insert(index int, value interface{}) interface{}
	Delete(index int) interface{}
	Remove(dist interface{}) interface{}
	Replace(src, dist interface{}) interface{}
	String() string
	Join(sep ...string) string
	Exist(value interface{}) bool
	Clone(src *SliceString) interface{}
	Set(value interface{}) interface{}
	Value() interface{}
	ForEach(callable SliceCallable)
	Len() int
	Reverse()
	Empty()
}

type SliceString struct {
	lock  sync.Mutex
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
	return fmt.Sprintf("%s", this.value)
}

func (this *SliceString) Join(sep ...string) string {
	s := ""
	if len(sep) > 0 {
		s = sep[0]
	}
	return strings.Join(this.value, s)
}

func (this *SliceString) Append(value interface{}) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value = append(this.value, value.(string))
	return this
}

func (this *SliceString) Insert(index int, value interface{}) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	tmp := make([]string, 0)
	tmp = append(tmp, this.value[index:]...)
	this.value = append(this.value[0:index], value.(string))
	this.value = append(this.value, tmp...)
	return this
}

func (this *SliceString) Remove(dist interface{}) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	tmp := make([]string, 0)
	this.ForEach(func(index int, value interface{}) {
		if value.(string) != dist.(string) {
			tmp = append(tmp, value.(string))
		}
	})
	this.value = tmp
	return this
}

func (this *SliceString) Delete(index int) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value = append(this.value[:index], this.value[index+1:]...)
	return this
}

func (this *SliceString) Replace(src, dist interface{}) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.ForEach(func(index int, value interface{}) {
		if value.(string) == src.(string) {
			this.value[index] = dist.(string)
		}
	})
	return this
}

func (this *SliceString) Empty() {
	this.lock.Lock()
	defer this.lock.Unlock()
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
	this.lock.Lock()
	defer this.lock.Unlock()
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
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value = value.([]string)
	return this
}

func (this *SliceString) Value() interface{} {
	return this.value
}

func (this *SliceString) Clone(src *SliceString) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	copy(this.value, src.value)
	return this
}

//byte切片倒序
func ByteReverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

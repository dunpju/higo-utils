package sliceutil

import (
	"fmt"
	"strings"
	"sync"
)

type ISlice[T any] interface {
	Append(value T) T
	Insert(index int, value T) T
	Delete(index int) T
	Remove(dist T) T
	Replace(src, dist T) T
	String() string
	Join(sep ...string) string
	Exist(value T) bool
	Clone(src *T) T
	Set(value T) T
	Value() T
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

func (this *SliceString) Append(value string) *SliceString {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value = append(this.value, value)
	return this
}

func (this *SliceString) Insert(index int, value string) *SliceString {
	this.lock.Lock()
	defer this.lock.Unlock()
	tmp := make([]string, 0)
	tmp = append(tmp, this.value[index:]...)
	this.value = append(this.value[0:index], value)
	this.value = append(this.value, tmp...)
	return this
}

func (this *SliceString) Remove(dist string) *SliceString {
	this.lock.Lock()
	defer this.lock.Unlock()
	tmp := make([]string, 0)
	this.ForEach(func(index int, value interface{}) bool {
		if value.(string) != dist {
			tmp = append(tmp, value.(string))
		}
		return true
	})
	this.value = tmp
	return this
}

func (this *SliceString) Delete(index int) *SliceString {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value = append(this.value[:index], this.value[index+1:]...)
	return this
}

func (this *SliceString) Replace(src, dist string) *SliceString {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.ForEach(func(index int, value interface{}) bool {
		if value.(string) == src {
			this.value[index] = dist
		}
		return true
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

type SliceCallable func(index int, value interface{}) bool

func (this *SliceString) ForEach(callable SliceCallable) {
	for key, value := range this.value {
		if !callable(key, value) {
			break
		}
	}
}

func (this *SliceString) Reverse() {
	this.lock.Lock()
	defer this.lock.Unlock()
	for i, j := 0, len(this.value)-1; i < j; i, j = i+1, j-1 {
		this.value[i], this.value[j] = this.value[j], this.value[i]
	}
}

func (this *SliceString) Exist(dist string) bool {
	for _, value := range this.value {
		if value == dist {
			return true
		}
	}
	return false
}

func (this *SliceString) Set(value []string) *SliceString {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value = value
	return this
}

func (this *SliceString) Value() []string {
	return this.value
}

func (this *SliceString) Clone(src *SliceString) *SliceString {
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

type Sliceutil struct {
}

func (this *Sliceutil) New(src ...string) *SliceString {
	return NewSliceString(src...)
}

func (this *Sliceutil) ByteReverse(s []byte) []byte {
	return ByteReverse(s)
}

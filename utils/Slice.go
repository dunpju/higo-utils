package utils

import "strings"

type SliceString []string

func NewSliceString() SliceString {
	var sliceString SliceString
	return sliceString
}

func (this SliceString) String(sep ...string) string {
	s := ""
	if len(sep) > 0 {
		s = sep[0]
	}
	return strings.Join(this, s)
}

func (this SliceString) Append(value string) SliceString {
	this = append(this, value)
	return this
}

func (this SliceString) Remove(dist string) SliceString {
	var tmp = SliceString{}
	this.ForEach(func(index int, value string) {
		if value != dist {
			tmp.Append(value)
		}
	})
	return tmp
}

func (this SliceString) Replace(src, dist string) SliceString {
	this.ForEach(func(index int, value string) {
		if value == src {
			this[index] = dist
		}
	})
	return this
}

func (this SliceString) Empty() {
	this = nil
}

func (this SliceString) Len() int {
	return len(this)
}

type SliceStringCallable func(index int, value string)

func (this SliceString) ForEach(callable SliceStringCallable) {
	for key, value := range this {
		callable(key, value)
	}
}

//byte切片倒序
func ByteReverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

package ufuncutil

type Ufuncutil struct {
}

func (this *Ufuncutil) If(condition bool, a, b interface{}) interface{} {
	return If(condition, a, b)
}

func (this *Ufuncutil) Ifindex(slice []interface{}, index int) interface{} {
	return Ifindex(slice, index)
}

func (this *Ufuncutil) IfStringIndex(slice []string, index int) string {
	return IfStringIndex(slice, index)
}

func (this *Ufuncutil) IfIntIndex(slice []int, index int) int {
	return IfIntIndex(slice, index)
}

func (this *Ufuncutil) IfInt64Index(slice []int64, index int) int64 {
	return IfInt64Index(slice, index)
}

// If 三目运算
func If(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// Ifindex 如果index存在，则返回切片对应index值
func Ifindex(slice []interface{}, index int) interface{} {
	if len(slice) > index {
		return slice[index]
	}
	return nil
}

func IfStringIndex(slice []string, index int) string {
	if len(slice) > index {
		return slice[index]
	}
	panic("index nonexistent")
}

func IfIntIndex(slice []int, index int) int {
	if len(slice) > index {
		return slice[index]
	}
	panic("index nonexistent")
}

func IfInt64Index(slice []int64, index int) int64 {
	if len(slice) > index {
		return slice[index]
	}
	panic("index nonexistent")
}

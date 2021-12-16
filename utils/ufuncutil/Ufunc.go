package ufuncutil

// 三目运算
func If(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// 如果index存在，则返回切片对应index值
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

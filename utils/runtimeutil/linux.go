package runtimeutil

import (
	"bytes"
	"runtime"
	"strconv"
)

// 获取当前协程ID
func GoroutineID() (uint64, error) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	return n, err
}

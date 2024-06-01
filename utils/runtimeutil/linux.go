package runtimeutil

import (
	"fmt"
	"runtime"
)

// GoroutineID 获取当前协程ID
func GoroutineID() (uint64, error) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	var id uint64
	_, err := fmt.Sscanf(string(b), "goroutine %d", &id)
	return id, err
}

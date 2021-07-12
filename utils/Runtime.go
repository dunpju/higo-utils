package utils

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"syscall"
)

// 获取当前协程ID
func GoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

// 获取当前线程ID
func ThreadID() int {
	var user32 *syscall.DLL
	var GetCurrentThreadId *syscall.Proc
	var err error

	user32, err = syscall.LoadDLL("Kernel32.dll")
	if err != nil {
		panic(fmt.Sprintf("syscall.LoadDLL fail: %v\n", err.Error()))
		return 0
	}
	GetCurrentThreadId, err = user32.FindProc("GetCurrentThreadId")
	if err != nil {
		panic(fmt.Sprintf("user32.FindProc fail: %v\n", err.Error()))
		return 0
	}

	var pid uintptr
	pid, _, err = GetCurrentThreadId.Call()

	return int(pid)
}

// +build windows

package runtimeutil

import (
	"fmt"
	"syscall"
)

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

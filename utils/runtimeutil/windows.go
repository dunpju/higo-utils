//go:build windows
// +build windows

package runtimeutil

import (
	"syscall"
)

// ThreadID 获取当前线程ID
func ThreadID() (uint64, error) {
	var user32 *syscall.DLL
	var getCurrentThreadId *syscall.Proc
	var err error

	user32, err = syscall.LoadDLL("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	getCurrentThreadId, err = user32.FindProc("GetCurrentThreadId")
	if err != nil {
		return 0, err
	}

	var pid uintptr
	pid, _, err = getCurrentThreadId.Call()
	if err != nil {
		return 0, err
	}

	return uint64(pid), nil
}

func (this *Runtime) ThreadID() (uint64, error) {
	return ThreadID()
}

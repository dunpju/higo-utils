package runtimeutil

type Runtime struct {
}

func (this *Runtime) GoroutineID() uint64 {
	return GoroutineID()
}

func (this *Runtime) ThreadID() int {
	return ThreadID()
}

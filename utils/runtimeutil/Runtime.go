package runtimeutil

type Runtime struct {
}

func (this *Runtime) GoroutineID() (uint64, error) {
	return GoroutineID()
}

func (this *Runtime) ThreadID() (uint64, error) {
	return ThreadID()
}

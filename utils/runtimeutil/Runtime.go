package runtimeutil

type Runtime struct {
}

func (this *Runtime) GoroutineID() (uint64, error) {
	return GoroutineID()
}

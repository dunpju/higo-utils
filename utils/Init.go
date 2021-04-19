package utils

import "os"

var (
	pathSeparator = string(os.PathSeparator)
	modePerm      = os.ModePerm
)

func PathSeparator() string {
	return pathSeparator
}

func SetPathSeparator(sep string) string {
	pathSeparator = sep
	return pathSeparator
}

func ModePerm() os.FileMode {
	return modePerm
}

func SetModePerm(mode os.FileMode) os.FileMode {
	modePerm = mode
	return modePerm
}

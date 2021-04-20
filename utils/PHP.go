package utils

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//path文件名
func Basename(path string, suffix ...string) string {
	suff := ""
	if len(suffix) > 0 {
		suff = suffix[0]
	}
	paths := strings.Split(path, pathSeparator)
	name := IfStringIndex(paths[len(paths)-1:], 0)
	if suff != "" {
		names := strings.Split(name, ".")
		name = IfStringIndex(names, 0)
	}
	return name
}

func Dirname(path string) string {
	paths := strings.Split(path, pathSeparator)
	paths = paths[:len(paths)-1]
	return strings.Join(paths, pathSeparator)
}

func DirBasename(path string) string {
	paths := strings.Split(path, pathSeparator)
	paths = paths[len(paths)-2 : len(paths)-1]
	return strings.Join(paths, pathSeparator)
}

//目录path切片
func Dirslice(path string) []string {
	paths := strings.Split(path, pathSeparator)
	return paths[:len(paths)-1]
}

// 目录是否存在
func DirExist(dirname string) bool {
	if _, err := os.Stat(dirname); err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

//创建目录
func Mkdir(dirname string, perm ...os.FileMode) bool {
	if len(perm) > 0 {
		modePerm = perm[0]
	}
	var dir []string
	for _, p := range Dirslice(dirname) {
		dir = append(dir, p)
		tmpPath := strings.Join(dir, pathSeparator)
		if _, err := os.Stat(tmpPath); err != nil {
			if os.IsNotExist(err) {
				if tmpPath != "" {
					if err := os.Mkdir(tmpPath, modePerm); err != nil {
						panic(err)
					}
				}
			} else {
				panic(err)
			}
		}
	}
	return true
}

//删除目录
func Rmdir(dirname string) bool {
	err := os.RemoveAll(dirname);
	if err != nil {
		panic(err)
	}
	return true
}

//清空目录
func Emdir(dirname string) bool {
	dir, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, d := range dir {
		err = os.RemoveAll(path.Join([]string{dirname, d.Name()}...))
		if err != nil {
			panic(err)
		}
	}
	return true
}

//删除文件
func Remove(filename string) bool {
	err := os.Remove(filename);
	if err != nil {
		panic(err)
	}
	return true
}

// 创建文件
func Mkfile(filename string) *os.File {
	// 目录不存在，并创建
	Mkdir(filename)
	//创建文件
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return f
}

//path切片 -> string
func Pathstring(paths []string) string {
	return strings.Join(paths, pathSeparator)
}

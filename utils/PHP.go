package utils

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var PathSeparator = string(os.PathSeparator)
var ModePerm = os.ModePerm

func Basename(path string, suffix ...string) string {
	suff := ""
	if len(suffix) > 0 {
		suff = suffix[0]
	}
	paths := strings.Split(path, PathSeparator)
	name := IfStringIndex(paths[len(paths)-1:], 0)
	if suff != "" {
		names := strings.Split(name, ".")
		name = IfStringIndex(names, 0)
	}
	return name
}

func Dirname(path string) string {
	paths := strings.Split(path, PathSeparator)
	paths = paths[:len(paths)-1]
	return strings.Join(paths, PathSeparator)
}

func Dirslice(path string) []string {
	paths := strings.Split(path, PathSeparator)
	return paths[:len(paths)-1]
}

func Mkdir(dirname string, perm ...os.FileMode) bool {
	if len(perm) > 0 {
		ModePerm = perm[0]
	}
	var dir []string
	for _, p := range Dirslice(dirname) {
		dir = append(dir, p)
		tmpPath := strings.Join(dir, PathSeparator)
		if _, err := os.Stat(tmpPath); os.IsNotExist(err) && tmpPath != "" {
			if err := os.Mkdir(tmpPath, ModePerm); err != nil {
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

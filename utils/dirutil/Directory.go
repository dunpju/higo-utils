package dirutil

import (
	"github.com/dengpju/higo-utils/utils/ufuncutil"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

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

type Directory struct {
	path   string   // 扫描完整路径
	suffix string   // 文件后缀
	list   []string // 文件列表
	scan   bool     // 是否执行
}

func Dir(path string) *Directory {
	return &Directory{path: path, suffix: "*", scan: false}
}

// 后缀
func (this *Directory) Suffix(suffix string) *Directory {
	this.suffix = suffix
	return this
}

// 获取
func (this *Directory) Get() []string {
	if this.scan {
		this.list = scanner(this.path, this.suffix)
	} else {
		panic("There is no scan")
	}
	return this.list
}

// 创建目录
func (this *Directory) Create() *Directory {
	// 目录不存在，并创建
	if _, err := os.Stat(this.path); os.IsNotExist(err) {
		if err := os.MkdirAll(this.path, os.ModePerm); err != nil {
			panic(err)
		}
	}
	return this
}

// 扫描
func (this *Directory) Scan() *Directory {
	this.scan = true
	return this
}

// 扫描者
func scanner(p string, suffix string) []string {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		panic(err)
	}
	var fileList []string
	for _, file := range files {
		if "*" != suffix {
			if path.Ext(file.Name()) == "."+strings.Trim(suffix, ".") {
				fileList = append(fileList, p+string(os.PathSeparator)+file.Name())
			}
		} else {
			fileList = append(fileList, p+string(os.PathSeparator)+file.Name())
		}
		if file.IsDir() {
			fileList = append(fileList, scanner(p+string(os.PathSeparator)+file.Name(), suffix)...)
		}
	}
	return fileList
}

//path文件名
func Basename(path string, suffix ...string) string {
	suff := ""
	if len(suffix) > 0 {
		suff = suffix[0]
	}
	paths := strings.Split(path, PathSeparator())
	name := ufuncutil.IfStringIndex(paths[len(paths)-1:], 0)
	if suff != "" {
		names := strings.Split(name, ".")
		name = ufuncutil.IfStringIndex(names, 0)
	}
	return name
}

func Dirname(path string) string {
	paths := strings.Split(path, PathSeparator())
	paths = paths[:len(paths)-1]
	return strings.Join(paths, PathSeparator())
}

func DirBasename(path string) string {
	paths := strings.Split(path, PathSeparator())
	paths = paths[len(paths)-2 : len(paths)-1]
	return strings.Join(paths, PathSeparator())
}

//目录path切片
func Dirslice(path string) []string {
	paths := strings.Split(path, PathSeparator())
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
		SetModePerm(perm[0])
	}
	var dir []string
	for _, p := range Dirslice(dirname) {
		dir = append(dir, p)
		tmpPath := strings.Join(dir, PathSeparator())
		if _, err := os.Stat(tmpPath); err != nil {
			if os.IsNotExist(err) {
				if tmpPath != "" {
					if err := os.Mkdir(tmpPath, ModePerm()); err != nil {
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
	err := os.RemoveAll(dirname)
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
	err := os.Remove(filename)
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
	return strings.Join(paths, PathSeparator())
}

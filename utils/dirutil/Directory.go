package dirutil

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

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

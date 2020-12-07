package utils

import (
	"io/ioutil"
	"log"
	"os"
)

type Dir struct {
	Path string // 扫描完整路径
	List []string
}

func NewDir(path string) *Dir {
	return &Dir{Path: path}
}

// 扫描
func (this *Dir) Scan() *Dir {
	this.List = scanner(this.Path)
	return this
}

// 扫描者
func scanner(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	var fileList []string
	for _, file := range files {
		fileList = append(fileList, path + string(os.PathSeparator) + file.Name())
		if file.IsDir() {
			fileList = append(fileList, scanner(path + string(os.PathSeparator) + file.Name())...)
		}
	}
	return fileList
}
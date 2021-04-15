package utils

import (
	"fmt"
	"os"
	"reflect"
)

// 自定义文件结构体
type File struct {
	Name string //文件名(完整路径)
	file *os.File
}

// 构造函数
func NewFile(name string) *File {
	return &File{Name: name}
}

// 创建
func (this *File) Create() *File {
	this.file = Mkfile(this.Name)
	return this
}

// 关闭文件句柄
func (this *File) Close() *File {
	defer this.file.Close()
	return this
}

// 删除
func (this *File) Remove() *File {
	Remove(this.Name)
	return this
}

// 获取文件创建时间戳
func (this *File) CreateTimestamp() int64 {
	fileInfo, _ := os.Stat(this.Name)
	t := reflect.ValueOf(fileInfo.Sys())
	if "syscall.Win32FileAttributeData" == fmt.Sprintf("%s", t.Elem().Type()) {
		lowDateTime := t.Elem().FieldByName("CreationTime").FieldByName("LowDateTime").Uint()
		highDateTime := t.Elem().FieldByName("CreationTime").FieldByName("HighDateTime").Uint()
		return Nanoseconds(uint32(lowDateTime), uint32(highDateTime)) / 1e9
	}
	return int64(t.Elem().FieldByName("Ctim").FieldByName("Sec").Int())
}

// 获取文件更新时间戳
func (this *File) ModifyTimestamp() int64 {
	fileInfo, _ := os.Stat(this.Name)
	t := reflect.ValueOf(fileInfo.Sys())
	if "syscall.Win32FileAttributeData" == fmt.Sprintf("%s", t.Elem().Type()) {
		lowDateTime := t.Elem().FieldByName("LastWriteTime").FieldByName("LowDateTime").Uint()
		highDateTime := t.Elem().FieldByName("LastWriteTime").FieldByName("HighDateTime").Uint()
		return Nanoseconds(uint32(lowDateTime), uint32(highDateTime)) / 1e9
	}
	return int64(t.Elem().FieldByName("Mtim").FieldByName("Sec").Int())
}

// 获取文件访问时间戳
func (this *File) AccessTimestamp() int64 {
	fileInfo, _ := os.Stat(this.Name)
	t := reflect.ValueOf(fileInfo.Sys())
	if "syscall.Win32FileAttributeData" == fmt.Sprintf("%s", t.Elem().Type()) {
		lowDateTime := t.Elem().FieldByName("LastAccessTime").FieldByName("LowDateTime").Uint()
		highDateTime := t.Elem().FieldByName("LastAccessTime").FieldByName("HighDateTime").Uint()
		return Nanoseconds(uint32(lowDateTime), uint32(highDateTime)) / 1e9
	}
	return int64(t.Elem().FieldByName("Atim").FieldByName("Sec").Int())
}

// 获取文件创建时间
func (this *File) CreateTime() string {
	timestamp := this.CreateTimestamp()
	return Date(timestamp)
}

// 获取文件更新时间
func (this *File) ModifyTime() string {
	timestamp := this.ModifyTimestamp()
	return Date(timestamp)
}

// 获取文件访问时间
func (this *File) AccessTime() string {
	timestamp := this.AccessTimestamp()
	return Date(timestamp)
}

// 文件是否存在
func (this *File) Exist() bool {
	if _, err := os.Stat(this.Name); os.IsNotExist(err) {
		return false
	}
	return true
}

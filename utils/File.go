package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

var FileFlag = os.O_APPEND

// 自定义文件结构体
type File struct {
	Name string //文件名(完整路径)
	file *os.File
}

// 构造函数
func NewFile(name string) *File {
	f := &File{Name: name}
	if f.Exist() {
		file, err := os.OpenFile(f.Name, FileFlag, ModePerm)
		if err != nil {
			panic(err)
		}
		f.file = file
	} else {
		f.Create()
	}
	return f
}

// 创建
func (this *File) Create() *File {
	this.file = Mkfile(this.Name)
	return this
}

//文件句柄
func (this *File) File() *os.File {
	return this.file
}

//读取所有
func (this *File) ReadAll() []byte {
	defer this.Close()
	b, err := ioutil.ReadAll(this.file)
	if err != nil {
		panic(err)
	}
	return b
}

func (this *File) ReadBlock(filePth string, bufSize int, hookFunc func([]byte)) error {
	defer this.Close()
	buf := make([]byte, bufSize) //一次读取多少个字节
	bfRd := bufio.NewReader(this.file)
	for {
		n, err := bfRd.Read(buf)
		hookFunc(buf[:n]) // n 是成功读取字节数

		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// 关闭文件句柄
func (this *File) Close() bool {
	err := this.file.Close()
	if err != nil {
		panic(err)
	}
	return true
}

// 删除
func (this *File) Remove() *File {
	Remove(this.Name)
	return this
}

// 获取文件创建时间戳
func (this *File) CreateTimestamp() int64 {
	fileInfo, err := os.Stat(this.Name)
	if err != nil {
		panic(err)
	}
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
	fileInfo, err := os.Stat(this.Name)
	if err != nil {
		panic(err)
	}
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
	fileInfo, err := os.Stat(this.Name)
	if err != nil {
		panic(err)
	}
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

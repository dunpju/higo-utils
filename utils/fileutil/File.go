package fileutil

import (
	"bufio"
	"fmt"
	"github.com/dengpju/higo-utils/utils/dirutil"
	"github.com/dengpju/higo-utils/utils/timeutil"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

//覆盖 os.O_WRONLY | os.O_TRUNC | os.O_CREATE
//追加 os.O_APPEND
var fileFlag = os.O_APPEND

type Fileutil struct {
}

// 构造函数
func (this *Fileutil) New(name string, flag int, perm os.FileMode) *File {
	return NewFile(name, flag, perm)
}

// 读取文件
func (this *Fileutil) ReadFile(name string) *File {
	return ReadFile(name)
}

// 文件是否存在
func (this *Fileutil) FileExist(name string) bool {
	return FileExist(name)
}

type Filer interface {
	File() *os.File
	Path() string
}

// 自定义文件结构体
type File struct {
	Name      string //文件名(完整路径)
	SplitFunc bufio.SplitFunc
	MaxBuffer int
	file      *os.File
	isClose   bool
}

func (this *File) SetClose(is bool) *File {
	this.isClose = is
	return this
}

// 构造函数
func NewFile(name string, flag int, perm os.FileMode) *File {
	f := &File{Name: name, isClose: true}
	if f.Exist() {
		file, err := os.OpenFile(f.Name, flag, perm)
		if err != nil {
			panic(err)
		}
		f.file = file
	} else {
		f.Create()
	}
	return f
}

// 读取文件
func ReadFile(name string) *File {
	f := &File{Name: name, isClose: true}
	if f.Exist() {
		file, err := os.Open(f.Name)
		if err != nil {
			panic(err)
		}
		f.file = file
	} else {
		panic(fmt.Errorf("file non exist"))
	}
	return f
}

// 创建
func (this *File) Create() *File {
	this.file = dirutil.Mkfile(this.Name)
	return this
}

//文件句柄
func (this *File) File() *os.File {
	return this.file
}

//写文件
func (this *File) Write(b []byte) (n int, err error) {
	return this.file.Write(b)
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

//读取所有文件字符串
func (this *File) ReadAllString() string {
	return string(this.ReadAll())
}

//遍历
func (this *File) ForEach(callable func(line int, b []byte)) error {
	defer this.Close()
	// Splits on newlines by default.
	scanner := bufio.NewScanner(this.file)
	if this.MaxBuffer > 0 {
		buf := make([]byte, this.MaxBuffer)
		scanner.Buffer(buf, this.MaxBuffer)
	}
	if this.SplitFunc != nil {
		scanner.Split(this.SplitFunc)
	}
	l := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		callable(l, scanner.Bytes())
		l++
	}
	return scanner.Err()
}

//扫描
func (this *File) Scan(callable func(scanner *bufio.Scanner)) {
	defer this.Close()
	// Splits on newlines by default.
	scanner := bufio.NewScanner(this.file)
	if this.MaxBuffer > 0 {
		buf := make([]byte, this.MaxBuffer)
		scanner.Buffer(buf, this.MaxBuffer)
	}
	if this.SplitFunc != nil {
		scanner.Split(this.SplitFunc)
	}
	callable(scanner)
}

//扫描者
func (this *File) Scanner() *bufio.Scanner {
	// Splits on newlines by default.
	scanner := bufio.NewScanner(this.file)
	if this.MaxBuffer > 0 {
		buf := make([]byte, this.MaxBuffer)
		scanner.Buffer(buf, this.MaxBuffer)
	}
	if this.SplitFunc != nil {
		scanner.Split(this.SplitFunc)
	}
	return scanner
}

//块读取
func (this *File) ReadBlock(bufSize int, hookFunc func([]byte)) error {
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
	if this.isClose {
		err := this.file.Close()
		if err != nil {
			panic(err)
		}
		return true
	}
	return false
}

// 删除
func (this *File) Remove() *File {
	dirutil.Remove(this.Name)
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
		return timeutil.Nanoseconds(uint32(lowDateTime), uint32(highDateTime)) / 1e9
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
		return timeutil.Nanoseconds(uint32(lowDateTime), uint32(highDateTime)) / 1e9
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
		return timeutil.Nanoseconds(uint32(lowDateTime), uint32(highDateTime)) / 1e9
	}
	return int64(t.Elem().FieldByName("Atim").FieldByName("Sec").Int())
}

// 获取文件创建时间
func (this *File) CreateTime() string {
	timestamp := this.CreateTimestamp()
	return timeutil.Date(timestamp)
}

// 获取文件更新时间
func (this *File) ModifyTime() string {
	timestamp := this.ModifyTimestamp()
	return timeutil.Date(timestamp)
}

// 获取文件访问时间
func (this *File) AccessTime() string {
	timestamp := this.AccessTimestamp()
	return timeutil.Date(timestamp)
}

// 文件是否存在
func (this *File) Exist() bool {
	if _, err := os.Stat(this.Name); os.IsNotExist(err) {
		return false
	}
	return true
}

// 文件是否是目录
func (this *File) IsDir() bool {
	fi, e := os.Stat(this.Name)
	if e != nil {
		return false
	}
	return fi.IsDir()
}

// 文件大小
func (this *File) Size() (int64, error) {
	f, err := os.Stat(this.Name)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

// 文件是否存在
func FileExist(name string) bool {
	f := &File{Name: name}
	return f.Exist()
}

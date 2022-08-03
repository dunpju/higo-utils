package Slice

import (
	"fmt"
	"github.com/dengpju/higo-utils/utils"
	"log"
	"strings"
)

type GenericLogger[T any] interface {
	WithField(a T, b T) T
	Info(T)
}

type MyLogger struct {
	fields []string
}

func (m *MyLogger) WithField(k string, v string) *MyLogger {
	m.fields = append(m.fields, k+"="+v)
	return m
}

func (m *MyLogger) Info(msg string) {
	log.Printf("%s : %s", strings.Join(m.fields, ","), msg)
}

func DoStuff[T GenericLogger[T]](t T) {
	//t.WithField("go", "1.18").Info("is awesome")
}

type MyIntLogger struct {
	fields []int
}

func (this *MyIntLogger) WithField(k, v int) *MyIntLogger {
	this.fields = append(this.fields, k, v)
	return this
}

func (this *MyIntLogger) Info(s int) {
	log.Printf("%d : %d ", this.fields, s)
}

func Test_slice() {
	// T
	//sliceutil.DoStuff(&sliceutil.MyLogger{})
	mylog := &MyLogger{}
	mylog.WithField("ff", "hh").Info("hh")
	(&MyIntLogger{}).WithField(1, 2).Info(3)

	// Slice
	sl := utils.Slice.New()
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.Append("2")
	sl.Insert(0, "1")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.Append("3")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) bool {
		fmt.Println(index, value)
		return true
	})
	sl.Insert(1, "11")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) bool {
		fmt.Println(index, value)
		return true
	})
	sl.Insert(3, "114")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) bool {
		fmt.Println(index, value)
		return true
	})
	sl.Remove("114")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) bool {
		fmt.Println(index, value)
		return true
	})
	sl.Delete(1)
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) bool {
		fmt.Println(index, value)
		return true
	})
	sl.Replace("1", "11")
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) bool {
		fmt.Println(index, value)
		return true
	})
	sl.Reverse()
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) bool {
		fmt.Println(index, value)
		return true
	})
	fmt.Println(sl.Exist("3"))
	fmt.Println(sl.Value()[2:])
	sl.Empty()
	fmt.Printf("%p\n", sl)
	fmt.Println(sl)
	sl.ForEach(func(index int, value interface{}) bool {
		fmt.Println(index, value)
		return true
	})
}

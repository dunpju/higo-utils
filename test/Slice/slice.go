package Slice

import (
	"fmt"
	"github.com/dengpju/higo-utils/utils"
	"log"
)

type GenericLogger[T any] interface {
	WithField(T, T) T
}

type MyLogger[T any] struct {
	fields []T
}

func (m *MyLogger[T]) WithField(k, v T) *MyLogger[T] {
	m.fields = append(m.fields, k, v)
	fmt.Println(m)
	return m
}

func DoStuff[T GenericLogger[T]](t T) {
	//t.WithField("go", "1.18")
	//fmt.Println(t)
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

type DataProcessor[T any] interface {
	Process(oriData T) (newData T)
	Save(data T) error
}

type DataProcessor2[T any] interface {
	int | ~struct{ Data interface{} }

	Process(data T) (newData T)
	Save(data T) error
}

type CSVProcessor struct {
}

// 注意，方法中 oriData 等的类型是 string
func (c *CSVProcessor) Process(oriData string) (newData string) {
	return oriData
}

func (c *CSVProcessor) Save(oriData string) error {
	return nil
}

func tt[T DataProcessor[string]](t T) {
	t.Process("gg")
}

func Test_slice() {

	csv := &CSVProcessor{}
	csv.Process("ggg")
	csv.Save("hhh")

	// T
	mylog := &MyLogger[int]{}
	//mylog.WithField("ff", "hh").Info("hh")
	//(&MyIntLogger{}).WithField(1, 2).Info(3)
	//DoStuff(&MyLogger{})
	mylog.WithField(10, 20)
	(&MyLogger[string]{}).WithField("ff", "gg")
	return
	//DoStuff(&MyIntLogger{})

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

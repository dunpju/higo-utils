package utils

import (
	"encoding/json"
	"strconv"
)

// map
type Map struct {
	Map1 map[string]interface{} // map1
	Map2 map[string]interface{} // map2
}

// 构造方法
func NewMap(map1 map[string]interface{}, map2 map[string]interface{}) *Map {
	return &Map{map1, map2}
}

// 合并map2到map1
func (this *Map) Merge() *map[string]interface{} {
	for k, v := range this.Map1 {
		this.Map2[k] = v
	}
	return &this.Map1
}

type IMap interface {
	Put(key string, value interface{}) IMap
	Push(value interface{}) IMap
	Shift() interface{}
	Pop() interface{}
	Current() interface{}
	End() interface{}
	Get(key string) interface{}
	Exist(key string) bool
	Replace(key string, value interface{}) bool
	Remove(key string) bool
	ForEach(callable Callable)
	Clear() bool
	Len() int
	Merge(M IMap)
	String() string
}

type ArrayMap struct {
	sort  []string
	value map[string]interface{}
}

type MapInt ArrayMap

func Array() *ArrayMap {
	return &ArrayMap{sort: make([]string, 0), value: make(map[string]interface{}, 0)}
}

// 添加元素
func (this *ArrayMap) Put(key string, value interface{}) IMap {
	if _, ok := this.value[key]; ok {
		this.Replace(key, value)
	} else {
		this.value[key] = value
		this.sort = append(this.sort, key)
	}
	return this
}

func (this *ArrayMap) Push(value interface{}) IMap {
	this.Put(strconv.Itoa(len(this.sort)), value)
	return this
}

//弹出第一个元素
func (this *ArrayMap) Shift() interface{} {
	if len(this.sort) > 0 {
		key := this.sort[0]
		v := this.value[key]
		this.sort = append(this.sort[:0], this.sort[0+1:]...)
		delete(this.value, key)
		return map[string]interface{}{key: v}
	}
	return nil
}

func (this *ArrayMap) Pop() interface{} {
	length := len(this.sort)
	if length > 0 {
		key := this.sort[length-1]
		v := this.value[key]
		this.sort = this.sort[:len(this.sort)-1]
		delete(this.value, key)
		return map[string]interface{}{key: v}
	}
	return nil
}

func (this *ArrayMap) Current() interface{} {
	if len(this.sort) > 0 {
		key := this.sort[0]
		v := this.value[key]
		return map[string]interface{}{key: v}
	}
	return nil
}

func (this *ArrayMap) End() interface{} {
	length := len(this.sort)
	if length > 0 {
		key := this.sort[length-1]
		v := this.value[key]
		return map[string]interface{}{key: v}
	}
	return nil
}

// 修改元素
func (this *ArrayMap) Replace(key string, value interface{}) bool {
	if _, ok := this.value[key]; ok {
		this.value[key] = value
	} else {
		return false
	}
	return true
}

// 删除元素
func (this *ArrayMap) Remove(key string) bool {
	if _, ok := this.value[key]; ok {
		delete(this.value, key)
	} else {
		return false
	}
	return true
}

// 清除map
func (this *ArrayMap) Clear() bool {
	this.value = make(map[string]interface{}, 0)
	this.sort = this.sort[0:0]
	return true
}

// 长度
func (this *ArrayMap) Len() int {
	return len(this.sort)
}

type Callable func(key string, value interface{})

// 遍历元素
func (this *ArrayMap) ForEach(callable Callable) {
	for _, key := range this.sort {
		callable(key, this.value[key])
	}
}

// 查询元素
func (this *ArrayMap) Get(key string) interface{} {
	if value, ok := this.value[key]; ok {
		return value
	} else {
		panic("`" + key + "` The key doesn't exist in the map")
	}
}

// 元素是否存在
func (this *ArrayMap) Exist(key string) bool {
	if _, ok := this.value[key]; ok {
		return true
	}
	return false
}

// 合并
func (this *ArrayMap) Merge(m IMap) {
	m.ForEach(func(key string, value interface{}) {
		this.Put(key, value)
	})
}

func (this *ArrayMap) String() string {
	mjson, err := json.Marshal(this.value)
	if err != nil {
		panic(err)
	}
	return string(mjson)
}

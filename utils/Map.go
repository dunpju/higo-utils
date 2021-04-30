package utils

import (
	"encoding/json"
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
	Put(key string, value interface{}) MapString
	Get(key string) interface{}
	Exist(key string) bool
	Replace(key string, value interface{}) bool
	Remove(key string) bool
	ForEach(callable Callable)
	Clear() bool
	Len() int
	Merge(M MapString)
	String() string
}

type MapString map[string]interface{}

type MapInt MapString

func MapOperation(m MapString) MapString {
	return m
}

// 添加元素
func (this MapString) Put(key string, value interface{}) MapString {
	this[key] = value
	return this
}

// 修改元素
func (this MapString) Replace(key string, value interface{}) bool {
	if _, ok := this[key]; ok {
		this[key] = value
	} else {
		return false
	}
	return true
}

// 删除元素
func (this MapString) Remove(key string) bool {
	if _, ok := this[key]; ok {
		delete(this, key)
	} else {
		return false
	}
	return true
}

// 清除map
func (this MapString) Clear() bool {
	for key, _ := range this {
		delete(this, key)
	}
	return true
}

// 长度
func (this MapString) Len() int {
	return len(this)
}

type Callable func(key string, value interface{})

// 遍历元素
func (this MapString) ForEach(callable Callable) {
	for key, value := range this {
		callable(key, value)
	}
}

// 查询元素
func (this MapString) Get(key string) interface{} {
	if value, ok := this[key]; ok {
		return value
	} else {
		panic("`" + key + "` The key doesn't exist in the map")
	}
}

// 元素是否存在
func (this MapString) Exist(key string) bool {
	if _, ok := this[key]; ok {
		return true
	}
	return false
}

// 合并
func (this MapString) Merge(M MapString) {
	M.ForEach(func(key string, value interface{}) {
		this.Put(key, value)
	})
}

func (this MapString) String() string {
	mjson, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	return string(mjson)
}

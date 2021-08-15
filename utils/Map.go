package utils

import (
	"encoding/json"
	"strconv"
	"sync"
)

type IMap interface {
	Put(key string, value interface{}) IMap
	Push(value interface{}) IMap
	Shift() interface{}
	Pop() interface{}
	Current() interface{}
	End() interface{}
	Get(key string) interface{}
	Value() interface{}
	Exist(key string) bool
	Replace(key string, value interface{}) bool
	Remove(key string) bool
	ForEach(callable Callable)
	Clear() bool
	Len() int
	Merge(m IMap)
	String() string
}

type KeyValue struct {
	Key   interface{}
	Value interface{}
}

func NewKeyValue(key interface{}, value interface{}) *KeyValue {
	return &KeyValue{Key: key, Value: value}
}

type ArrayMap struct {
	index int64
	sort  []string
	value map[string]interface{}
	lock  sync.Mutex
}

func Array() *ArrayMap {
	return &ArrayMap{sort: make([]string, 0), value: make(map[string]interface{}, 0)}
}

func (this *ArrayMap) Put(key string, value interface{}) IMap {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.value[key]; ok {
		this.Replace(key, value)
	} else {
		this.value[key] = value
		this.sort = append(this.sort, key)
	}
	this.index++
	return this
}

func (this *ArrayMap) Push(value interface{}) IMap {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.Put(strconv.FormatInt(this.index, 10), value)
	return this
}

func (this *ArrayMap) Shift() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if len(this.sort) > 0 {
		key := this.sort[0]
		v := this.value[key]
		this.sort = append(this.sort[:0], this.sort[0+1:]...)
		delete(this.value, key)
		return NewKeyValue(key, v)
	}
	return nil
}

func (this *ArrayMap) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	length := len(this.sort)
	if length > 0 {
		key := this.sort[length-1]
		v := this.value[key]
		this.sort = this.sort[:len(this.sort)-1]
		delete(this.value, key)
		return NewKeyValue(key, v)
	}
	return nil
}

func (this *ArrayMap) Current() interface{} {
	if len(this.sort) > 0 {
		key := this.sort[0]
		v := this.value[key]
		return NewKeyValue(key, v)
	}
	return nil
}

func (this *ArrayMap) End() interface{} {
	length := len(this.sort)
	if length > 0 {
		key := this.sort[length-1]
		v := this.value[key]
		return NewKeyValue(key, v)
	}
	return nil
}

func (this *ArrayMap) Replace(key string, value interface{}) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.value[key]; ok {
		this.value[key] = value
	} else {
		return false
	}
	return true
}

func (this *ArrayMap) Remove(key string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.value[key]; ok {
		tmp := make([]string, 0)
		for _, v := range this.sort {
			if v != key {
				tmp = append(tmp, v)
			}
		}
		this.sort = tmp
		delete(this.value, key)
	} else {
		return false
	}
	return true
}

func (this *ArrayMap) Clear() bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.value = make(map[string]interface{}, 0)
	this.sort = this.sort[0:0]
	return true
}

func (this *ArrayMap) Len() int {
	return len(this.sort)
}

type Callable func(key string, value interface{})

func (this *ArrayMap) ForEach(callable Callable) {
	for _, key := range this.sort {
		callable(key, this.value[key])
	}
}

func (this *ArrayMap) Get(key string) interface{} {
	if value, ok := this.value[key]; ok {
		return value
	} else {
		panic("`" + key + "` The key doesn't exist in the map")
	}
}

func (this *ArrayMap) Value() interface{} {
	return this.value
}

func (this *ArrayMap) Exist(key string) bool {
	if _, ok := this.value[key]; ok {
		return true
	}
	return false
}

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

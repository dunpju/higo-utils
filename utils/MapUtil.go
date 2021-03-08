package utils

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

type MapString map[string]interface{}

func MapOperation(m MapString) MapString {
	return m
}

func (this MapString) Append(key string, value interface{}) MapString {
	this[key] = value
	return this
}

func MapAppend(m map[string]interface{}, key string, value interface{}) map[string]interface{} {
	m[key] = value
	return m
}
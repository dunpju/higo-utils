package main

import (
	"fmt"
	"github.com/dengpju/higo-utils/utils"
)

func main() {
	fmt.Println(utils.Dir("./utils").Scan().Suffix("go").Get())
	result := utils.MapOperation(make(map[string]interface{})).Put("key1", "value1").Put("key2", "value2")
	fmt.Println(result)
	result.ForEach(func(key string, value interface{}) {
		fmt.Println(key, value)
	})
	result.Replace("key1", "k")
	fmt.Println(result.Get("key1"))
	result.Remove("k")
	result.Remove("key2")
	fmt.Println(result.Len())
	fmt.Println(result)
	result.Clear()
	result.ForEach(func(key string, value interface{}) {
		fmt.Println(key, value, "for")
	})

}

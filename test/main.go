package main

import (
	"fmt"
	"github.com/dengpju/higo-utils/utils"
)

func main() {
	fmt.Println(utils.Dir("./utils").Scan().Suffix("go").Get())
	result := utils.MapOperation(make(map[string]interface{})).Append("key1", "value1").Append("key2", "value2")
	fmt.Println(result)
}

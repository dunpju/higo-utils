package main

import (
	"fmt"
	"github.com/dengpju/higo-utils/utils"
)

func main()  {
	fmt.Println(utils.Dir("./utils").Suffix("go").Scan().List())
}

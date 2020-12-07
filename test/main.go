package main

import (
	"fmt"
	"github.com/dengpju/higo-utils/utils"
)

func main()  {
	d := utils.NewDir("./utils")
	fmt.Println(d.Scan().List)
}

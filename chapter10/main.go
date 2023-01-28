package main

import (
	"config"
	"fmt"
)

func main() {
	c, err := config.New()
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(c)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	config := Config{}
	arguments := os.Args
	config.name = arguments[1]
	config.sex = arguments[2]
	config.age = arguments[3]

	fmt.Println(arguments[1])
	fmt.Println(arguments[2])
	fmt.Println(arguments[3])
}

package main

import "fmt"

func main() {
	var i interface{}
	i=1
	fmt.Println(i)
	i=3.123123
	fmt.Println(i)
}

type Dync struct {
	getscret string
}

func (d *Dync) getscrent() string {
	return "a"
}

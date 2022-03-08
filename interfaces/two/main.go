package main

import "fmt"

func main() {
	security := Assets{assets: []Asset{
		&GlassDoor{},
		&WoodDoor{},
	}}
	fmt.Println("开始上班")
	security.DoStartWrok()
	fmt.Println("八小时后")
	security.DoStopWrok()
	fmt.Println("Done")
}

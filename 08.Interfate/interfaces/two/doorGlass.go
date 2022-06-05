package main

import "fmt"

//如果没有完全实现door的方法，则会报错
//必须要实现接口中的所有方法，否则报错
var _ Door = &GlassDoor{}

type GlassDoor struct {
}

func (g *GlassDoor) Open() {
	fmt.Println("GlassDoor open")
}
func (g *GlassDoor) Close() {
	fmt.Println("GlassDoor close")
}
func (g *GlassDoor) Unlock() {
	fmt.Println("GlassDoor Unlock")
}
func (g *GlassDoor) Lock() {
	fmt.Println("GlassDoor lock")
}

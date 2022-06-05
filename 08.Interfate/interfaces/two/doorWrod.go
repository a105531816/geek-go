package main

import "fmt"

var _ Door = &GlassDoor{}
type WoodDoor struct {
}

func (g *WoodDoor) Open() {
	fmt.Println("WoodDoor open")
}
func (g *WoodDoor) Close() {
	fmt.Println("WoodDoor close")
}
func (g *WoodDoor) Unlock() {
	fmt.Println("WoodDoor Unlock")
}
func (g *WoodDoor) Lock() {
	fmt.Println("WoodDoor lock")
}

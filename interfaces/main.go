package main

import (
	"fmt"
)

func main() {
	var elephant Elephant
	var refrigerator Refrigerator

	//fmt.Println(elephant.Name, refrigerator.Size)

	//var pute PutElephantIntoRefrigerator
	//pute.OpenTheDoorOfRefrigerator(refrigerator)
	//pute.PutElephantIntoRefrigerator(elephant, refrigerator)
	//pute.CloseTheDoorOfRefrigerator(refrigerator)

	var legend PutElephantIntoRefrigerator = PutElephantIntoRefrigeratorImpl{}
	legend.OpenTheDoorOfRefrigerator(refrigerator)
	legend.PutElephantIntoRefrigerator(elephant, refrigerator)
	legend.CloseTheDoorOfRefrigerator(refrigerator)
	// todo show the elephant in refrigerator

	var i interface{}
	i = 1
	i = func() { fmt.Println(1) }
	fmt.Println(i)

}

type PutElephantIntoRefrigerator interface {
	OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error
	PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error
	CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error
}

type Refrigerator struct {
	Size int
}
type Elephant struct {
	Name string
}

type PutElephantIntoRefrigeratorImpl struct {
}

func (c PutElephantIntoRefrigeratorImpl) OpenTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("打开冰箱门")
	return nil
}

func (c PutElephantIntoRefrigeratorImpl) PutElephantIntoRefrigerator(elephant Elephant, refrigerator Refrigerator) error {
	// todo
	fmt.Println("装进去")
	return nil
}

func (c PutElephantIntoRefrigeratorImpl) CloseTheDoorOfRefrigerator(refrigerator Refrigerator) error {
	// todo
	fmt.Println("关闭冰箱门")
	return nil
}

type Open interface {
	Open() error
}

type Close interface {
	Close() error
}

type Box interface {
	Open
	Close
}

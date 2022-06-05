package main

import "fmt"

func main() {
	getNumNum()
	fmt.Println(".....")
}

//
func getNum() func(a int) int {
	c := 1

	//当函数被调用时，返回一个函数，此时函数的作用域在getNum9()；
	//返回的函数可以使用自身周边的变量。
	return func(a int) int {
		a = c + a
		fmt.Printf("闭包，%v\n", a)
		return a
	}
}

func getNumNum() {
	imp := getNum()
	imp(1)
	imp(2)
	imp(3)
	imp(4)
	imp(5)
}

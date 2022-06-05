package main

import "fmt"

var (
	//如果想要成为全局匿名函数就必须赋值给一个全局变量
	//Func1 是一个全局匿名函数
	Func1 = func(n1, n2 int) int {
		return n1 + n2
	}
)

func main() {
	//在main函数中调用全局匿名函数
	res4 := Func1(4, 9)
	fmt.Println(res4)
}

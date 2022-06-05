package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	test()                       //会报错，但是会被defer+ recover拦截
	fmt.Println("yes1231231231") //不会影响进度
	panic("发生严重错误")
}

func test() {
	defer funcName()
	defer funcName1()

	//此方式无法捕获panic，因为recover()脱离了原作用域
	defer func() {
		funcName()
	}()
	num1 := 10
	num2 := 0
	num3 := num1 / num2

	fmt.Println(num3) //报错，分母为0
}

func funcName() {
	//可以在if中声明变量并赋值，并判断
	if err := recover(); err != nil {
		fmt.Println("有错误:", err)
		debug.PrintStack() //打印错误堆栈
	}
}

func funcName1() {
	//使用匿名函数的方式进行错误捕获，会导致无法捕获错误
	//匿名函数的捕获房子导致recover()的作用域被限制，无法捕获全局的panic
	func() {
		if err := recover(); err != nil {
			fmt.Println("有错误:", err)
			debug.PrintStack() //打印错误堆栈
		}
	}()
}

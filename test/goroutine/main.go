package main

import (
	"fmt"
	"strconv"
	"time"
)

//○ 在主线程（进程）中，开启一个goroutine，该协程每隔一秒输出“hello world”
//○ 在主线程中也每隔一秒输出“hello world”，输出10次后，退出程序
//○ 要求主线程和goroutine同时执行

//编写一个函数，每间隔一秒输出helloword

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello wolld" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
func main() {
	for j := 0; j < 10; j++ {
		go test()
	}

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello wolld" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

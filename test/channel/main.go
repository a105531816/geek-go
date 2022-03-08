package main

import "fmt"

func main() {
	var intChan chan int
	//channel必须初始化才可以使用
	intChan = make(chan int, 3)
	intChan1 := make(chan int, 3)
	//channel是引用类型，其底层是一个地址
	fmt.Printf("channel 的值：%v\n", intChan)
	fmt.Printf("channel 本身的地址%p\n", intChan)
	for i := 0; i < 2; i++ {
		//往channel中写入数据
		intChan <- i
	}
	fmt.Printf("channel 的长度：%v，channel 的容量：%v\n", len(intChan), cap(intChan))
	//取channel中的数据，如果取的数据个数大于channel中存放数据的个数，就会报错deadlock
	a := <-intChan
	fmt.Println(a)

	select {
	case c := <-intChan:
		fmt.Println("intChan准备好，可以读取！", c)
	case c := <-intChan1:
		fmt.Println("intChan1准备好，可以读取！", c)
	default:
		fmt.Println("intChan1 和intChan都未准备好，不可以读取！")
	}
}

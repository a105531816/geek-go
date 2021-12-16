package main

import (
	"fmt"
	"sync"
)

//现在需要计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中

//编写一个函数，用来计算数值的阶乘
//我们启动的协程是多个，统一将结果放入的map中
//map是全局

var myMap = make(map[uint64]uint64, 300)
var (
	//声明一个全局互斥锁
	//lock 是一个全局的互斥锁
	//sync包
	//Mutex：是x锁互斥锁
	lock sync.Mutex
)

func jiecheng(num int) {
	var a uint64 = 1
	for i := 1; i <= num; i++ {
		a *= uint64(i)
	}
	lock.Lock()
	myMap[uint64(num)] = uint64(a)
	lock.Unlock()
}
func main() {
	for i := 1; i <= 200; i++ {
		go jiecheng(i)
	}
	//time.Sleep(2 * time.Second)

	for i1, i2 := range myMap {
		fmt.Printf("%v的阶乘是：%v\n", i1, i2)
	}

}

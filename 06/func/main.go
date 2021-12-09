package main

import "fmt"

func main() {
	getMonitor(1, "男", setMonitor) //传输参数，使用setMonitor这个函数执行
}

//恰好getMointor对这两个参数也有用处，可以直接拿来用，传参个数相同
func getMonitor(id int, name string, geta func(id int, name string)) {
	geta(id, name) //其实相当于直接调用了回调函数
}

//函数setMonitor传入两个参数，进行判断，输出
func setMonitor(id int, name string) {
	if id == 1 && name == "男" {
		id = 2
	}
	fmt.Println("a")
}

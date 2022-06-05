package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	//设置最大使用的CPU核心数（预留一个）
	runtime.GOMAXPROCS(cpuNum - 1)
}

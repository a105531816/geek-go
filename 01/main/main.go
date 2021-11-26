package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello\tworld\n!")   // \t ：制表位，实现对齐功能
	fmt.Println("hello\nworld\n!")   // \n ：换行符
	fmt.Println("hello\\ world\\ !") // \\ ：一个\
	fmt.Println("hello\rworld\r!")   // \r ：一个回车，会从当前行的最前面开始输出，替换掉以前的版本。
}

package main

//计算两个圆的面积和

import "fmt"

//声明常量π
const (
	pi float32 = 3.1415926
)

func main() {
	var r1 float32 = 11
	var r2 float32 = 15
	var sum2 = (pi * r1 * r1) + (pi * r2 * r2)
	fmt.Printf("两个圆的面积和为：%.3f", sum2)
}

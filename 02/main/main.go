package main

//创建一个数组，反转数组

import "fmt"

func main() {
	var arry1 = [4]int{}
	var arry2 = [4]int{}
	c := len(arry1) - 1
	d := 0
	arry1[0] = 100
	arry1[1] = 120
	arry1[2] = 233
	arry1[3] = 424
	for _, i := range arry1 {
		arry2[c] = i
		c--
	}
	for _, i := range arry2 {
		fmt.Printf("数组B的第%v个元素的值为：%v\n", d, i)
		d++
	}
}

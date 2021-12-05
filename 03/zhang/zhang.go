package main

import (
	"fmt"
)

func main() {
	arry1 := [12][]int{}
	var NewYear int = 4
	var run bool
	//fmt.Println("请输入您的年龄：")
	//fmt.Scanln(&NewYear)
	//判断是否是闰年，能被400整除，能被4整除，但是不能被100整除
	if NewYear%400 == 0 || NewYear%4 == 0 {
		if NewYear%100 != 0 {
			run = true
		}
	}

	for i := 0; i <= 11; i++ {
		if i == 2 && run {
			for j := 1; j <= 29; j++ {
				var day = 0
				arry1[i][day] = j
				day++
			}
			continue
		} else if i == 1 || i == 3 || i == 5 || i == 7 || i == 8 || i == 10 || i == 12 {
			for j := 1; j <= 31; j++ {
				var day = 0

				arry1[i][day] = j
				day++
			}
		} else if i == 2 {
			for j := 1; j <= 29; j++ {
				var day = 0
				arry1[i][day] = j
				day++
			}
		} else {
			for j := 1; j <= 30; j++ {
				var day = 0
				fmt.Println(i, day, j)
				arry1[i][day] = j
				day++
			}
		}
	}
	fmt.Println(arry1)
}

package main

import "fmt"

//用map管理20人的人数
//算出平均分
//根据分数高低对20人排名，高分在前，同分在同一行
//判断真map或者假map

func main() {
	list := map[string]int{
		"张1": 1, "张2": 12, "张3": 13, "张4": 41, "张5": 51,
		"张6": 71, "张7": 91, "张8": 17, "张9": 13, "张10": 12,
		"张11": 13, "张12": 14, "张13": 31, "张14": 17, "张15": 12,
		"张16": 21, "张17": 14, "张18": 41, "张19": 19, "张20": 18,
	}
	var NewMap map[string][]int
	var a string
	var b int
	var c []string
	var d []int
	var e int
	for i1, i2 := range list {
		if i2 > b {
			d[e] = b
			c[e] = i1
			b = i2
		}
	}
}

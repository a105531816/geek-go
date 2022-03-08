package main

import "fmt"

func main() {
	//引用类型
	var a map[string]string //定义一个空的map
	//var a = make(map[string]string, 10) //传统map的定义方式，必须指定实例
	a["name"] = "张"        //在空的map中输入值
	fmt.Println(a["name"]) //报空map错误

	//特殊的引用类型-slice
	var b []string         //定义一个空的slice
	b[0] = "张"             //在空的slice中输入值,报错
	b = append(b, "adsfa") //append不会报错
	fmt.Println(b)

}

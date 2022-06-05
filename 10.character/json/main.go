package main

import (
	"encoding/json"
	"fmt"
	"geek/pkg-bak/apis"
	_ "gopkg.in/yaml.v3"
	"log"
)

func main() {
	zhangshao := apis.Person{
		Name:   "张少",
		Sex:    "男",
		Age:    24,
		Weight: 75,
		Height: 170,
	}
	yam
	jszhangshao, err := json.Marshal(zhangshao)
	if err != nil {
		log.Fatal("转换失败")
		return
	}
	fmt.Println("JSON未使用string：", jszhangshao)
	fmt.Println("JSON使用strin g：", string(jszhangshao))
	var data []byte
	json.Unmarshal(data, &zhangshao)
	fmt.Println("JSON未使用string：", zhangshao)
}

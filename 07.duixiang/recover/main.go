package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime/debug"
)

type Config struct {
	name string
	sex  string
	age  int
}

func main() {
	config := Config{}
	cmds := &cobra.Command{
		Use:   "如何使用",
		Short: "体脂计算器",
		Long:  "体脂计算器基于xxxx",
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("name", config.name)
			//fmt.Println("sex", config.sex)
			//fmt.Println("age", config.age)
		},
	}
	cmds.Flags().StringVar(&config.name, "name", "张少", "姓名")
	cmds.Flags().StringVar(&config.sex, "sex", "男", "性别")
	cmds.Flags().IntVar(&config.age, "age", 10, "年龄")
	cmds.Execute()
	fmt.Println(config.name, config.sex, config.age)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(111)
			debug.PrintStack()
		}
	}()
}


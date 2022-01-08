package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Config struct {
	name string
	sex  string
	age  int
}

func main() {
	config := Config{}
	goArgs(config)
}

func goArgs(config Config) {
	cmds := &cobra.Command{
		Use:   "如何使用",
		Short: "体脂计算器",
		Long:  "体脂计算器基于xxxx",
		Run: func(cmd *cobra.Command, args []string) { //此处写要执行的东西
			fmt.Println("name", config.name)
			fmt.Println("sex", config.sex)
			fmt.Println("age", config.age)
		},
	}
	cmds.Flags().StringVar(&config.name, "name", "", "姓名")
	cmds.Flags().StringVar(&config.sex, "sex", "", "性别")
	cmds.Flags().IntVar(&config.age, "age", 1, "年龄")
	cmds.Execute() //使帮助生效
}

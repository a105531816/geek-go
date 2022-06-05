package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"os"
)

func main() {
	person := []Person{{
		name:   "张",
		sex:    "男",
		tall:   1.70,
		weight: 70,
		age:    25,
	},
	}
	for _, i2 := range person {
		bmi, err := gobmi.BMI(i2.weight, i2.tall)
		if err != nil {
			fmt.Printf("计算出错，错误为%v\n", err)
			os.Exit(10)
		}
		fmt.Println(bmi)
		fmt.Println(i2.TestPerson())
		i2.aMap = make(map[string]string)
		i2.aMap["a"] = "f"
		fmt.Println(i2.aMap)
	}

}

type Person struct {
	name   string
	sex    string
	tall   float64
	weight float64
	age    int
	aMap   map[string]string
}

func (p *Person) TestPerson() (res int) {
	if p.tall == 1.70 {
		return 1
	}
	return 0
}

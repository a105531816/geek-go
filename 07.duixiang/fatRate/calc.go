package main

import (
	gobmi "geek/pkg/go-bmi"
	"log"
)

type Calc struct {
}

//如果没有成员变量，则不需要使用(c Calc)这种方式
func (Calc) BMI(person *Person) error {
	bmi, err := gobmi.BMI(person.weight, person.tall)
	if err != nil {
		log.Println("计算BMI错误！")
		return err
	}
	person.bmi = bmi
	return nil
}

func (Calc) FatRate(person *Person) error {
	fatRate := gobmi.CalcFatRate(person.bmi, person.age, person.sex)
	person.fatRate = fatRate
	return nil
}

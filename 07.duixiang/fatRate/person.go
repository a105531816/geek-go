package main

import (
	gobmi "geek/pkg/go-bmi"
	"log"
)

type Person struct {
	name    string
	sex     string
	tall    float64
	weight  float64
	age     int
	bmi     float64
	fatRate float64
}

func (p *Person) calcBMI() error {
	bmi, err := gobmi.BMI(p.weight, p.tall)
	if err != nil {
		log.Println("输出错误")
		return err
	}
	p.bmi = bmi
	return nil
}

func (p *Person) calcFatRate() error {
	p.fatRate = gobmi.CalcFatRate(p.bmi, p.age, p.sex)
	return nil
}

package calc

import (
	"fmt"
	"log"
	"testing"
)

//测试身高负数
func TestCalc01(t *testing.T) {
	var tall float64 = -1
	var weight float64 = 70
	bmi, err := CalcBMI(tall, weight)
	if err != nil {
		log.Printf("当身高为负数时，会报错：%v\n", err)
	} else {
		log.Fatalf("当身高为负数时，没有报错\n")
	}
	fmt.Println(bmi)
}

//测试身高为0
func TestCalc02(t *testing.T) {

	var tall float64 = 0
	var weight float64 = 70
	bmi, err := CalcBMI(tall, weight)
	if err != nil {
		log.Printf("当身高为0时，会报错：%v\n", err)
	} else {
		log.Fatalf("当身高为0时，没有报错\n")
	}
	fmt.Println(bmi)
}

//测试体重负数
func TestCalc03(t *testing.T) {
	var tall float64 = 1.70
	var weight float64 = -1
	bmi, err := CalcBMI(tall, weight)
	if err != nil {
		log.Printf("当体重为负数时，会报错：%v\n", err)
	} else {
		log.Fatalf("当体重为负数时，没有报错\n")
	}
	fmt.Println(bmi)
}

//测试体重为0
func TestCalc04(t *testing.T) {
	var tall float64 = 1.70
	var weight float64 = 0
	bmi, err := CalcBMI(tall, weight)
	if err != nil {
		log.Printf("当体重为0时，会报错：%v\n", err)
	} else {
		log.Fatalf("当体重为0时，没有报错\n")
	}
	fmt.Println(bmi)
}

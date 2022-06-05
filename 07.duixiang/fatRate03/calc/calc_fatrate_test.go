package calc

import (
	"fmt"
	"log"
	"testing"
)

//测试年龄相关
func TestCalcFatRate(t *testing.T) {
	_, err := CalcFatRate(0, 25, "男")
	fmt.Println(err)
	if err != nil {
		log.Printf("BMI不可为0，%v\n", err)
	}
}

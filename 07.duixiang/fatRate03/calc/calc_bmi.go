package calc

import (
	"fmt"
	"github.com/pkg/errors"
)

func CalcBMI(tall float64, weight float64) (bmi float64, err error) {

	if tall <= 0 {
		fmt.Println("身高不能为0或者负数！")
		return 0, errors.New("身高不能为0或者负数！")
	}
	if weight <= 0 {
		fmt.Println("体重不能为0或者负数！")
		return 0, errors.New("体重不能为0或者负数！")
	}
	return weight / (tall * tall), nil
}

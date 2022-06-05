package test

import "errors"

//定义一个函数，返回一个error类型的err变量
func Test(num1, num2 int) (err error) {
	//判断
	if num2 != 0 {
		return nil
	} else {
		//如果分母为0，就返回一个错误
		return errors.New("分母为0")
	}
}

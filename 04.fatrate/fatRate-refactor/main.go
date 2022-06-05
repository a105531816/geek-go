package main

import (
	"fmt"
)

func init() {
	fmt.Println("init 函数")
}

func main() {
	for {
		mainFatRatBody()
		if res := whetherContinue(); !res {
			break
		}
	}

}





//判断是否就继续录入
func whetherContinue() bool {
	var whetherContinue string
	fmt.Println("是否录入下一个 (y/n)？")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}

func mainFatRatBody() {
	weight, tall, age, sexWeight, sex := getMaterialsFromInput()

	fatRate := caclFatRate(weight, tall, age, sexWeight)

	//将变量声明为方法
	var checkHealthinessFunc func(age int, fatRate float64)
	if sex == "男" {
		checkHealthinessFunc = getHealthinessSuggestionsForMale

	} else {
		checkHealthinessFunc = getHealthinessSuggestionsForFemale
	}

	//可以直接使用这个变量。和函数的使用方式一样
	checkHealthinessFunc(age, fatRate)

	//将计算方法提取成工具，进行回调
	//HealthinessSuggestions(sex, age, fatRate, getHealthinessSuggestionsForMale)
	//HealthinessSuggestions(sex, age, fatRate, getHealthinessSuggestionsForFemale)
}

//将函数以参数的方式传进去
func HealthinessSuggestions(sex string, age int, fatRate float64, getSuggestion func(age int, fatRate float64)) {
	if sex == "男" {
		//  编写男性的体脂率与体脂状态表i
		//getHealthinessSuggestionsForMale(age, fatRate)
		getSuggestion(age, fatRate)
	} else {
		//getHealthinessSuggestionsForMale(age, fatRate)
		getSuggestion(age, fatRate)
	}
}

//男生健康指标
func getHealthinessSuggestionsForMale(age int, fatRate float64) {
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 {
			fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("目前是：肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 40 && age <= 59 {
		// todo
	} else if age >= 60 {
		// todo
	} else {
		fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
	}
}

func getHealthinessSuggestionsForFemale(age int, fatRate float64) {
	if age >= 18 && age <= 39 {
		if fatRate <= 0.1 {
			fmt.Println("目前是：偏瘦。要多吃多锻炼，增强体质。")
		} else if fatRate > 0.1 && fatRate <= 0.16 {
			fmt.Println("目前是：标准。太棒了，要保持哦")
		} else if fatRate > 0.16 && fatRate <= 0.21 {
			fmt.Println("目前是：偏胖。吃完饭多散散步，消化消化")
		} else if fatRate > 0.21 && fatRate <= 0.26 {
			fmt.Println("目前是：肥胖。少吃点，多运动运动")
		} else {
			fmt.Println("目前是：非常肥胖。健身游泳跑步，即刻开始")
		}
	} else if age >= 40 && age <= 59 {
		// todo
	} else if age >= 60 {
		// todo
	} else {
		fmt.Println("我们不看未成年人的体脂率，变化太大，无法评判")
	}
}

//计算体脂率
func caclFatRate(weight float64, tall float64, age int, sexWeight int) float64 {
	//计算体脂率
	bmi := weight / (tall * tall)
	fatRate := (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("体脂率是：", fatRate)
	return fatRate
}

//读取输入内容
func getMaterialsFromInput() (inWeight float64, inTall float64, inAge int, inSexWeight int, inSex string) {
	var name string
	fmt.Print("姓名：")
	fmt.Scanln(&name)

	var weight float64
	fmt.Print("体重（千克）：")
	fmt.Scanln(&weight)

	var tall float64
	fmt.Print("身高（米）：")
	fmt.Scanln(&tall)

	var age int
	fmt.Print("年龄：")
	fmt.Scanln(&age)

	var sexWeight int
	var sex string = "男"
	fmt.Print("性别（男/女）：")
	fmt.Scanln(&sex)

	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	return weight, tall, age, sexWeight, sex
}

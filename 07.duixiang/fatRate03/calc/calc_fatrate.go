package calc

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64, err error) {
	if bmi <= 0 {
		panic("BMI不可为0或负数！")
		//fmt.Println("BMI不可为0或负数！")
		//return 1, err
	}

	if age <= 0 {
		panic("年龄不可为0或负数！")
	} else if age > 150 {
		panic("年龄不可超过150岁！")
	}
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else if sex == "女" {
		sexWeight = 0
	} else {
		panic("输入的性别为非法数据！")
	}
	fatRate = (1.2*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return fatRate, nil
}

//根据不同年龄体脂必中不同，不定期调整
func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23

	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return
}

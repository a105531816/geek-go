package main

import "fmt"

//计算多个人的平均体脂
//实现完整的体脂计算器
//连续输入三人的姓名、性别、身高、体重、年龄信息
//计算每个人的BMI、体脂率
//输出：
//   每个人的姓名、BMI、体脂率、建议
//   总人数、平均体脂率
func main() {
	var name, sex [3]string
	var height, weight, age [3]float64
	var SexWeight [3]int
	//输入三个人的信息
	for i := 0; i < len(name); i++ {

		fmt.Printf("请第%v位同学输入个人信息！！\n", i+1)
		fmt.Printf("请输入姓名：")
		fmt.Scanln(&name[i])
		fmt.Printf("请输入性别：")
		fmt.Scanln(&sex[i])
		if sex[i] == "男" {
			SexWeight[i] = 1
		}
		fmt.Printf("请输入身高(米)：")
		fmt.Scanln(&height[i])
		fmt.Printf("请输入体重(公斤)：")
		fmt.Scanln(&weight[i])
		fmt.Printf("请输入年龄：")
		fmt.Scanln(&age[i])
	}
	//输出体脂率及建议
	for i := 0; i < len(name); i++ {
		bmi := weight[i] / (height[i] * height[i])
		fatRate := 1.2*bmi + 0.23*float64(age[i]) - 5.4 - 10.8*float64(SexWeight[i])
		switch SexWeight[i] {
		case 0:
			if age[i] >= 18 && age[i] <= 39 {
				if fatRate <= 0.20 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "多吃点")
				} else if fatRate <= 0.27 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "稍微吃点")
				} else if fatRate <= 0.34 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "可以吃点")
				} else if fatRate <= 0.39 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "别吃了")
				} else {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "减肥")
				}
			} else if age[i] > 39 && age[i] <= 59 {
				if fatRate <= 0.21 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "多吃点")
				} else if fatRate <= 0.28 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "稍微吃点")
				} else if fatRate <= 0.35 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "可以吃点")
				} else if fatRate <= 0.40 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "别吃了")
				} else {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "减肥")
				}
			} else if age[i] >= 60 {
				if fatRate <= 0.22 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "多吃点")
				} else if fatRate <= 0.29 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "稍微吃点")
				} else if fatRate <= 0.36 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "可以吃点")
				} else if fatRate <= 0.41 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "别吃了")
				} else {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "减肥")
				}
			} else {
				fmt.Println("未成年人不进行体脂率计算！")
			}
		case 1:
			if age[i] >= 18 && age[i] <= 39 {
				if fatRate <= 0.10 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "多吃点")
				} else if fatRate <= 0.16 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "稍微吃点")
				} else if fatRate <= 0.21 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "可以吃点")
				} else if fatRate <= 0.26 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "别吃了")
				} else {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "减肥")
				}
			} else if age[i] > 39 && age[i] <= 59 {
				if fatRate <= 0.11 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "多吃点")
				} else if fatRate <= 0.17 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "稍微吃点")
				} else if fatRate <= 0.22 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "可以吃点")
				} else if fatRate <= 0.27 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "别吃了")
				} else {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "减肥")
				}
			} else if age[i] >= 60 {
				if fatRate <= 0.13 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "多吃点")
				} else if fatRate <= 0.19 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "稍微吃点")
				} else if fatRate <= 0.24 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "可以吃点")
				} else if fatRate <= 0.29 {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "别吃了")
				} else {
					fmt.Printf("姓名：%v\tBMI：%.2f\t体脂率：%.2f\t建议：%v\n", name[i], bmi, fatRate, "减肥")
				}
			} else {
				fmt.Println("未成年人不进行体脂率计算！")
			}
		}
	}
}

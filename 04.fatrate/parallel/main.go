package main

import "fmt"

//判断两条线是否平行
//提示：
//• 两点决定一条直线
//• 两条线是否平行取决于两条线的斜率是否一样
//K＝(y2-y1)/(x2-x1)

func main() {
	var x1, x2, y1, y2, i1, i2, j1, j2 int
	fmt.Println("请输入直线在坐标系中的第一个点（x,y）：")
	fmt.Scanln(&x1)
	fmt.Scanln(&y1)
	fmt.Println("请输入直线在坐标系中的第二个点（x,y）：")
	fmt.Scanln(&x2)
	fmt.Scanln(&y2)
	fmt.Println("请输入直线在坐标系中的第一个点（x,y）：")
	fmt.Scanln(&i1)
	fmt.Scanln(&j1)
	fmt.Println("请输入直线在坐标系中的第二个点（x,y）：")
	fmt.Scanln(&i2)
	fmt.Scanln(&j2)
	//输出两条直线
	fmt.Printf("第一条直线的点(%d,%d),(%d,%d)",x1,y1,x2,y2)
	fmt.Printf("第一条直线的点(%d,%d),(%d,%d)",i1,j1,i2,j2)
	//求b
	var b1 int
	var b2 int
	var k1 int
	var k2 int
	if (x2-x1) == 0 && (i2-i1) == 0 {
		fmt.Println("两条直线均垂直于x轴，平行")
	} else if (x2-x1) == 1 && (i2-i1) == 1 {
		fmt.Println("两条直线均垂直于y轴，平行")
	} else if (x2-x1) == 0 || (x2-x1) == 1 {
		fmt.Println("两条直线不平行-1")
	} else if (i2-i1) == 0 || (i2-i1) == 1 {
		fmt.Println("两条直线不平行-2")
	} else {
		k1 = (y2 - y1) / (x2 - x1)
		b1 = x2 - (x1 * k1)
		k2 = (j2 - j1) / (i2 - i1)
		b2 = x2 - (x1 * k1)
		if k1 == k2 && b1 == b2 {
			fmt.Println("两条直线平行，并且重合")
		} else if k1 == k2 {
			fmt.Println("两条直线平行")
		} else {
			fmt.Println("两条直线不平行-3")
		}
	}
}

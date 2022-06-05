package main

import "time"

//创建电梯对象
type dianti struct {
}

//电梯上行
func (d *dianti) diantiUp() {
	//每秒中走一层
	time.Sleep(1 * time.Second)
}

//电梯下行
func (d *dianti) diantiDown() {
	//每秒中走一层
	time.Sleep(1 * time.Second)
}

//电梯停留
func (d *dianti) diantistay() {

}

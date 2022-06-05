package main

type Person struct {
	//Id     int     `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true;column:id"`
	Id     int     `json:"id,omitempty" gorm:"column:id"`
	Name   string  `json:"name,omitempty"   gorm:"column:name"`
	Sex    string  `json:"sex,omitempty"   gorm:"column:sex"`
	Tall   float64 `json:"tall,omitempty"   gorm:"column:tall"`
	Weight float64 `json:"weight,omitempty"   gorm:"column:weight"`
	Age    int     `json:"age,omitempty"   gorm:"column:age"`
}

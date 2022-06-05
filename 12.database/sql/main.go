package main

import (
	_ "database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
)

func main() {
	//connection
	conn := ConnectDatabase()

	//insert
	//if err := newPerson(conn); err != nil {
	//	fmt.Println("插入失败")
	//}

	//select
	if err := ormSelect(conn); err != nil {
		fmt.Println("查询失败")
	}

	//select 范围查询
	//if err := ormSelectWith(conn); err != nil {
	//	fmt.Println("查询成功")
	//}

	//update
	//if err := ormUpdate(conn); err != nil {
	//	fmt.Println("更新成功")
	//}

	//指定字段update
	if err := ormUpdateFields(conn); err != nil {
		fmt.Println("更新失败")
	}
	//
	////delete
	if err := ormDelete(conn); err != nil {
		fmt.Println("删除失败")
	}
}

//Connect
func ConnectDatabase() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:root@tcp(10.0.0.44:3308)/test"))
	if err != nil {
		log.Fatal("连接数据库失败：", err)
	}
	log.Printf("连接到数据库\n")
	return conn
}

//Insert
func newPerson(conn *gorm.DB) error {
	res := conn.Create(&Person{
		Name:   "9马",
		Sex:    "男",
		Tall:   1.65,
		Weight: 64,
		Age:    20,
	})
	if err := res.Error; err != nil {
		fmt.Println("shibai：", err)
	}
	return nil
}

//实现Person的方法，才能精确定位表名
func (p *Person) TableName() string {
	return "personal_info"
}

//select
func ormSelect(conn *gorm.DB) error {
	result := make([]*Person, 10)
	resp := conn.Where(&Person{Name: "2牛"}).Find(&result)
	if err := resp.Error; err != nil {
		fmt.Println("查询错误", resp.Error)
		return err
	}
	fmt.Println("查询成功")
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	//for _, j := range result {
	//	fmt.Println(j.Name, j.Age)
	//}
	return nil
}

//select 范围查询
func ormSelectWith(conn *gorm.DB) error {
	result := make([]*Person, 10)
	resp := conn.Where("tall = ?", "1.7").Where("weight >=60").Find(&result)
	if err := resp.Error; err != nil {
		fmt.Println("查询错误", resp.Error)
		return err
	}
	fmt.Println("查询成功")
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	return nil
}

//Update
func ormUpdate(conn *gorm.DB) error {
	resp := conn.Updates(&Person{
		Id:     2,
		Name:   "213",
		Sex:    "34",
		Tall:   1.6,
		Weight: 71,
		Age:    30,
	})
	if err := resp.Error; err != nil {
		fmt.Println("更新错误", resp.Error)
		return err
	}
	fmt.Println("更新成功！")
	return nil
}

//Update
func ormUpdateFields(conn *gorm.DB) error {
	up := &Person{
		Id:     2,
		Name:   "12猪",
		Sex:    "34",
		Tall:   1.5,
		Weight: 71,
		Age:    30,
	}
	//只更新name,tall字段
	resp := conn.Model(up).Select("name", "tall").Updates(up)
	if err := resp.Error; err != nil {
		fmt.Println("更新错误", resp.Error)
		return err
	}
	fmt.Println("更新成功！")
	return nil
}

//Delete
func ormDelete(conn *gorm.DB) error {
	de := &Person{
		Id: 3,
	}
	resp := conn.Delete(de)
	if err := resp.Error; err != nil {
		fmt.Println("删除错误", resp.Error)
		return err
	}
	return nil
}

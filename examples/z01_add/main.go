package main

/*
@Time : 2022/6/20 23:54
@Author : 张大鹏
@File : main
@Software: Goland2021.3.1
@Description: 创建记录
*/

import (
	"github.com/zhangdapeng520/zdpgo_orm"
	"github.com/zhangdapeng520/zdpgo_orm/gorm"
)

type User struct {
	gorm.Model
	Name string // 姓名
	Age  uint8  // 年龄
}

func main() {
	o, err := zdpgo_orm.NewWithConfig(&zdpgo_orm.Config{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "book",
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建数据库表
	o.CreateTables(&User{})

	// 创建用户
	var user = User{Name: "张大鹏", Age: 18}
	o.Add(&user) // 通过数据的指针来创建
}

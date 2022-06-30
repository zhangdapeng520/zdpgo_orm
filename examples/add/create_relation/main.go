package main

/*
@Time : 2022/6/20 23:54
@Author : 张大鹏
@File : main
@Software: Goland2021.3.1
@Description: 关联创建
*/

import (
	"github.com/zhangdapeng520/zdpgo_orm/driver/mysql"
	"github.com/zhangdapeng520/zdpgo_orm/gorm"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gorm.Open(dsn), &mysql.Config{})
	if err != nil {
		panic(err)
	}

	// 创建数据库表
	db.AutoMigrate(&User{}, &CreditCard{})
	db.Create(&User{
		Name:       "jinzhu",
		CreditCard: CreditCard{Number: "411111111111"},
	})
}

package main

/*
@Time : 2022/6/20 23:54
@Author : 张大鹏
@File : main
@Software: Goland2021.3.1
@Description: 创建记录
*/

import (
	"database/sql"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_orm"
	"github.com/zhangdapeng520/zdpgo_orm/driver/mysql"
	"time"
)

type User struct {
	zdpgo_orm.Model
	Name         string         // 姓名
	Email        string         // 邮箱
	Age          uint8          // 年龄
	Birthday     time.Time      // 生日
	MemberNumber sql.NullString // 身份证号
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := zdpgo_orm.Open(mysql.Open(dsn), &zdpgo_orm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建数据库表
	db.AutoMigrate(&User{})

	// 创建用户
	var user = User{Name: "张大鹏", Age: 18, Birthday: time.Now()}
	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println(result)
}

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
	"github.com/zhangdapeng520/zdpgo_orm/driver/mysql"
	"github.com/zhangdapeng520/zdpgo_orm/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string         // 姓名
	Email        string         // 邮箱
	Age          uint8          // 年龄
	Birthday     time.Time      // 生日
	MemberNumber sql.NullString // 身份证号
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建数据库表
	db.AutoMigrate(&User{})

	// 要有效地插入大量记录，请将一个 slice 传递给 Create 方法。 GORM 将生成单独一条SQL语句来插入所有数据，并回填主键的值，钩子方法也会被调用。
	var users = []User{{Name: "张珊珊"}, {Name: "李思思"}, {Name: "王舞儿"}}
	db.Create(&users)

	for _, user := range users {
		fmt.Println(user.ID)
	}
}

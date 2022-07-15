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
	"github.com/zhangdapeng520/zdpgo_orm/gorm"
	"github.com/zhangdapeng520/zdpgo_orm/mysql"
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
		panic(err)
	}

	// 创建数据库表
	db.AutoMigrate(&User{})

	// 创建记录并更新给出的字段。
	var user = User{Name: "王五", Age: 18, Birthday: time.Now()}
	db.Select("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("王五", 18, "2020-07-04 11:05:21.775")
}

package main

/*
@Time : 2022/6/20 23:54
@Author : 张大鹏
@File : main
@Software: Goland2021.3.1
@Description: 根据map创建
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
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	// 创建数据库表
	db.AutoMigrate(&User{})
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "张大鹏", "Age": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "张大鹏_1", "Age": 18},
		{"Name": "张大鹏_2", "Age": 20},
	})
}

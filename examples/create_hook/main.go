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

//GORM 允许用户定义的钩子有 BeforeSave, BeforeCreate, AfterSave, AfterCreate 创建记录时将调用这些钩子方法，请参考 Hooks 中关于生命周期的详细信息
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("在创建之前执行了。。。")
	return
}
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("在创建之后执行了。。。")
	return
}

func (u *User) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println("在创建之后执行了111。。。")
	return
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 创建数据库表
	db.AutoMigrate(&User{})

	// 使用 CreateInBatches 分批创建时，你可以指定每批的数量，例如：
	var users = []User{{Name: "张珊珊2"}, {Name: "李思思2"}, {Name: "王舞儿2"}}
	db.CreateInBatches(users, 100)

	for _, user := range users {
		fmt.Println(user.ID)
	}
}

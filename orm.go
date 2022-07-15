package zdpgo_orm

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_orm/gorm"
	"github.com/zhangdapeng520/zdpgo_orm/mysql"
)

/*
@Time : 2022/6/30 20:56
@Author : 张大鹏
@File : orm.go
@Software: Goland2021.3.1
@Description:
*/

type Orm struct {
	Config *Config
	Db     *gorm.DB
}

func New() (*Orm, error) {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) (*Orm, error) {
	o := &Orm{}

	// 配置
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		config.Port = 3306
	}
	if config.Username == "" {
		config.Username = "root"
	}
	if config.Password == "" {
		config.Password = "root"
	}
	if config.Database == "" {
		config.Database = "test"
	}
	o.Config = config

	// 连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	o.Db = db

	// 返回
	return o, nil
}

func (o *Orm) CreateTables(dest ...interface{}) error {
	return o.Db.AutoMigrate(dest...)
}

func (o *Orm) Add(data interface{}) {
	o.Db.Create(data)
}

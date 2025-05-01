package gormx

import (
	"course_learn/common/configx"
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB 对象
var (
	DB    *gorm.DB
	SQLDB *sql.DB
)

func Init() (db *gorm.DB, sqlDB *sql.DB) {
	username := configx.MysqlConfigData.UserName
	password := configx.MysqlConfigData.PassWord
	ip := configx.MysqlConfigData.HostName
	port := configx.MysqlConfigData.Port
	database := configx.MysqlConfigData.DataBase

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, ip, port, database)
	if configx.LogConfigData.Flag == 1 {
		fmt.Println("mysql连接DNS:", dsn)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // DSN data source name
		//DefaultStringSize:         256,   // string 类型字段的默认长度
		//DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	if err != nil {
		if configx.LogConfigData.Flag == 1 {
			panic(err)
		}
	}

	DB = db

	SQLDB, err = DB.DB()
	if err != nil {
		if configx.LogConfigData.Flag == 1 {
			panic(err)
		}
	}
	if configx.LogConfigData.Flag == 1 {
		fmt.Println("mysql连接成功")
	}
	return
}

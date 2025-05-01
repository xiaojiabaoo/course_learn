package configx

import (
	"course_learn/model"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	MysqlConfigData model.MysqlConfig
	RedisConfigData model.RedisConfig
	AppConfigData   model.AppConfig
	LogConfigData   model.LogConfig
)

func InitViper() {
	LogConfigData.Flag = 0 //默认关打印
	confPath, _ := os.Getwd()
	viper.AddConfigPath(confPath)
	viper.SetConfigName("app")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("初始化Viper出错：", err)
	}
}

func InitConfigData() {
	//读取日志配置数据
	LogConfigData.Flag = viper.GetInt64("log.flag")
	LogConfigData.Level = viper.GetString("log.level")

	//读取APP配置数据
	AppConfigData.Name = viper.GetString("app.name")
	AppConfigData.Port = viper.GetInt64("app.port")

	//读取Mysql配置数据
	MysqlConfigData.HostName = viper.GetString("mysql.hostname")
	MysqlConfigData.Port = viper.GetInt64("mysql.port")
	MysqlConfigData.DataBase = viper.GetString("mysql.database")
	MysqlConfigData.UserName = viper.GetString("mysql.username")
	MysqlConfigData.PassWord = viper.GetString("mysql.password")

	//读取Redis配置数据
	RedisConfigData.HostName = viper.GetString("redis.hostname")
	RedisConfigData.Port = viper.GetInt64("redis.port")
	RedisConfigData.DataBase = viper.GetInt64("redis.database")
	RedisConfigData.UserName = viper.GetString("redis.username")
	RedisConfigData.PassWord = viper.GetString("redis.password")
}

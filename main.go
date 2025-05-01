package main

import (
	"course_learn/common/configx"
	"course_learn/common/gormx"
	"course_learn/common/redisx"
	"course_learn/router"
	"course_learn/tools/util"
)

func main() {
	//获取配置文件，读取文件中的配置信息
	configx.InitViper()
	//初始化配置数据
	configx.InitConfigData()
	//初始化mysql
	gormx.Init()
	//初始化redis
	redisx.Init()
	//启动服务
	router.Run("0.0.0.0:" + util.Int2Str(int(configx.AppConfigData.Port)))
}

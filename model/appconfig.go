package model

type AppConfig struct {
	Name string `bson:"name" json:"name"`
	Port int64  `bson:"port" json:"port"`
}

type MysqlConfig struct {
	HostName string `bson:"hostname" json:"hostname"`
	UserName string `bson:"username" json:"username"`
	PassWord string `bson:"password" json:"password"`
	Port     int64  `bson:"port" json:"port"`
	DataBase string `bson:"database" json:"database"`
}

type RedisConfig struct {
	HostName string `bson:"hostname" json:"hostname"`
	UserName string `bson:"username" json:"username"`
	PassWord string `bson:"password" json:"password"`
	Port     int64  `bson:"port" json:"port"`
	DataBase int64  `bson:"database" json:"database"`
}

type LogConfig struct {
	Flag  int64  `bson:"flag" json:"flag"`
	Level string `bson:"lever" json:"lever"`
}

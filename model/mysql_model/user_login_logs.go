package mysql_model

import "time"

type UserLoginLogs struct {
	Id          int64     `gorm:"not null pk autoincr INT"`
	UserId      int64     `gorm:"default 0 comment('用户ID') INT"`
	LoginTime   int64     `gorm:"default 0 comment('用户登录时间') INT"`
	LoginIp     string    `gorm:"comment('用户登录IP') VARCHAR(255)"`
	LoginArea   string    `gorm:"comment('用户登录地区') VARCHAR(255)"`
	LoginMethod int64     `gorm:"comment('户登录方式：1.手机号码登录；2.邮箱地址登录；.......') VARCHAR(255)"`
	DateTime    time.Time `gorm:"DATETIME"`
}

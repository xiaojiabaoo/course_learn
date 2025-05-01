package mysql_model

import "time"

type UserParam struct {
	Id       int64       `gorm:"not null pk autoincr INT"`
	UserId   int64       `gorm:"comment('用户ID') INT"`
	TryCount int64       `gorm:"comment('用户的试用课程次数') INT"`
	TryTime  int64       `gorm:"comment('用户每次试用的时间，单位：天') INT"`
	AddTime  int64       `gorm:"comment('更新时间，与date_time时间相同') INT"`
	UpTime   int64       `gorm:"comment('更新时间，与date_time时间相同') INT"`
	DateTime time.Time `gorm:"DATETIME"`
}

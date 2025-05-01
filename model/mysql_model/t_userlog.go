package mysql_model

import "time"

type TUserlog struct {
	Id          int64       `gorm:"not null pk autoincr INT"`
	ProjectName string    `gorm:"comment('项目名称') VARCHAR(255)"`                       // 项目名称
	Platform    string    `gorm:"comment('平台') VARCHAR(255)"`                         // 平台
	UserId      int64       `gorm:"default 0 comment('用户ID') INT"`            // 用户ID
	UserToken   string    `gorm:"comment('用户token') VARCHAR(255)"`                    // 用户token
	UserInfo    string    `gorm:"comment('用户信息：身份|userId|Token|个推CID') VARCHAR(255)"` // 用户信息：身份|userId|Token|个推CID
	MsgContent  string    `gorm:"comment('日志内容') VARCHAR(255)"`                       // 日志内容
	DetailInfo  string    `gorm:"comment('详情信息') VARCHAR(255)"`                       // 详情信息
	CreateTime  time.Time `gorm:"DATETIME"`                                           // 创建时间
}

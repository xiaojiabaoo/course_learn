package mysql_model

import "time"

type Manager struct {
	Id         int64     `gorm:"not null pk autoincr INT"`                     // 管理员ID
	Account    string    `gorm:"comment('管理员账号') VARCHAR(255)"`                // 账号
	Email      string    `gorm:"comment('管理邮箱') VARCHAR(255)"`                 // 邮箱
	Password   string    `gorm:"comment('管理密码') VARCHAR(255)"`                 // 密码
	Identifier string    `gorm:"comment('识别码') VARCHAR(255)"`                  // 识别码
	Status     int64     `gorm:"comment('账号状态：1.启用；2.已删除') INT"`               // 识别码
	AddTime    int64     `gorm:"default 0 comment('添加时间/用户注册时间') INT"`         // 添加时间
	UpdateTime int64     `gorm:"default 0 comment('更新时间，与date_time时间相同') INT"` // 更新时间，与date_time时间相同
	DateTime   time.Time `gorm:"DATETIME"`                                     // 更新时间，与update_time时间相同
}

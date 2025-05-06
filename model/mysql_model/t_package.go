package mysql_model

import "time"

type TPackage struct {
	PackageId   int64     `gorm:"not null pk autoincr INT"`                     // 专业ID
	PackageName string    `gorm:"comment('专业名称') VARCHAR(255)"`                 // 专业名称
	Status      int64     `gorm:"comment('状态：1.上架；2.下架；3.已删除；') VARCHAR(255)"`  // 状态：1.上架；2.下架；3.已删除；
	AddTime     int64     `gorm:"default 0 comment('添加时间/用户注册时间') INT"`         // 添加时间
	UpdateTime  int64     `gorm:"default 0 comment('更新时间，与date_time时间相同') INT"` // 更新时间，与date_time时间相同
	DateTime    time.Time `gorm:"DATETIME"`                                     // 更新时间，与update_time时间相同
}

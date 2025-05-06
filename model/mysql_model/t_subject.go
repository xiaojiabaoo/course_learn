package mysql_model

import "time"

type TSubject struct {
	KnowledgeTreeId int64     `gorm:"not null pk autoincr INT"`      // 知识树ID
	SubjectId       int64     `gorm:"default 0 comment('课程ID') INT"` // 课程ID
	SubjectName     string    `gorm:"comment('课程名') VARCHAR(255)"`   // 课程名
	PackageId       int64     `gorm:"default 0 comment('专业ID') INT"` // 专业ID
	Sort            int64     `gorm:"default 0 comment('排序') INT"`   // 排序
	Desc            string    `gorm:"default 0 comment('课程简介说明') LONGTEXT"`
	IsHot           int64     `gorm:"default 0 comment('排序') INT"`                  // 是否热门：0.否；1.是
	Status          int64     `gorm:"comment('状态：1.上架；2.下架；3.已删除；') VARCHAR(255)"`  // 状态：1.上架；2.下架；3.已删除；
	AddTime         int64     `gorm:"default 0 comment('添加时间/用户注册时间') INT"`         // 添加时间
	UpdateTime      int64     `gorm:"default 0 comment('更新时间，与date_time时间相同') INT"` // 更新时间，与date_time时间相同
	DateTime        time.Time `gorm:"DATETIME"`                                     // 更新时间，与update_time时间相同
}

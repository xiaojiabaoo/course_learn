package mysql_model

import "time"

type TSection struct {
	Id            int64     `gorm:"not null pk autoincr INT"`                     // 节ID
	Sequence      string    `gorm:"comment('节顺序') VARCHAR(255)"`                  // 节顺序
	Name          string    `gorm:"comment('节名称') VARCHAR(255)"`                  // 节名称
	QuestionCount int64     `gorm:"default 0 comment('节问题总数') INT"`               // 节问题总数
	ChapterId     int64     `gorm:"default 0 comment('所属章节ID') INT"`              // 所属章节ID
	Status        int64     `gorm:"comment('状态：1.上架；2.下架；3.已删除；') VARCHAR(255)"`  // 状态：1.上架；2.下架；3.已删除；
	AddTime       int64     `gorm:"default 0 comment('添加时间/用户注册时间') INT"`         // 添加时间
	UpdateTime    int64     `gorm:"default 0 comment('更新时间，与date_time时间相同') INT"` // 更新时间，与date_time时间相同
	DateTime      time.Time `gorm:"DATETIME"`                                     // 更新时间，与update_time时间相同
}

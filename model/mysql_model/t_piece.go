package mysql_model

import "time"

type TPiece struct {
	Id            int64     `gorm:"not null pk autoincr INT"`                     // 块ID
	Name          string    `gorm:"comment('块名称') VARCHAR(255)"`                  // 块名称
	QuestionCount int64     `gorm:"default 0 comment('块问题总数') INT"`               // 块问题总数
	SectionId     int64     `gorm:"default 0 comment('所属节ID') INT"`               // 所属节ID
	Status        int64     `gorm:"comment('状态：1.上架；2.下架；3.已删除；') VARCHAR(255)"`  // 状态：1.上架；2.下架；3.已删除；
	AddTime       int64     `gorm:"default 0 comment('添加时间/用户注册时间') INT"`         // 添加时间
	UpdateTime    int64     `gorm:"default 0 comment('更新时间，与date_time时间相同') INT"` // 更新时间，与date_time时间相同
	DateTime      time.Time `gorm:"DATETIME"`                                     // 更新时间，与update_time时间相同
}

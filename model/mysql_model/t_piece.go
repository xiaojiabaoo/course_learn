package mysql_model

type TPiece struct {
	Id            int64  `gorm:"not null pk autoincr INT"`       // 块ID
	Name          string `gorm:"comment('块名称') VARCHAR(255)"`    // 块名称
	QuestionCount int64  `gorm:"default 0 comment('块问题总数') INT"` // 块问题总数
	SectionId     int64  `gorm:"default 0 comment('所属节ID') INT"` // 所属节ID
}

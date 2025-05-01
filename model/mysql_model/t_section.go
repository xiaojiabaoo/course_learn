package mysql_model

type TSection struct {
	Id            int64  `gorm:"not null pk autoincr INT"`        // 节ID
	Sequence      string `gorm:"comment('节顺序') VARCHAR(255)"`     // 节顺序
	Name          string `gorm:"comment('节名称') VARCHAR(255)"`     // 节名称
	QuestionCount int64  `gorm:"default 0 comment('节问题总数') INT"`  // 节问题总数
	ChapterId     int64  `gorm:"default 0 comment('所属章节ID') INT"` // 所属章节ID
}

package mysql_model

type TChapter struct {
	Id            int64  `gorm:"not null pk autoincr INT"`    // 章ID
	Sequence      string `gorm:"comment('章顺序') VARCHAR(255)"` // 章顺序
	Name          string `gorm:"comment('章名称') VARCHAR(255)"` // 章名称
	QuestionCount int64  `gorm:"comment('章问题总数') INT"`        // 章问题总数
	SubjectId     int64  `gorm:"comment('所属课程ID') INT"`       // 所属课程ID
	Sort          int64  `gorm:"comment('排序') INT"`
}

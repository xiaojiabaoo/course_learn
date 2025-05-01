package mysql_model

type TSubject struct {
	KnowledgeTreeId int64  `gorm:"not null pk autoincr INT"`      // 知识树ID
	SubjectId       int64  `gorm:"default 0 comment('课程ID') INT"` // 课程ID
	SubjectName     string `gorm:"comment('课程名') VARCHAR(255)"`   // 课程名
	PackageId       int64  `gorm:"default 0 comment('专业ID') INT"` // 专业ID
	Sort            int64  `gorm:"default 0 comment('排序') INT"`   // 排序
	Desc            string `gorm:"default 0 comment('课程简介说明') LONGTEXT"`
	IsHot           int64  `gorm:"default 0 comment('排序') INT"` // 是否热门：0.否；1.是
}

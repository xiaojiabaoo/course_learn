package mysql_model

import "time"

type UserLearnRecord struct {
	Id          int64     `gorm:"not null pk autoincr INT"`
	UserId      int64     `gorm:"default 0 comment('用户ID') INT"` // 用户ID
	PackageId   int64     `gorm:"default 0 comment('学习专业ID') INT"`
	PackageName string    `gorm:"comment('学习专业名称') VARCHAR(255)"`
	SubjectId   int64     `gorm:"default 0 comment('学习课程ID') INT"`
	SubjectName string    `gorm:"comment('学习课程名称') VARCHAR(255)"`
	ChapterId   int64     `gorm:"default 0 comment('学习章ID') INT"`
	ChapterName string    `gorm:"comment('学习章名称') VARCHAR(255)"`
	SectionId   int64     `gorm:"default 0 comment('学习节ID') INT"`
	SectionName string    `gorm:"comment('学习节名称') VARCHAR(255)"`
	PieceId     int64     `gorm:"default 0 comment('学习节ID') INT"`
	PieceName   string    `gorm:"comment('学习节名称') VARCHAR(255)"`
	AddTime     int64     `gorm:"default 0 comment('学习时间') INT"`
	DateTime    time.Time `gorm:"DATETIME"` // 创建时间
}

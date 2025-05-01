package mysql_model

type TSubjectChapter struct {
	TSubject  `gorm:"extends"`
	UserOrder `gorm:"extends"`
}

type TChapterSection struct {
	TChapter `gorm:"extends"`
	TSection `gorm:"extends"`
	TPiece   `gorm:"extends"`
}

type TQuestionOption struct {
	TQuestion `gorm:"extends"`
	TOption   `gorm:"extends"`
}

type TChildQuestionOption struct {
	TChildQuestion `gorm:"extends"`
	TOption        `gorm:"extends"`
}

type SubjectByPieceId struct {
	TPiece   `gorm:"extends"`
	TSection `gorm:"extends"`
	TChapter `gorm:"extends"`
	TSubject `gorm:"extends"`
}

type SubjectByPiece struct {
	PackageId   int64  `json:"package_id"`   // 专业ID
	PackageName string `json:"package_name"` // 专业名称
	SubjectId   int64  `json:"subject_id"`   // 课程ID
	SubjectName string `json:"subject_name"` // 课程名
	ChapterId   int64  `json:"chapter_id"`   // 章ID
	ChapterName string `json:"chapter_name"` // 章名称
	SectionId   int64  `json:"section_id"`   // 节ID
	SectionName string `json:"section_name"` // 节名称
	PieceId     int64  `json:"piece_id"`     // 块ID
	PieceName   string `json:"piece_name"`   // 块名称
}

type UserOrderList struct {
	UserOrder   `gorm:"extends"`
	SubjectName string `json:"subject_name"`
}

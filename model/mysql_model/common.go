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
	/*TQuestion `gorm:"extends"`
	TOption   `gorm:"extends"`*/
	Id                int64  `json:"id"`
	PieceId           int64  `json:"piece_id"`
	Sequence          int64  `json:"sequence"`
	QuestionType      string `json:"question_type"`
	QuestionSource    string `json:"question_source"`
	QuestionContent   string `json:"question_content"`
	QuestionAnswer    string `json:"question_answer"`
	Score             string `json:"score"`
	MainNodeId        int64  `json:"main_node_id"`
	MainNodeName      string `json:"main_node_name"`
	MainNodeFrequency string `json:"main_node_frequency"`
	Favorite          int64  `json:"favorite"`
	CanPhoto          int64  `json:"can_photo"`
	AvgCorrectRate    string `json:"avg_correct_rate"`
	AnalysisType      int64  `json:"analysis_type"`
	Analysis          string `json:"analysis"`
	OptionId          int64  `json:"option_id"`
	OptionSequence    int64  `json:"option_sequence"`
	OptionTitle       string `json:"option_title"`
	OptionContent     string `json:"option_content"`
	OptionCorrect     int64  `json:"option_correct"`
}

type TChildQuestionOption struct {
	/*TChildQuestion `gorm:"extends"`
	TOption        `gorm:"extends"`*/
	Id                int64  `json:"id"`
	PieceId           int64  `json:"piece_id"`
	Sequence          int64  `json:"sequence"`
	QuestionType      string `json:"question_type"`
	QuestionSource    string `json:"question_source"`
	QuestionContent   string `json:"question_content"`
	QuestionAnswer    string `json:"question_answer"`
	Score             string `json:"score"`
	MainNodeId        int64  `json:"main_node_id"`
	MainNodeName      string `json:"main_node_name"`
	MainNodeFrequency string `json:"main_node_frequency"`
	Favorite          int64  `json:"favorite"`
	CanPhoto          int64  `json:"can_photo"`
	AvgCorrectRate    string `json:"avg_correct_rate"`
	AnalysisType      int64  `json:"analysis_type"`
	Analysis          string `json:"analysis"`
	OptionId          int64  `json:"option_id"`
	OptionSequence    int64  `json:"option_sequence"`
	OptionTitle       string `json:"option_title"`
	OptionContent     string `json:"option_content"`
	OptionCorrect     int64  `json:"option_correct"`
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
	NickName    string `json:"nick_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

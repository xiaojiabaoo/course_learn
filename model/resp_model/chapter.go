package resp_model

type GetChapterList struct {
	Id            int64     `json:"id"`
	Sequence      string    `json:"sequence"`
	Name          string    `json:"name"`
	QuestionCount int64     `json:"question_count"`
	SubjectId     int64     `json:"subject_id"`
	Sort          int64     `json:"sort"`
	Section       []Section `json:"section"`
}

type Section struct {
	Id            int64   `json:"id"`
	Sequence      string  `json:"sequence"`
	Name          string  `json:"name"`
	QuestionCount int64   `json:"question_count"`
	ChapterId     int64   `json:"chapter_id"`
	Piece         []Piece `json:"piece"`
}

type Piece struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	QuestionCount int64  `json:"question_count"`
	SectionId     int64  `json:"section_id"`
}

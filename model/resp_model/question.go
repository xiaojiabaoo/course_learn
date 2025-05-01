package resp_model

type GetQuestionData struct {
	Id                int64               `json:"id"`
	PieceId           int64               `json:"piece_id"`
	Sequence          int64               `json:"sequence"`
	QuestionType      string              `json:"question_type"`
	QuestionSource    string              `json:"question_source"`
	QuestionContent   string              `json:"question_content"`
	QuestionAnswer    string              `json:"question_answer"`
	Score             string              `json:"score"`
	MainNodeId        int64               `json:"main_node_id"`
	MainNodeName      string              `json:"main_node_name"`
	MainNodeFrequency string              `json:"main_node_frequency"`
	Favorite          int64               `json:"favorite"`
	CanPhoto          int64               `json:"can_photo"`
	AvgCorrectRate    string              `json:"avg_correct_rate"`
	AnalysisType      int64               `json:"analysis_type"`
	Analysis          string              `json:"analysis"`
	Option            []Option            `json:"option"`
	ChildQuestionData []ChildQuestionData `json:"child_question_data"`
}

type ChildQuestionData struct {
	Id                int64    `json:"id"`
	PieceId           int64    `json:"piece_id"`
	Sequence          int64    `json:"sequence"`
	QuestionType      string   `json:"question_type"`
	QuestionSource    string   `json:"question_source"`
	QuestionContent   string   `json:"question_content"`
	QuestionAnswer    string   `json:"question_answer"`
	Score             string   `json:"score"`
	MainNodeId        int64    `json:"main_node_id"`
	MainNodeName      string   `json:"main_node_name"`
	MainNodeFrequency string   `json:"main_node_frequency"`
	Favorite          int64    `json:"favorite"`
	CanPhoto          int64    `json:"can_photo"`
	AvgCorrectRate    string   `json:"avg_correct_rate"`
	AnalysisType      int64    `json:"analysis_type"`
	Analysis          string   `json:"analysis"`
	Option            []Option `json:"option"`
}

type Option struct {
	Id         int64  `json:"id"`
	QuestionId int64  `json:"question_id"`
	Sequence   int64  `json:"sequence"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Correct    int64  `json:"correct"`
}

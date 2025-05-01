package req_model

type GetQuestionData struct {
	PieceId   int64 `json:"piece_id"` //块ID
	UserId    int64 `json:"user_id"`
	SubjectId int64 `json:"subject_id"`
}

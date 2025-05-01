package mysql_model

type TChildQuestion struct {
	Pid               int64  `gorm:"not null pk autoincr INT"`     // 父问题ID
	Id                int64  `gorm:"not null pk autoincr INT"`     // 问题ID
	PieceId           int64  `gorm:"default 0 comment('块ID') INT"` // 块ID
	Sequence          int64  `gorm:"default 0 comment('顺序') INT"`  // 顺序
	QuestionType      string `gorm:"comment('问题类型') VARCHAR(255)"` // 问题类型
	QuestionSource    string `gorm:"comment('问题源') VARCHAR(255)"`  // 问题源
	QuestionContent   string `gorm:"comment('问题内容') VARCHAR(255)"` // 问题内容
	QuestionAnswer    string `gorm:"comment('问题答案') VARCHAR(255)"` // 问题答案
	Score             string `gorm:"comment('分数') VARCHAR(255)"`
	MainNodeId        int64  `gorm:"default 0 comment('主节点ID') INT"`
	MainNodeName      string `gorm:"comment('主节点名称') VARCHAR(255)"`
	MainNodeFrequency string `gorm:"comment('主节点') VARCHAR(255)"`
	Favorite          int64  `gorm:"default 0 comment('类型') INT"`
	CanPhoto          int64  `gorm:"default 0 comment('类型') INT"`
	AvgCorrectRate    string `gorm:"comment('类型') VARCHAR(255)"`
	AnalysisType      int64  `gorm:"default 0 comment('解析类型') INT"` // 解析类型
	Analysis          string `gorm:"comment('解析描述') VARCHAR(255)"`  // 解析描述
}

package mysql_model

import "time"

type TQuestion struct {
	Id                int64     `gorm:"not null pk autoincr INT"`     // 问题ID
	PieceId           int64     `gorm:"default 0 comment('块ID') INT"` // 块ID
	Sequence          int64     `gorm:"default 0 comment('顺序') INT"`  // 顺序
	QuestionType      string    `gorm:"comment('问题类型') VARCHAR(255)"` // 问题类型
	QuestionSource    string    `gorm:"comment('问题源') VARCHAR(255)"`  // 问题源
	QuestionContent   string    `gorm:"comment('问题内容') VARCHAR(255)"` // 问题内容
	QuestionAnswer    string    `gorm:"comment('问题答案') VARCHAR(255)"` // 问题答案
	Score             string    `gorm:"comment('分数') VARCHAR(255)"`
	MainNodeId        int64     `gorm:"default 0 comment('主节点ID') INT"`
	MainNodeName      string    `gorm:"comment('主节点名称') VARCHAR(255)"`
	MainNodeFrequency string    `gorm:"comment('主节点') VARCHAR(255)"`
	Favorite          int64     `gorm:"default 0 comment('类型') INT"`
	CanPhoto          int64     `gorm:"default 0 comment('类型') INT"`
	AvgCorrectRate    string    `gorm:"comment('类型') VARCHAR(255)"`
	AnalysisType      int64     `gorm:"default 0 comment('解析类型') INT"`                // 解析类型
	Analysis          string    `gorm:"comment('解析描述') VARCHAR(255)"`                 // 解析描述
	Status            int64     `gorm:"comment('状态：1.上架；2.下架；3.已删除；') VARCHAR(255)"`  // 状态：1.上架；2.下架；3.已删除；
	AddTime           int64     `gorm:"default 0 comment('添加时间/用户注册时间') INT"`         // 添加时间
	UpdateTime        int64     `gorm:"default 0 comment('更新时间，与date_time时间相同') INT"` // 更新时间，与date_time时间相同
	DateTime          time.Time `gorm:"DATETIME"`                                     // 更新时间，与update_time时间相同
}

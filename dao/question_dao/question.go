package question_dao

import (
	"course_learn/common/gormx"
	"course_learn/model/mysql_model"
	"github.com/pkg/errors"
)

type Question interface {
	GetQuestionData(pieceId int64) ([]mysql_model.TQuestionOption, error)
	GetChildQuestionData(questionId, pieceId int64) ([]mysql_model.TChildQuestionOption, error)
}

type ApiQuestion struct{}

func (a ApiQuestion) GetQuestionData(pieceId int64) ([]mysql_model.TQuestionOption, error) {
	data := make([]mysql_model.TQuestionOption, 0, 0)
	err := gormx.DB.Table("t_question q").Select("q.id, q.piece_id, q.sequence, "+
		"q.question_type, q.question_source, q.question_content, q.question_answer, "+
		"q.score, q.main_node_id, q.main_node_name, q.main_node_frequency, "+
		"q.favorite, q.can_photo, q.avg_correct_rate, q.analysis_type, "+
		"q.analysis, o.id option_id, o.sequence option_sequence, o.title option_title, o.content option_content, "+
		"o.correct option_correct").
		Joins("LEFT JOIN t_option o ON q.id = o.question_id").
		Where("q.piece_id=?", pieceId).
		Where("q.status=1").
		Order("q.sequence ASC").
		Order("o.sequence ASC").
		Scan(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "根据块ID获取题目列表失败")
	}
	return data, nil
}

func (a ApiQuestion) GetChildQuestionData(questionId, pieceId int64) ([]mysql_model.TChildQuestionOption, error) {
	data := make([]mysql_model.TChildQuestionOption, 0, 0)
	err := gormx.DB.Debug().Table("t_child_question cq").Select("cq.id, cq.piece_id, "+
		"cq.sequence, cq.question_type, cq.question_source, "+
		"cq.question_content, cq.question_answer, cq.score, "+
		"cq.main_node_id, cq.main_node_name, cq.main_node_frequency, "+
		"cq.favorite, cq.can_photo, cq.avg_correct_rate, "+
		"cq.analysis_type, cq.analysis, o.id option_id, "+
		"o.sequence option_sequence, o.title option_title, o.content option_content, o.correct option_correct").
		Joins("LEFT JOIN t_option o ON cq.id = o.question_id").
		Where("cq.pid=? AND cq.piece_id=?", questionId, pieceId).
		Order("cq.sequence ASC").Order("o.sequence ASC").Scan(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "根据块ID获取题目列表失败")
	}
	return data, nil
}

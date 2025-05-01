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
	err := gormx.DB.Table("t_question").Select("t_question.id, t_question.piece_id, t_question.sequence, " +
		"t_question.question_type, t_question.question_source, t_question.question_content, t_question.question_answer, " +
		"t_question.score, t_question.main_node_id, t_question.main_node_name, t_question.main_node_frequency, " +
		"t_question.favorite, t_question.can_photo, t_question.avg_correct_rate, t_question.analysis_type, " +
		"t_question.analysis, t_option.id, t_option.question_id, t_option.sequence, t_option.title, t_option.content, " +
		"t_option.correct").
		Joins("LEFT JOIN t_option ON t_question.id = t_option.question_id").
		Where("t_question.piece_id=?", pieceId).
		Order("t_question.sequence ASC").
		Order("t_option.sequence ASC").
		Scan(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "根据块ID获取题目列表失败")
	}
	return data, nil
}

func (a ApiQuestion) GetChildQuestionData(questionId, pieceId int64) ([]mysql_model.TChildQuestionOption, error) {
	data := make([]mysql_model.TChildQuestionOption, 0, 0)
	err := gormx.DB.Table("t_child_question").Select("t_child_question.id, t_child_question.piece_id, " +
		"t_child_question.sequence, t_child_question.question_type, t_child_question.question_source, " +
		"t_child_question.question_content, t_child_question.question_answer, t_child_question.score, " +
		"t_child_question.main_node_id, t_child_question.main_node_name, t_child_question.main_node_frequency, " +
		"t_child_question.favorite, t_child_question.can_photo, t_child_question.avg_correct_rate, " +
		"t_child_question.analysis_type, t_child_question.analysis, t_option.id, t_option.question_id, " +
		"t_option.sequence, t_option.title, t_option.content, t_option.correct").
		Joins("LEFT JOIN t_option ON t_child_question.id = t_option.question_id").
		Where("t_child_question.pid=? AND t_child_question.piece_id=?", questionId, pieceId).
		Order("t_child_question.sequence ASC").Order("t_option.sequence ASC").Scan(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "根据块ID获取题目列表失败")
	}
	return data, nil
}

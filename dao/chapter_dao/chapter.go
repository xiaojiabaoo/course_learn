package chapter_dao

import (
	"course_learn/common/gormx"
	"course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"github.com/pkg/errors"
)

type Chapter interface {
	GetChapterList(req_model.GetChapterList) ([]mysql_model.TChapterSection, error)
}

type ApiChapter struct{}

func (a ApiChapter) GetChapterList(param req_model.GetChapterList) ([]mysql_model.TChapterSection, error) {
	section := make([]mysql_model.TChapterSection, 0, 0)
	err := gormx.DB.Table("t_chapter").
		Joins("LEFT JOIN t_section ON t_chapter.id = t_section.chapter_id").
		Joins("LEFT JOIN t_piece ON t_piece.section_id = t_section.id").
		Where("t_chapter.subject_id=?", param.SubjectId).
		Order("t_chapter.sort").Order("t_section.id").Order("t_piece.id").
		Scan(&section).Error
	if err != nil {
		return section, errors.Wrap(err,"根据科目ID查询章节信息失败")
	}
	return section, nil
}
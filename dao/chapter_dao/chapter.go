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
	err := gormx.DB.Table("t_chapter c").
		Joins("LEFT JOIN t_section s ON c.id = s.chapter_id").
		Joins("LEFT JOIN t_piece p ON p.section_id = s.id").
		Where("c.subject_id=?", param.SubjectId).
		Where("c.status=1 and s.status=1 and p.status=1").
		Order("c.sort").Order("s.id").Order("p.id").
		Scan(&section).Error
	if err != nil {
		return section, errors.Wrap(err, "根据科目ID查询章节信息失败")
	}
	return section, nil
}

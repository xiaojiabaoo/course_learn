package subject_dao

import (
	"course_learn/common/gormx"
	"course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"github.com/pkg/errors"
)

type Subject interface {
	GetSubjectData(userId, types, time int64) ([]mysql_model.TSubjectChapter, error)
	GetUserTrySubList(req_model.TrySubject) ([]mysql_model.UserOrder, error)
	AddTrySubject(mysql_model.UserOrder) error
	GetSubjectDetail(subjectId int64) (*mysql_model.TSubject, error)
	GetSubjectByPieceId(pieceId int64) ([]map[string]string, error)
	GetSubjectPiece(subjectId []int64) ([]mysql_model.SubjectByPieceId, error)
}

type ApiSubject struct{}

func (a *ApiSubject) GetSubjectData(userId, types, time int64) ([]mysql_model.TSubjectChapter, error) {
	subject := make([]mysql_model.TSubjectChapter, 0, 0)
	joins := gormx.DB.Table("t_subject").Select("t_subject.knowledge_tree_id, t_subject.subject_id, t_subject.subject_name, " +
		"t_subject.package_id, user_order.type, user_order.user_id, user_order.expiration_time").
		Joins("LEFT JOIN user_order ON user_order.subject_id = t_subject.subject_id")
	if types == 1 || types == 3 {
		joins = joins.Where("user_order.user_id=? and is_valid=? and user_order.expiration_time > ?", userId, 1, time)
	}
	if types == 4 {
		joins = joins.Where("t_subject.is_hot=?", 1)
	}
	err := joins.Order("t_subject.sort asc").Order("user_order.expiration_time desc").Scan(&subject).Error
	if err != nil {
		return subject, errors.Wrap(err, "查询课程表失败")
	}
	return subject, nil
}

func (a *ApiSubject) GetUserTrySubList(param req_model.TrySubject) ([]mysql_model.UserOrder, error) {
	order := make([]mysql_model.UserOrder, 0, 0)
	err := gormx.DB.Where("user_id=?", param.UserId).Where("type=?", 1).Find(&order).Error
	if err != nil {
		return order, errors.Wrap(err, "查询用户订单失败")
	}
	return order, nil
}

func (a *ApiSubject) AddTrySubject(order mysql_model.UserOrder) error {
	err := gormx.DB.Create(&order).Error
	if err != nil {
		return errors.Wrap(err, "添加课程试用失败")
	}
	return nil
}

func (a *ApiSubject) GetSubjectDetail(subjectId int64) (*mysql_model.TSubject, error) {
	subject := &mysql_model.TSubject{}
	err := gormx.DB.Where("subject_id=?", subjectId).First(subject).Error
	if err != nil {
		return subject, errors.Wrap(err, "查询单个课程信息失败")
	}
	return subject, nil
}

func (a *ApiSubject) GetSubjectByPieceId(pieceId int64) ([]mysql_model.SubjectByPiece, error) {
	data := make([]mysql_model.SubjectByPiece, 0, 0)
	/*err := gormx.DB.Raw("SELECT " +
	"tpa.package_id, " +
	"tpa.package_name, " +
	"tsu.subject_id, " +
	"tsu.subject_name," +
	"tc.id chapter_id, tc.`name` chapter_name, " +
	"tse.id section_id, tse.`name` section_name, " +
	"tpi.id piece_id, tpi.`name` piece_name " +
	"FROM t_piece tpi " +
	"INNER JOIN t_section tse ON tse.id = tpi.section_id " +
	"INNER JOIN t_chapter tc ON tc.id = tse.chapter_id " +
	"INNER JOIN t_subject tsu ON tsu.subject_id = tc.subject_id " +
	"INNER JOIN t_package tpa ON tpa.package_id = tsu.package_id WHERE tpi.id = " + util.Int2Str(int(pieceId))).Scan(&maps).Error
	*/
	err := gormx.DB.Debug().Table("t_piece").
		Select("t_package.package_id,t_package.package_name,t_subject.subject_id,t_subject.subject_name,"+
			"t_chapter.id chapter_id, t_chapter.`name` chapter_name,t_section.id section_id, t_section.`name` section_name,"+
			"t_piece.id piece_id,t_piece.`name` piece_name").
		Joins("INNER JOIN t_section ON t_section.id = t_piece.section_id").
		Joins("INNER JOIN t_chapter ON t_chapter.id = t_section.chapter_id").
		Joins("INNER JOIN t_subject ON t_subject.subject_id = t_chapter.subject_id").
		Joins("INNER JOIN t_package ON t_package.package_id = t_subject.package_id").
		Where("t_piece.id = ?", pieceId).Scan(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询课程信息失败")
	}
	return data, nil
}

func (a *ApiSubject) GetSubjectPiece(subjectId []int64) ([]mysql_model.SubjectByPieceId, error) {
	subject := make([]mysql_model.SubjectByPieceId, 0, 0)
	err := gormx.DB.Table("t_subject").Select("t_subject.subject_id, t_subject.subject_name, t_piece.id, t_piece.name, t_chapter.subject_id").
		Joins("LEFT JOIN t_chapter ON t_subject.subject_id = t_chapter.subject_id").
		Joins("LEFT JOIN t_section ON t_section.chapter_id = t_chapter.id").
		Joins("LEFT JOIN t_piece ON t_piece.section_id = t_section.id").
		Where("t_subject.subject_id in ?", subjectId).Scan(&subject).Error
	if err != nil {
		return subject, errors.Wrap(err, "查询课程下的节块信息失败")
	}
	return subject, nil
}

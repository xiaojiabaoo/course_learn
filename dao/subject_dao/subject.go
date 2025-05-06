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
	GetSubjectByPieceId(pieceId int64) ([]mysql_model.SubjectByPiece, error)
	GetSubjectPiece(subjectId []int64) ([]mysql_model.SubjectByPieceId, error)
}

type ApiSubject struct{}

func (a *ApiSubject) GetSubjectData(userId, types, time int64) ([]mysql_model.TSubjectChapter, error) {
	subject := make([]mysql_model.TSubjectChapter, 0, 0)
	joins := gormx.DB.Table("t_subject s").Select("s.knowledge_tree_id, s.subject_id, s.subject_name, " +
		"s.package_id, uo.type, uo.user_id, uo.expiration_time").
		Joins("LEFT JOIN user_order uo ON uo.subject_id = s.subject_id")
	if types == 1 || types == 3 {
		joins = joins.Where("uo.user_id=? and is_valid=? and uo.expiration_time > ?", userId, 1, time)
	}
	if types == 4 {
		joins = joins.Where("s.is_hot=?", 1)
	}
	joins = joins.Where("s.status=1")
	err := joins.Order("s.sort asc").Order("uo.expiration_time desc").Scan(&subject).Error
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
	err := gormx.DB.Where("subject_id=? and status=1", subjectId).First(subject).Error
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
	err := gormx.DB.Debug().Table("t_piece pi").
		Select("p.package_id,p.package_name,s.subject_id,s.subject_name,"+
			"c.id chapter_id, c.`name` chapter_name,se.id section_id, se.`name` section_name,"+
			"pi.id piece_id,pi.`name` piece_name").
		Joins("INNER JOIN t_section se ON se.id = pi.section_id").
		Joins("INNER JOIN t_chapter c ON c.id = se.chapter_id").
		Joins("INNER JOIN t_subject s ON s.subject_id = c.subject_id").
		Joins("INNER JOIN t_package p ON p.package_id = s.package_id").
		Where("pi.id = ?", pieceId).
		Where("pi.status=1 and se.status=1 and c.status=1 and s.status=1 and p.status=1").
		Scan(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询课程信息失败")
	}
	return data, nil
}

func (a *ApiSubject) GetSubjectPiece(subjectId []int64) ([]mysql_model.SubjectByPieceId, error) {
	subject := make([]mysql_model.SubjectByPieceId, 0, 0)
	err := gormx.DB.Table("t_subject s").Select("s.subject_id, s.subject_name, p.id, p.name, c.subject_id").
		Joins("LEFT JOIN t_chapter c ON s.subject_id = c.subject_id").
		Joins("LEFT JOIN t_section se ON se.chapter_id = c.id").
		Joins("LEFT JOIN t_piece p ON p.section_id = se.id").
		Where("s.subject_id in ?", subjectId).
		Where("se.status=1 and c.status=1 and s.status=1 and p.status=1").
		Scan(&subject).Error
	if err != nil {
		return subject, errors.Wrap(err, "查询课程下的节块信息失败")
	}
	return subject, nil
}

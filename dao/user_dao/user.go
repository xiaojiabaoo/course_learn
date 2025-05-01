package user_dao

import (
	"course_learn/common/gormx"
	mysql_model "course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User interface {
	AddUsers(*mysql_model.Users) error
	AddUserParam(*mysql_model.UserParam) error
	AddUserLoginLogs(*mysql_model.UserLoginLogs) error
	AddUserLeanRecord(mysql_model.UserLearnRecord) error
	GetUserHasSub(subjectId, userId, time int64) (mysql_model.UserOrder, error)
	GetUserLeanRecord(userId int64, subjectId []int64) ([]mysql_model.UserLearnRecord, error)
	GetUserData(userId int64) (mysql_model.Users, error)
	SetUserData(mysql_model.Users) error
	ChangePassword(mysql_model.Users) error
	GetUserLoginRecord(param req_model.GetLoginRecord) ([]mysql_model.UserLoginLogs, int64, error)
}

type ApiUser struct{}

func (a *ApiUser) AddUsers(user *mysql_model.Users) error {
	err := gormx.DB.Create(user).Error
	if err != nil {
		return errors.Wrap(err, "添加用户失败")
	}
	return nil
}

func (a *ApiUser) AddUserParam(param *mysql_model.UserParam) error {
	err := gormx.DB.Create(param).Error
	if err != nil {
		return errors.Wrap(err, "添加用户参数失败")
	}
	return nil
}

func (a *ApiUser) AddUserLoginLogs(logs *mysql_model.UserLoginLogs) error {
	err := gormx.DB.Create(logs).Error
	if err != nil {
		return errors.Wrap(err, "添加用户登录日志失败")
	}
	return nil
}

func (a *ApiUser) GetUserParam(id int64) (*mysql_model.UserParam, error) {
	param := &mysql_model.UserParam{}
	err := gormx.DB.Where("user_id=?", id).First(param).Error
	if err != nil {
		return param, errors.Wrap(err, "获取有用户参数信息失败")
	}
	return param, nil
}

func (a *ApiUser) AddUserLeanRecord(param mysql_model.UserLearnRecord) error {
	err := gormx.DB.Create(&param).Error
	if err != nil {
		return errors.Wrap(err, "添加用户学习记录失败")
	}
	return nil
}

func (a *ApiUser) GetUserHasSub(subjectId, userId, time int64) (mysql_model.UserOrder, error) {
	order := mysql_model.UserOrder{}
	err := gormx.DB.Where("subject_id=? and user_id=? and is_valid=1 and expiration_time>?",
		subjectId, userId, time).First(&order).Error
	if err != nil {
		return order, errors.Wrap(err, "查询用户与产品关系信息失败")
	}
	return order, err
}

func (a *ApiUser) GetUserLeanRecord(userId int64, subjectId []int64) ([]mysql_model.UserLearnRecord, error) {
	order := make([]mysql_model.UserLearnRecord, 0, 0)
	err := gormx.DB.Where("user_id=? and subject_id in ?", userId, subjectId).Find(&order).Error
	if err != nil {
		return order, errors.Wrap(err, "查询用户学习记录信息失败")
	}
	return order, err
}

func (a *ApiUser) GetUserData(userId int64) (mysql_model.Users, error) {
	users := mysql_model.Users{}
	err := gormx.DB.Where("id=?", userId).First(&users).Error
	if err != nil {
		return users, errors.Wrap(err, "查询用户信息失败")
	}
	return users, nil
}

func (a *ApiUser) SetUserData(param mysql_model.Users) error {
	err := gormx.DB.Table("users").Where("id=?", param.Id).Updates(map[string]interface{}{
		"nick_name": param.NickName,
		"gender":    param.Gender,
		"logo":      param.Logo,
		"phone":     param.Phone,
		"email":     param.Email,
		"area":      param.Area,
	}).Error
	if err != nil {
		return errors.Wrap(err, "修改用户信息失败")
	}
	return nil
}

func (a *ApiUser) ChangePassword(param mysql_model.Users) error {
	err := gormx.DB.Table("users").Where("id=?", param.Id).Updates(map[string]interface{}{
		"password": param.Password,
	}).Error
	if err != nil {
		return errors.Wrap(err, "修改用户密码失败")
	}
	return nil
}

func (a *ApiUser) GetUserLoginRecord(param req_model.GetLoginRecord) ([]mysql_model.UserLoginLogs, int64, error) {
	var (
		total int64
		err   error
		data  = make([]mysql_model.UserLoginLogs, 0)
	)
	err = gormx.DB.Table("user_login_logs").Where("user_id=?", param.UserId).Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户登录日志失败")
	}
	err = gormx.DB.Table("user_login_logs").Where("user_id=?", param.UserId).Order("login_time desc").
		Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户登录日志失败")
	}
	return data, total, nil
}

func (a *ApiUser) GetOrderList(param req_model.GetOrderList) ([]mysql_model.UserOrderList, int64, error) {
	var (
		total int64
		err   error
		data  = make([]mysql_model.UserOrderList, 0)
	)
	db := gormx.DB.Debug().Table("user_order uo").Select("uo.*,ts.subject_name").
		Joins("LEFT JOIN t_subject ts ON uo.subject_id=ts.subject_id").
		Where("uo.user_id=?", param.UserId)
	if param.SubjectId > 0 {
		db = db.Where("uo.subject_id=?", param.SubjectId)
	}
	db = db.Order("uo.add_time desc")
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户订单列表失败")
	}
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户订单列表失败")
	}
	return data, total, nil
}

func (a *ApiUser) GetLearnRecord(param req_model.GetLearnRecord) ([]mysql_model.UserLearnRecord, int64, error) {
	var (
		total int64
		err   error
		data  = make([]mysql_model.UserLearnRecord, 0)
	)
	db := gormx.DB.Debug().Table("user_learn_record").
		Select("package_name,subject_name,chapter_name,section_name,piece_name,MAX(add_time) add_time").
		Where("user_id=?", param.UserId)
	sql := fmt.Sprintf(`SELECT 
			package_name,subject_name,chapter_name,section_name,piece_name,MAX(add_time) add_time 
			FROM user_learn_record WHERE user_id=2`)
	if param.PackageId > 0 {
		sql += fmt.Sprintf(` and package_id=%d`, param.PackageId)
		db = db.Where("package_id=?", param.PackageId)
	}
	if param.SubjectId > 0 {
		sql += fmt.Sprintf(` and subject_id=%d`, param.SubjectId)
		db = db.Where("subject_id=?", param.SubjectId)
	}
	if param.ChapterId > 0 {
		sql += fmt.Sprintf(` and chapter_id=%d`, param.ChapterId)
		db = db.Where("chapter_id=?", param.ChapterId)
	}
	if param.SectionId > 0 {
		sql += fmt.Sprintf(` and section_id=%d`, param.SectionId)
		db = db.Where("section_id=?", param.SectionId)
	}
	if param.PieceId > 0 {
		sql += fmt.Sprintf(` and piece_id=%d`, param.PieceId)
		db = db.Where("piece_id=?", param.PieceId)
	} else {
		sql += fmt.Sprintf(` and piece_id is not null and piece_id > 0 `)
		db = db.Where("piece_id is not null and piece_id > 0")
	}
	sql += fmt.Sprintf(`GROUP BY package_name,subject_name,chapter_name,section_name,piece_name`)
	err = gormx.DB.Debug().Raw(fmt.Sprintf(`SELECT COUNT(1) FROM (%s) temp`, sql)).Scan(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户学习记录失败")
	}
	db = db.Group("package_name,subject_name,chapter_name,section_name,piece_name")
	db = db.Order("add_time desc")
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户学习记录失败")
	}
	return data, total, nil
}

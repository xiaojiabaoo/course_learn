package user_dao

import (
	"course_learn/common/gormx"
	mysql_model "course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type User interface {
	GetUserByIds(userIds []int64) ([]mysql_model.Users, error)
	GetUserByPhone(phone string) ([]mysql_model.Users, error)
	GetUserByEmail(email string) ([]mysql_model.Users, error)
	GetUserByNickName(nickName string) ([]mysql_model.Users, error)
	AddUsers(user *mysql_model.Users) error
	AddUserParam(param *mysql_model.UserParam) error
	AddUserLoginLogs(logs *mysql_model.UserLoginLogs) error
	GetUserParam(id int64) (*mysql_model.UserParam, error)
	AddUserLeanRecord(mysql_model.UserLearnRecord) error
	GetUserHasSub(subjectId, userId, time int64) (mysql_model.UserOrder, error)
	GetUserLeanRecord(userId int64, subjectId []int64) ([]mysql_model.UserLearnRecord, error)
	GetUserData(userId int64) (mysql_model.Users, error)
	SetUserData(mysql_model.Users) error
	ChangePassword(param mysql_model.Users) error
	GetUserLoginRecord(param req_model.GetLoginRecord) ([]resp_model.UserLoginLogs, int64, error)
	GetOrderList(param req_model.GetOrderList) ([]mysql_model.UserOrderList, int64, error)
	GetLearnRecord(param req_model.GetLearnRecord) ([]mysql_model.UserLearnRecord, int64, error)
	GetUserList(param req_model.GetUserList) ([]mysql_model.Users, int64, error)
	EditUserData(param req_model.EditUserData) error
}

type ApiUser struct{}

func (a *ApiUser) GetUserByIds(userIds []int64) ([]mysql_model.Users, error) {
	var (
		users = make([]mysql_model.Users, 0)
		err   error
	)
	err = gormx.DB.Debug().Table("users").Where("id in ?", userIds).Find(&users).Error
	if err != nil {
		return users, errors.Wrap(err, "查询用户失败")
	}
	return users, nil
}

func (a *ApiUser) GetUserByPhone(phone string) ([]mysql_model.Users, error) {
	users := make([]mysql_model.Users, 0)
	err := gormx.DB.Where("phone like ?", "%"+phone+"%").Find(&users).Error
	if err != nil && err.Error() != "record not found" {
		return nil, errors.Wrap(err, "根据手机号码获取用户信息失败")
	}
	return users, nil
}

func (a *ApiUser) GetUserByEmail(email string) ([]mysql_model.Users, error) {
	users := make([]mysql_model.Users, 0)
	err := gormx.DB.Where("email like ?", "%"+email+"%").Find(&users).Error
	if err != nil && err.Error() != "record not found" {
		return nil, errors.Wrap(err, "根据邮箱获取用户信息失败")
	}
	return users, nil
}

func (a *ApiUser) GetUserByNickName(nickName string) ([]mysql_model.Users, error) {
	users := make([]mysql_model.Users, 0)
	err := gormx.DB.Where("nick_name like ?", "%"+nickName+"%").Find(&users).Error
	if err != nil && err.Error() != "record not found" {
		return nil, errors.Wrap(err, "根据昵称获取用户信息失败")
	}
	return users, nil
}

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

func (a *ApiUser) GetUserLoginRecord(param req_model.GetLoginRecord) ([]resp_model.UserLoginLogs, int64, error) {
	var (
		total int64
		err   error
		data  = make([]resp_model.UserLoginLogs, 0)
	)
	db := gormx.DB.Table("user_login_logs ull").
		Select("ull.*,u.nick_name,u.email,u.phone").
		Joins("LEFT JOIN users u ON ull.user_id = u.id")
	if param.UserId > 0 {
		db = db.Where("ull.user_id=?", param.UserId)
	}
	if len(param.NickName) > 0 {
		db = db.Where("nick_name like ?", "%"+param.NickName+"%")
	}
	if len(param.Email) > 0 {
		db = db.Where("email like ?", "%"+param.Email+"%")
	}
	if len(param.Phone) > 0 {
		db = db.Where("phone like ?", "%"+param.Phone+"%")
	}
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户登录日志失败")
	}
	err = db.Order("ull.login_time desc").
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
	db := gormx.DB.Debug().Table("user_order uo").
		Select("uo.*,ts.subject_name,u.nick_name,u.email,u.phone").
		Joins("LEFT JOIN t_subject ts ON uo.subject_id=ts.subject_id").
		Joins("LEFT JOIN users u ON u.id=uo.user_id")
	if param.UserId > 0 {
		db = db.Where("uo.user_id=?", param.UserId)
	}
	if param.SubjectId > 0 {
		db = db.Where("uo.subject_id=?", param.SubjectId)
	}
	if len(param.NickName) > 0 {
		db = db.Where("nick_name like ?", "%"+param.NickName+"%")
	}
	if len(param.Email) > 0 {
		db = db.Where("email like ?", "%"+param.Email+"%")
	}
	if len(param.Phone) > 0 {
		db = db.Where("phone like ?", "%"+param.Phone+"%")
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
		Select("user_id,package_name,subject_name,chapter_name,section_name,piece_name,MAX(add_time) add_time")
	sql := fmt.Sprintf(`SELECT 
			user_id,package_name,subject_name,chapter_name,section_name,piece_name,MAX(add_time) add_time 
			FROM user_learn_record WHERE 1=1 `)
	// user 参数
	if param.UserId > 0 {
		sql += fmt.Sprintf(` and user_id=%d`, param.UserId)
		db = db.Where("user_id=?", param.UserId)
	}
	// manager 参数
	if len(param.UserIds) > 0 {
		if param.UserIds[0] == -1 {
			db = db.Where("user_id in ?", param.UserIds)
			sql += fmt.Sprintf(` and user_id in (%d)`, -1)
		} else {
			db = db.Where("user_id in ?", param.UserIds)
			userIds := ""
			for k, v := range param.UserIds {
				userIds += fmt.Sprintf(`%d`, v)
				if k != len(param.UserIds)-1 {
					userIds += ","
				}
			}
			sql += fmt.Sprintf(` and user_id in (%s)`, userIds)
		}
	}
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
	sql += fmt.Sprintf(`GROUP BY user_id,package_name,subject_name,chapter_name,section_name,piece_name`)
	err = gormx.DB.Debug().Raw(fmt.Sprintf(`SELECT COUNT(1) FROM (%s) temp`, sql)).Scan(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户学习记录失败")
	}
	db = db.Group("user_id,package_name,subject_name,chapter_name,section_name,piece_name")
	db = db.Order("add_time desc")
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户学习记录失败")
	}
	return data, total, nil
}

func (a *ApiUser) GetUserList(param req_model.GetUserList) ([]mysql_model.Users, int64, error) {
	var (
		total int64
		err   error
		data  = make([]mysql_model.Users, 0)
	)
	db := gormx.DB.Debug().Table("users")
	if len(param.UserName) > 0 {
		db = db.Where("nick_name like ?", "%"+param.UserName+"%")
	}
	if len(param.Email) > 0 {
		db = db.Where("email like ?", "%"+param.Email+"%")
	}
	if len(param.Phone) > 0 {
		db = db.Where("phone like ?", "%"+param.Phone+"%")
	}
	if len(param.Gender) > 0 {
		db = db.Where("gender=?", param.Gender)
	}
	db = db.Order("add_time desc")
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户列表失败")
	}
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询用户列表失败")
	}
	return data, total, nil
}

func (a *ApiUser) EditUserData(param req_model.EditUserData) error {
	var (
		err error
	)
	err = gormx.DB.Debug().Table("users").Where("id=?", param.UserId).
		Updates(map[string]interface{}{
			"account_status": param.AccountStatus,
			"update_time":    time.Now().Unix(),
		}).Error
	if err != nil {
		return errors.Wrap(err, "修改用户信息失败")
	}
	return nil
}

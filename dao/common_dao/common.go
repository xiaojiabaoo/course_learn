package common_dao

import (
	"course_learn/common/gormx"
	do "course_learn/model/mysql_model"
	"github.com/pkg/errors"
)

type Common interface {
	GetUserDataByAccPass(account, password string) (*do.Users, error)
	GetUserDataByAccount(account string) (*do.Users, error)
}

type ApiCommon struct{}

func (a *ApiCommon) GetUserDataByAccPass(account, password string) (*do.Users, error) {
	users := &do.Users{}
	err := gormx.DB.Where("password=? and (phone=? or email=?)", password, account, account).First(users).Error
	if err != nil && err.Error() != "record not found" {
		return nil, errors.Wrap(err, "根据账号和密码获取用户信息失败")
	}
	return users, nil
}

func (a *ApiCommon) GetUserDataByAccount(account string) (*do.Users, error) {
	users := &do.Users{}
	err := gormx.DB.Where("phone=? or email=?", account, account).First(users).Error
	if err != nil && err.Error() != "record not found" {
		return nil, errors.Wrap(err, "根据账号获取用户信息失败")
	}
	return users, nil
}

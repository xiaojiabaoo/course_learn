package user_ser

import (
	"course_learn/common/redisx"
	"course_learn/dao/subject_dao"
	"course_learn/dao/user_dao"
	"course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/tools"
	custErr "course_learn/tools/error"
	"course_learn/tools/util"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type User interface {
	GetUserLeanRecord(req_model.GetUserLeanRecord) ([]resp_model.GetUserLeanRecord, error)
	AddUserLeanRecord(req_model.AddUserLeanRecord) error
	GetUserData(req_model.GetUserData) (resp_model.GetUserData, error)
	SetUserData(req_model.SetUserData) error
	ChangePassword(req_model.ChangePassword) error
	GetLoginRecord(param req_model.GetLoginRecord) (resp_model.GetLoginRecordData, error)
	GetOrderList(param req_model.GetOrderList) (resp_model.GetOrderListData, error)
	GetLearnRecord(param req_model.GetLearnRecord) (resp_model.GetLearnRecordData, error)
	GetUserList(param req_model.GetUserList) (resp_model.GetUserListData, error)
	EditUserData(param req_model.EditUserData) error
}

type ApiUser struct{}

func (a *ApiUser) GetUserLeanRecord(param req_model.GetUserLeanRecord) ([]resp_model.GetUserLeanRecord, error) {
	var (
		err  error
		list = make([]resp_model.GetUserLeanRecord, 0, 0)
	)
	if param.UserId == tools.EMPTY_INT {
		return list, custErr.New(custErr.USER_ID_EMPTY_ERROR, "")
	}
	return list, err
}

func (a *ApiUser) AddUserLeanRecord(param req_model.AddUserLeanRecord) error {
	var (
		err        error
		apiUser    = user_dao.ApiUser{}
		apiSubject = subject_dao.ApiSubject{}
		now        = time.Now()
		piece      = make([]mysql_model.SubjectByPiece, 0, 0)
	)
	if param.UserId == tools.EMPTY_INT {
		return custErr.New(custErr.USER_ID_EMPTY_ERROR, "")
	}
	if param.PieceId == tools.EMPTY_INT {
		return custErr.New(custErr.USER_PIECEID_EMPTY_ERROR, "")
	}
	piece, err = apiSubject.GetSubjectByPieceId(param.PieceId)
	if err != nil {
		return err
	}
	if len(piece) == 0 {
		return nil
	}
	record := mysql_model.UserLearnRecord{
		UserId:      param.UserId,
		PackageId:   piece[0].PackageId,
		PackageName: piece[0].PackageName,
		SubjectId:   piece[0].SubjectId,
		SubjectName: piece[0].SubjectName,
		ChapterId:   piece[0].ChapterId,
		ChapterName: piece[0].ChapterName,
		SectionId:   piece[0].SectionId,
		SectionName: piece[0].SectionName,
		PieceId:     piece[0].PieceId,
		PieceName:   piece[0].PieceName,
		AddTime:     now.Unix(),
		DateTime:    now,
	}
	err = apiUser.AddUserLeanRecord(record)
	if err != nil {
		return err
	}
	return nil
}

func (a *ApiUser) GetUserData(param req_model.GetUserData) (resp_model.GetUserData, error) {
	var (
		data    = resp_model.GetUserData{}
		err     error
		apiUser = user_dao.ApiUser{}
		users   = mysql_model.Users{}
	)
	if param.UserId == tools.EMPTY_INT {
		return data, custErr.New(custErr.USER_ID_EMPTY_ERROR, "")
	}
	users, err = apiUser.GetUserData(param.UserId)
	if err != nil {
		return data, err
	}
	data.NickName = users.NickName
	data.Logo = users.Logo
	data.Gender = users.Gender
	data.Phone = users.Phone
	data.Email = users.Email
	data.Area = users.Area
	data.AddTime = users.AddTime
	data.AccountStatus = users.AccountStatus
	return data, nil
}

func (a *ApiUser) SetUserData(param req_model.SetUserData) error {
	var (
		apiUser = user_dao.ApiUser{}
	)
	if param.UserId == tools.EMPTY_INT {
		return custErr.New(custErr.USER_ID_EMPTY_ERROR, "")
	}
	users := mysql_model.Users{
		Id:       param.UserId,
		NickName: param.NickName,
		Gender:   param.Gender,
		Logo:     param.Logo,
		Phone:    param.Phone,
		Email:    param.Email,
		Area:     param.Area,
	}
	return apiUser.SetUserData(users)
}

func (a *ApiUser) ChangePassword(param req_model.ChangePassword) error {
	var (
		apiUser = user_dao.ApiUser{}
		err     error
		user    = mysql_model.Users{}
		code    string
	)
	if param.UserId == tools.EMPTY_INT {
		return custErr.New(custErr.USER_ID_EMPTY_ERROR, "")
	}
	if param.SecurityCode <= 0 {
		return custErr.New(custErr.SECURITY_EMPTY_ERROR, "")
	}
	if param.ConfirmPassword != param.NewPassword {
		return custErr.New(custErr.USER_CONFIRM_PASSWORD_ERROR, "")
	}
	user, err = apiUser.GetUserData(param.UserId)
	if err != nil {
		return err
	}
	// 检测验证码是否正确
	code = redisx.Redis.Get(fmt.Sprintf(tools.SECURITY_CODE_KEY, user.Email))
	if err != nil {
		return errors.Wrap(err, "获取验证码失败")
	}
	if param.SecurityCode != int64(util.Str2Int(code)) {
		//验证码错误
		return custErr.New(custErr.SECURITY_CODE_ERROR, "")
	}
	// 检测旧密码是否正确
	if user.Password != param.OldPassword {
		return custErr.New(custErr.USER_PASSWORD_ERROR, "")
	}

	err = redisx.Redis.Del(fmt.Sprintf(tools.SECURITY_CODE_KEY, user.Email))
	if err != nil {
		return errors.Wrap(err, "删除验证码失败")
	}
	// 验证码通过后开始修改密码
	err = apiUser.ChangePassword(mysql_model.Users{
		Id:       param.UserId,
		Password: fmt.Sprintf(tools.USER_DEFAULT_PASS, param.NewPassword),
	})
	if err != nil {
		return err
	}
	tokenUuid := redisx.Redis.Get(fmt.Sprintf(tools.USER_LOGIN_KEY, param.UserId))
	err = redisx.Redis.Del(fmt.Sprintf(tools.USER_LOGIN_TOKEN_KEY, tokenUuid))
	if err != nil {
		return errors.Wrap(err, "删除用户原有Token失败")
	}
	err = redisx.Redis.Del(fmt.Sprintf(tools.USER_LOGIN_KEY, param.UserId))
	if err != nil {
		return errors.Wrap(err, "删除用户原有Token失败")
	}
	return nil
}

func (a *ApiUser) GetLoginRecord(param req_model.GetLoginRecord) (resp_model.GetLoginRecordData, error) {
	var (
		apiUser = user_dao.ApiUser{}
		total   int64
		err     error
		data    = make([]resp_model.UserLoginLogs, 0)
		list    = resp_model.GetLoginRecordData{}
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiUser.GetUserLoginRecord(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, resp_model.LoginRecordData{
			Id:          v.Id,
			UserId:      v.UserId,
			LoginTime:   v.LoginTime,
			LoginIp:     v.LoginIp,
			LoginArea:   v.LoginArea,
			LoginMethod: v.LoginMethod,
			NickName:    v.NickName,
			Email:       v.Email,
			Phone:       v.Phone,
		})
	}
	return list, nil
}

func (a *ApiUser) GetOrderList(param req_model.GetOrderList) (resp_model.GetOrderListData, error) {
	var (
		apiUser = user_dao.ApiUser{}
		total   int64
		err     error
		data    = make([]mysql_model.UserOrderList, 0)
		list    = resp_model.GetOrderListData{}
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiUser.GetOrderList(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, resp_model.OrderListData{
			Id:             v.Id,
			OrderNo:        v.OrderNo,
			Type:           v.Type,
			IsPayment:      v.IsPayment,
			Amount:         v.Amount,
			BuyType:        v.BuyType,
			IsValid:        v.IsValid,
			ExpirationTime: v.ExpirationTime,
			UserId:         v.UserId,
			SubjectId:      v.SubjectId,
			SubjectName:    v.SubjectName,
			AddTime:        v.AddTime,
			NickName:       v.NickName,
			Phone:          v.Phone,
			Email:          v.Email,
		})
	}
	return list, nil
}

func (a *ApiUser) GetLearnRecord(param req_model.GetLearnRecord) (resp_model.GetLearnRecordData, error) {
	var (
		apiUser   = user_dao.ApiUser{}
		total     int64
		err       error
		data      = make([]mysql_model.UserLearnRecord, 0)
		list      = resp_model.GetLearnRecordData{}
		userIds   = make([]int64, 0)
		userIdMap = make(map[int64]int64)
		users     = make([]mysql_model.Users, 0)
		userMap   = make(map[int64]mysql_model.Users)
		index     int
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Action == tools.MANAGER_LOGIN {
		userIds = make([]int64, 0)
		users = make([]mysql_model.Users, 0)
		if len(param.NickName) > 0 {
			users, err = apiUser.GetUserByNickName(param.NickName)
			if err != nil {
				return list, err
			}
			for _, v := range users {
				userIds = append(userIds, v.Id)
			}
			index++
		}
		users = make([]mysql_model.Users, 0)
		if len(userIds) > 0 && len(param.Phone) > 0 {
			users, err = apiUser.GetUserByPhone(param.Phone)
			if err != nil {
				return list, err
			}
			for _, v := range users {
				userIds = append(userIds, v.Id)
			}
			index++
		}
		users = make([]mysql_model.Users, 0)
		if len(userIds) > 0 && len(param.Email) > 0 {
			users, err = apiUser.GetUserByEmail(param.Email)
			if err != nil {
				return list, err
			}
			for _, v := range users {
				userIds = append(userIds, v.Id)
			}
			if len(users) == 0 {
				userIds = make([]int64, 0)
			}
			index++
		}
		for _, v := range userIds {
			userIdMap[v] += 1
		}

		if index >= 1 {
			if index == 1 {
				param.UserIds = userIds
			} else {
				for id, v := range userIdMap {
					if v == 1 {
						index = -1
					} else if v >= 2 {
						param.UserIds = append(param.UserIds, id)
					}
				}
			}
			if len(param.UserIds) == 0 || index == -1 {
				param.UserIds = append(param.UserIds, -1)
			}
		}
	}
	data, total, err = apiUser.GetLearnRecord(param)
	if err != nil {
		return list, err
	}
	if param.Action == tools.MANAGER_LOGIN {
		userIds = make([]int64, 0)
		users = make([]mysql_model.Users, 0)
		// 查询用户信息
		for _, v := range data {
			userIds = append(userIds, v.UserId)
		}
		users, err = apiUser.GetUserByIds(userIds)
		if err != nil {
			return list, err
		}
		for _, v := range users {
			userMap[v.Id] = v
		}
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, resp_model.LearnRecordData{
			Id:          v.Id,
			UserId:      v.UserId,
			NickName:    userMap[v.UserId].NickName,
			Email:       userMap[v.UserId].Email,
			Phone:       userMap[v.UserId].Phone,
			PackageId:   v.PackageId,
			PackageName: v.PackageName,
			SubjectId:   v.SubjectId,
			SubjectName: v.SubjectName,
			ChapterId:   v.ChapterId,
			ChapterName: v.ChapterName,
			SectionId:   v.SectionId,
			SectionName: v.SectionName,
			PieceId:     v.PieceId,
			PieceName:   v.PieceName,
			AddTime:     v.AddTime,
		})
	}
	return list, nil
}

func (a *ApiUser) GetUserList(param req_model.GetUserList) (resp_model.GetUserListData, error) {
	var (
		apiUser = user_dao.ApiUser{}
		total   int64
		err     error
		data    = make([]mysql_model.Users, 0)
		list    = resp_model.GetUserListData{}
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiUser.GetUserList(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, resp_model.UserListData{
			Id:            v.Id,
			NickName:      v.NickName,
			Gender:        v.Gender,
			Logo:          v.Logo,
			AreaCode:      v.AreaCode,
			Phone:         v.Phone,
			Email:         v.Email,
			Password:      v.Password,
			Area:          v.Area,
			AddTime:       v.AddTime,
			UpdateTime:    v.UpdateTime,
			AccountStatus: v.AccountStatus,
		})
	}
	return list, nil
}

func (a *ApiUser) EditUserData(param req_model.EditUserData) error {
	var (
		apiUser = user_dao.ApiUser{}
	)
	if param.UserId <= 0 {
		return custErr.New(custErr.USER_ID_EMPTY_ERROR, "")
	}
	if param.AccountStatus > 0 {
		// 删除用户登录状态
		tokenUuid := redisx.Redis.Get(fmt.Sprintf(tools.USER_LOGIN_KEY, param.UserId))
		err := redisx.Redis.Del(fmt.Sprintf(tools.USER_LOGIN_TOKEN_KEY, tokenUuid))
		if err != nil {
			return errors.Wrap(err, "删除用户原有Token失败")
		}
		err = redisx.Redis.Del(fmt.Sprintf(tools.USER_LOGIN_KEY, param.UserId))
		if err != nil {
			return errors.Wrap(err, "删除用户原有Token失败")
		}
	}
	return apiUser.EditUserData(param)
}

package user

import (
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/service/common_ser"
	"course_learn/service/user_ser"
	"course_learn/tools"
	custErr "course_learn/tools/error"
	"course_learn/tools/response"
	"github.com/gin-gonic/gin"
)

func GetUserLeanRecord(c *gin.Context) {
	var (
		param       = req_model.GetUserLeanRecord{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	info, err := apiCommon.GetUserInfo(c)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	param.UserId = info
	data, err := apiQuestion.GetUserLeanRecord(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func AddUserLeanRecord(c *gin.Context) {
	var (
		param       = req_model.AddUserLeanRecord{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	info, err := apiCommon.GetUserInfo(c)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	param.UserId = info
	err = apiQuestion.AddUserLeanRecord(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetUserData(c *gin.Context) {
	var (
		param       = req_model.GetUserData{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetUserData{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	info, err := apiCommon.GetUserInfo(c)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	param.UserId = info
	data, err = apiQuestion.GetUserData(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func SetUserData(c *gin.Context) {
	var (
		param       = req_model.SetUserData{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	info, err := apiCommon.GetUserInfo(c)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	param.UserId = info
	err = apiQuestion.SetUserData(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func ChangePassword(c *gin.Context) {
	var (
		param       = req_model.ChangePassword{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	userId, err := apiCommon.GetUserInfo(c)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	param.UserId = userId
	err = apiQuestion.ChangePassword(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetLoginRecord(c *gin.Context) {
	var (
		param       = req_model.GetLoginRecord{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetLoginRecordData{}
	)
	err = c.ShouldBind(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	if param.Action == tools.USER_LOGIN || len(param.Action) == 0 {
		info, err := apiCommon.GetUserInfo(c)
		if err != nil {
			response.ResponseData(c, err, nil)
			return
		}
		param.UserId = info
	} else if param.Action == tools.MANAGER_LOGIN {
		_, err = apiCommon.GetManageInfo(c)
		if err != nil {
			response.ResponseData(c, err, nil)
			return
		}
	}
	data, err = apiQuestion.GetLoginRecord(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func GetOrderList(c *gin.Context) {
	var (
		param       = req_model.GetOrderList{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetOrderListData{}
	)
	err = c.ShouldBind(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	if param.Action == tools.USER_LOGIN || len(param.Action) == 0 {
		info, err := apiCommon.GetUserInfo(c)
		if err != nil {
			response.ResponseData(c, err, nil)
			return
		}
		param.UserId = info
	} else if param.Action == tools.MANAGER_LOGIN {
		_, err = apiCommon.GetManageInfo(c)
		if err != nil {
			response.ResponseData(c, err, nil)
			return
		}
	}
	data, err = apiQuestion.GetOrderList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func GetLearnRecord(c *gin.Context) {
	var (
		param       = req_model.GetLearnRecord{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetLearnRecordData{}
	)
	err = c.ShouldBind(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	if param.Action == tools.USER_LOGIN || len(param.Action) == 0 {
		info, err := apiCommon.GetUserInfo(c)
		if err != nil {
			response.ResponseData(c, err, nil)
			return
		}
		param.UserId = info
	} else if param.Action == tools.MANAGER_LOGIN {
		_, err = apiCommon.GetManageInfo(c)
		if err != nil {
			response.ResponseData(c, err, nil)
			return
		}
	}
	data, err = apiQuestion.GetLearnRecord(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func GetUserList(c *gin.Context) {
	var (
		param       = req_model.GetUserList{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetUserListData{}
	)
	err = c.ShouldBind(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	_, err = apiCommon.GetManageInfo(c)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	data, err = apiQuestion.GetUserList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func EditUserData(c *gin.Context) {
	var (
		param       = req_model.EditUserData{}
		err         error
		apiQuestion = user_ser.ApiUser{}
		apiCommon   = common_ser.ApiCommon{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	_, err = apiCommon.GetManageInfo(c)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	err = apiQuestion.EditUserData(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

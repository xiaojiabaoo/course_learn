package subject

import (
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/service/common_ser"
	"course_learn/service/subject_ser"
	"course_learn/tools"
	custErr "course_learn/tools/error"
	"course_learn/tools/response"
	"github.com/gin-gonic/gin"
)

func GetSubjectList(c *gin.Context) {
	var (
		param      = req_model.GetSubjectList{}
		err        error
		data       = make([]resp_model.GetSubjectList, 0, 0)
		apiSubject = subject_ser.ApiSubject{}
		apiCommon  = common_ser.ApiCommon{}
		info       int64
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	if param.Action == tools.USER_LOGIN || len(param.Action) == 0 {
		info, err = apiCommon.GetUserInfo(c)
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
	data, err = apiSubject.GetSubjectList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func TrySubject(c *gin.Context) {
	var (
		param      = req_model.TrySubject{}
		err        error
		apiSubject = subject_ser.ApiSubject{}
		apiCommon  = common_ser.ApiCommon{}
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
	err = apiSubject.TrySubject(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetSubjectDetail(c *gin.Context) {
	var (
		param      = req_model.GetSubjectDesc{}
		err        error
		apiSubject = subject_ser.ApiSubject{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	data, err := apiSubject.GetSubjectDetail(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

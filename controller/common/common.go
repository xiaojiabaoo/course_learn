package common

import (
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/service/common_ser"
	custErr "course_learn/tools/error"
	"course_learn/tools/response"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

func Verification(c *gin.Context) {
	var (
		param  = req_model.Verification{}
		err    error
		apiSer = common_ser.ApiCommon{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	err = apiSer.Verification(param)
	response.ResponseData(c, err, nil)
}

func Login(c *gin.Context) {
	var (
		param  = req_model.Login{}
		err    error
		apiSer = common_ser.ApiCommon{}
		login  = resp_model.Login{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	login, err = apiSer.Login(param)
	response.ResponseData(c, err, login)
}

func Upload(c *gin.Context) {
	var (
		err       error
		apiCommon = common_ser.ApiCommon{}
		file      = &multipart.FileHeader{}
		upload    = resp_model.Upload{}
		param     = req_model.Upload{}
	)
	err = c.ShouldBind(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	file, err = c.FormFile("file")
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	if file == nil {
		response.ResponseData(c, custErr.New(custErr.IMAGE_EMPTY_ERROR, err.Error()), nil)
		return
	}
	upload, err = apiCommon.Upload(c, file, param)
	response.ResponseData(c, err, upload)
}

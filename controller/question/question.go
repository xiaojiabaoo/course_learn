package question

import (
	"course_learn/model/req_model"
	"course_learn/service/common_ser"
	"course_learn/service/question_ser"
	custErr "course_learn/tools/error"
	"course_learn/tools/response"
	"github.com/gin-gonic/gin"
)

func GetQuestionData(c *gin.Context) {
	var (
		param       = req_model.GetQuestionData{}
		err         error
		apiQuestion = question_ser.ApiQuestion{}
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
	data, err := apiQuestion.GetQuestionData(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

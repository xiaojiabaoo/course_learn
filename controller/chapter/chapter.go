package chapter

import (
	"course_learn/model/req_model"
	"course_learn/service/chapter_ser"
	custErr "course_learn/tools/error"
	"course_learn/tools/response"
	"github.com/gin-gonic/gin"
)

func GetChapterList(c *gin.Context) {
	var (
		param      = req_model.GetChapterList{}
		err        error
		apiChapter = chapter_ser.ApiChapter{}
	)
	err = c.ShouldBindJSON(&param)
	if err != nil {
		response.ResponseData(c, custErr.New(custErr.INVALID_PARAMS, err.Error()), nil)
		return
	}
	data, err := apiChapter.GetChapterList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

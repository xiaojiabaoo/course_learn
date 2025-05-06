package manage

import (
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/service/common_ser"
	"course_learn/service/manage_ser"
	custErr "course_learn/tools/error"
	"course_learn/tools/response"
	"github.com/gin-gonic/gin"
)

func GetManageList(c *gin.Context) {
	var (
		param       = req_model.GetManageList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetManageList{}
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
	data, err = apiQuestion.GetManageList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func EditManage(c *gin.Context) {
	var (
		param       = req_model.EditManage{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.EditManage(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetPackageList(c *gin.Context) {
	var (
		param       = req_model.GetPackageList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetPackageList{}
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
	data, err = apiQuestion.GetPackageList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func GetPackageSelectList(c *gin.Context) {
	var (
		param       = req_model.GetPackageSelectList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = make([]resp_model.PackageListData, 0)
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
	data, err = apiQuestion.GetPackageSelectList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func EditPackage(c *gin.Context) {
	var (
		param       = req_model.EditPackage{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.EditPackage(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func DelPackage(c *gin.Context) {
	var (
		param       = req_model.DelPackage{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.DelPackage(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetSubjectList(c *gin.Context) {
	var (
		param       = req_model.GetSubjectsList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetSubjectsList{}
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
	data, err = apiQuestion.GetSubjectList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func GetSubjectSelectList(c *gin.Context) {
	var (
		param       = req_model.GetSubjectSelectList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = make([]resp_model.SubjectsListData, 0)
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
	data, err = apiQuestion.GetSubjectSelectList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func EditSubject(c *gin.Context) {
	var (
		param       = req_model.EditSubject{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.EditSubject(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func DelSubject(c *gin.Context) {
	var (
		param       = req_model.DelSubject{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.DelSubject(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetChapterList(c *gin.Context) {
	var (
		param       = req_model.GetChaptersList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetChaptersList{}
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
	data, err = apiQuestion.GetChapterList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func GetChapterSelectList(c *gin.Context) {
	var (
		param       = req_model.GetChapterSelectList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = make([]resp_model.ChaptersListData, 0)
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
	data, err = apiQuestion.GetChapterSelectList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func EditChapter(c *gin.Context) {
	var (
		param       = req_model.EditChapter{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.EditChapter(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func DelChapter(c *gin.Context) {
	var (
		param       = req_model.DelChapter{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.DelChapter(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetSectionList(c *gin.Context) {
	var (
		param       = req_model.GetSectionList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetSectionList{}
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
	data, err = apiQuestion.GetSectionList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func GetSectionSelectList(c *gin.Context) {
	var (
		param       = req_model.GetSectionSelectList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = make([]resp_model.SectionList, 0)
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
	data, err = apiQuestion.GetSectionSelectList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func EditSection(c *gin.Context) {
	var (
		param       = req_model.EditSection{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.EditSection(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func DelSection(c *gin.Context) {
	var (
		param       = req_model.DelSection{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.DelSection(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetPieceList(c *gin.Context) {
	var (
		param       = req_model.GetPieceList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetPieceList{}
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
	data, err = apiQuestion.GetPieceList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func GetPieceSelectList(c *gin.Context) {
	var (
		param       = req_model.GetPieceSelectList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = make([]resp_model.PieceListData, 0)
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
	data, err = apiQuestion.GetPieceSelectList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func EditPiece(c *gin.Context) {
	var (
		param       = req_model.EditPiece{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.EditPiece(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func DelPiece(c *gin.Context) {
	var (
		param       = req_model.DelPiece{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.DelPiece(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetQuestionList(c *gin.Context) {
	var (
		param       = req_model.GetQuestionList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = resp_model.GetQuestionList{}
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
	data, err = apiQuestion.GetQuestionList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func EditQuestion(c *gin.Context) {
	var (
		param       = req_model.EditQuestion{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.EditQuestion(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func DelQuestion(c *gin.Context) {
	var (
		param       = req_model.DelQuestion{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.DelQuestion(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func GetQuestionOptionList(c *gin.Context) {
	var (
		param       = req_model.GetQuestionOptionList{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
		apiCommon   = common_ser.ApiCommon{}
		data        = make([]resp_model.GetQuestionOptionList, 0)
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
	data, err = apiQuestion.GetQuestionOptionList(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, data)
}

func EditQuestionOption(c *gin.Context) {
	var (
		param       = req_model.EditQuestionOption{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.EditQuestionOption(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

func DelQuestionOption(c *gin.Context) {
	var (
		param       = req_model.EditQuestionOption{}
		err         error
		apiQuestion = manage_ser.ApiManage{}
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
	err = apiQuestion.DelQuestionOption(param)
	if err != nil {
		response.ResponseData(c, err, nil)
		return
	}
	response.ResponseData(c, nil, nil)
}

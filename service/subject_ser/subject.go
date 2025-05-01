package subject_ser

import (
	"course_learn/dao/subject_dao"
	"course_learn/dao/user_dao"
	do "course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/tools"
	custErr "course_learn/tools/error"
	"course_learn/tools/util"
	"time"
)

type Subject interface {
	GetSubjectList(req_model.GetSubjectList) ([]resp_model.GetSubjectList, error)
	TrySubject(req_model.TrySubject) error
	GetSubjectDetail(req_model.GetSubjectDesc) (resp_model.GetSubjectDetail, error)
}

type ApiSubject struct{}

func (a *ApiSubject) GetSubjectList(param req_model.GetSubjectList) ([]resp_model.GetSubjectList, error) {
	var (
		resp       = make([]resp_model.GetSubjectList, 0, 0)
		err        error
		apiSubject = subject_dao.ApiSubject{}
		apiUser    = user_dao.ApiUser{}
		tempMap    = make(map[int64]struct{})
		data       = make([]do.TSubjectChapter, 0, 0)
		now        = time.Now()
		subjectId  = make([]int64, 0, 0)
		totalMap   = make(map[int64]int64)
		recordMap  = make(map[int64]map[int64]int64)
		piece      = make([]do.SubjectByPieceId, 0, 0)
		record     = make([]do.UserLearnRecord, 0, 0)
	)
	//查询课程表
	data, err = apiSubject.GetSubjectData(param.UserId, param.Type, now.Unix())
	if err != nil {
		return resp, err
	}
	if param.Type == 3 {
		for _, v1 := range data {
			subjectId = append(subjectId, v1.TSubject.SubjectId)
		}
		if len(subjectId) == 0 {
			return resp, nil
		}
		//查询课程下的节块总数信息
		piece, err = apiSubject.GetSubjectPiece(subjectId)
		if err != nil {
			return resp, err
		}
		for _, v := range piece {
			totalMap[v.TSubject.SubjectId] += 1
		}
		//查询用户的学习记录
		record, err = apiUser.GetUserLeanRecord(param.UserId, subjectId)
		if err != nil {
			return resp, err
		}
		for _, v := range record {
			if _, ok := recordMap[v.SubjectId]; !ok {
				recordMap[v.SubjectId] = map[int64]int64{}
			}
			if _, ok := recordMap[v.SubjectId][v.PieceId]; !ok {
				recordMap[v.SubjectId][v.PieceId] = 1
			}
			recordMap[v.SubjectId][v.PieceId] = 1
		}
	}

	for _, v1 := range data {
		/*if v1.UserOrder.UserId != 0 && param.UserId != v1.UserOrder.UserId {
			continue
		}*/
		if _, ok := tempMap[v1.TSubject.SubjectId]; ok {
			continue
		}
		list := resp_model.GetSubjectList{
			KnowledgeTreeId: v1.TSubject.KnowledgeTreeId,
			SubjectId:       v1.TSubject.SubjectId,
			SubjectName:     v1.TSubject.SubjectName,
			PackageId:       v1.TSubject.PackageId,
			Sort:            v1.TSubject.Sort,
		}
		if param.Type == 3 {
			recordCount := len(recordMap[v1.TSubject.SubjectId])
			totalCount := totalMap[v1.TSubject.SubjectId]
			list.Progress = float64(recordCount) / float64(totalCount) * 100
		}
		if param.UserId == v1.UserOrder.UserId {
			if v1.UserOrder.ExpirationTime > now.Unix() {
				list.Status = v1.UserOrder.Type
				//如果距离过期时间不到五天，给出即将过期提示
				if v1.UserOrder.ExpirationTime-now.Unix() < tools.USER_EXPIRE_DAY {
					list.Expire = true
				}
			}

		}
		resp = append(resp, list)
		tempMap[v1.TSubject.SubjectId] = struct{}{}
	}
	return resp, nil
}

func (a *ApiSubject) TrySubject(param req_model.TrySubject) error {
	var (
		now        = time.Now()
		apiSubject = subject_dao.ApiSubject{}
		apiUser    = user_dao.ApiUser{}
		err        error
		userParam  = &do.UserParam{}
		list       = make([]do.UserOrder, 0, 0)
		number     = util.CreateSerialNumber()
	)
	if param.SubjectId == 0 {
		return custErr.New(custErr.SUBJECT_EMPTY_ERROR, "")
	}
	//获取到用户的可试用次数和使用时间
	userParam, err = apiUser.GetUserParam(param.UserId)
	if err != nil {
		return err
	}
	//获取用户的所有试用过的产品 包括正在试用的和过期的
	list, err = apiSubject.GetUserTrySubList(param)
	if err != nil {
		return err
	}
	//检测试用次数是否已达到了最高
	if userParam.TryCount <= int64(len(list)) {
		return custErr.New(custErr.USER_TRY_MAX_ERROR, "")
	}
	//检测购买的课程是否正在正常试用中
	for _, v := range list {
		if v.IsValid == 1 && v.ExpirationTime > now.Unix() && v.SubjectId == param.SubjectId {
			return custErr.New(custErr.USER_TRY_SAME_ERROR, "")
		}
	}
	//获取试用到期时间
	date := now.AddDate(0, 0, int(userParam.TryTime))
	order := do.UserOrder{
		OrderNo:        number,
		Type:           1,
		IsPayment:      0,
		BuyType:        0,
		IsValid:        1,
		ExpirationTime: date.Unix(),
		UserId:         param.UserId,
		SubjectId:      param.SubjectId,
		AddTime:        now.Unix(),
		UpTime:         now.Unix(),
		DateTime:       now,
	}
	err = apiSubject.AddTrySubject(order)
	if err != nil {
		return err
	}
	return nil
}

func (a *ApiSubject) GetSubjectDetail(param req_model.GetSubjectDesc) (resp_model.GetSubjectDetail, error) {
	var (
		err        error
		detail     = resp_model.GetSubjectDetail{}
		apiSubject = subject_dao.ApiSubject{}
		data       = &do.TSubject{}
	)
	if param.SubjectId == tools.EMPTY_INT {
		return detail, custErr.New(custErr.SUBJECT_EMPTY_ERROR, "")
	}
	data, err = apiSubject.GetSubjectDetail(param.SubjectId)
	if err != nil {
		return detail, err
	}
	detail.KnowledgeTreeId = data.KnowledgeTreeId
	detail.SubjectId = data.SubjectId
	detail.SubjectName = data.SubjectName
	detail.PackageId = data.PackageId
	detail.Sort = data.Sort
	detail.Desc = data.Desc
	return detail, nil
}

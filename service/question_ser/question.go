package question_ser

import (
	"course_learn/dao/question_dao"
	"course_learn/dao/user_dao"
	"course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/tools"
	custErr "course_learn/tools/error"
	"fmt"
	"time"
)

type Question interface {
	GetQuestionData(req_model.GetQuestionData) ([]resp_model.GetQuestionData, error)
}

type ApiQuestion struct{}

func (a ApiQuestion) GetQuestionData(param req_model.GetQuestionData) ([]resp_model.GetQuestionData, error) {
	var (
		err               error
		apiQuestion       = question_dao.ApiQuestion{}
		data              = make([]resp_model.GetQuestionData, 0, 0)
		questionData      = make([]mysql_model.TQuestionOption, 0, 0)
		childQuestionData = make([]mysql_model.TChildQuestionOption, 0, 0)
		optionMap         = make(map[int64][]resp_model.Option)
		temp1Map          = make(map[int64]struct{})
		temp2Map          = make(map[int64]struct{})
		temp3Map          = make(map[string]struct{})
		temp4Map          = make(map[string]struct{})
		apiUser           = user_dao.ApiUser{}
		now               = time.Now()
	)
	if param.PieceId == tools.EMPTY_INT {
		return data, custErr.New(custErr.PIECE_EMPTY_ERROR, "")
	}
	if param.SubjectId == tools.EMPTY_INT {
		return data, custErr.New(custErr.SUBJECT_EMPTY_ERROR, "")
	}
	//查询该课程当前登录用户是否试用或购买
	sub, err := apiUser.GetUserHasSub(param.SubjectId, param.UserId, now.Unix())
	if err != nil {
		return data, err
	}
	if sub.Id == 0 {
		return data, custErr.New(custErr.USER_NOT_TRY_ORBUY_ERROR,"")
	}
	questionData, err = apiQuestion.GetQuestionData(param.PieceId)
	if err != nil {
		return data, err
	}
	for _, v := range questionData {
		if v.TOption.Id == 0 {
			continue
		}
		key := fmt.Sprintf(`%d%s`, v.TQuestion.Id, v.TOption.Title)
		if _, ok := temp3Map[key]; ok {
			continue
		}
		optionMap[v.TQuestion.Id] = append(optionMap[v.TQuestion.Id], resp_model.Option{
			Id:         v.TOption.Id,
			QuestionId: v.TOption.QuestionId,
			Sequence:   v.TOption.Sequence,
			Title:      v.TOption.Title,
			Content:    v.TOption.Content,
			Correct:    v.TOption.Correct,
		})
		temp3Map[key] = struct{}{}
	}
	for _, v := range questionData {
		var childQuestion = make([]resp_model.ChildQuestionData, 0, 0)
		if _, ok := temp1Map[v.TQuestion.Id]; ok {
			continue
		}
		if v.TQuestion.QuestionType == "COMPREHENSIVE" {
			childQuestionData, err = apiQuestion.GetChildQuestionData(v.TQuestion.Id, param.PieceId)
			if err != nil {
				return data, err
			}
			var childOptionMap = make(map[int64][]resp_model.Option)
			for _, v2 := range childQuestionData {
				if v2.TOption.Id == 0 {
					continue
				}
				key := fmt.Sprintf(`%d%s`, v2.TChildQuestion.Id, v2.TOption.Title)
				if _, ok := temp4Map[key]; ok {
					continue
				}
				childOptionMap[v2.TChildQuestion.Id] = append(childOptionMap[v2.TChildQuestion.Id], resp_model.Option{
					Id:         v2.TOption.Id,
					QuestionId: v2.TOption.QuestionId,
					Sequence:   v2.TOption.Sequence,
					Title:      v2.TOption.Title,
					Content:    v2.TOption.Content,
					Correct:    v2.TOption.Correct,
				})
				temp4Map[key] = struct{}{}
			}
			for _, v2 := range childQuestionData {
				if _, ok := temp2Map[v2.TChildQuestion.Id]; ok {
					continue
				}
				c := resp_model.ChildQuestionData{
					Id:                v2.TChildQuestion.Id,
					PieceId:           v2.TChildQuestion.PieceId,
					Sequence:          v2.TChildQuestion.Sequence,
					QuestionType:      v2.TChildQuestion.QuestionType,
					QuestionSource:    v2.TChildQuestion.QuestionSource,
					QuestionContent:   v2.TChildQuestion.QuestionContent,
					QuestionAnswer:    v2.TChildQuestion.QuestionAnswer,
					Score:             v2.TChildQuestion.Score,
					MainNodeId:        v2.TChildQuestion.MainNodeId,
					MainNodeName:      v2.TChildQuestion.MainNodeName,
					MainNodeFrequency: v2.TChildQuestion.MainNodeFrequency,
					Favorite:          v2.TChildQuestion.Favorite,
					CanPhoto:          v2.TChildQuestion.CanPhoto,
					AvgCorrectRate:    v2.TChildQuestion.AvgCorrectRate,
					AnalysisType:      v2.TChildQuestion.AnalysisType,
					Analysis:          v2.TChildQuestion.Analysis,
				}
				if _, ok := childOptionMap[v2.TChildQuestion.Id]; ok {
					c.Option = childOptionMap[v2.TChildQuestion.Id]
				} else {
					c.Option = []resp_model.Option{}
				}
				childQuestion = append(childQuestion, c)
				temp2Map[v2.TChildQuestion.Id] = struct{}{}
			}
		}
		getQuestionData := resp_model.GetQuestionData{
			Id:                v.TQuestion.Id,
			PieceId:           v.TQuestion.PieceId,
			Sequence:          v.TQuestion.Sequence,
			QuestionType:      v.TQuestion.QuestionType,
			QuestionSource:    v.TQuestion.QuestionSource,
			QuestionContent:   v.TQuestion.QuestionContent,
			QuestionAnswer:    v.TQuestion.QuestionAnswer,
			Score:             v.TQuestion.Score,
			MainNodeId:        v.TQuestion.MainNodeId,
			MainNodeName:      v.TQuestion.MainNodeName,
			MainNodeFrequency: v.TQuestion.MainNodeFrequency,
			Favorite:          v.TQuestion.Favorite,
			CanPhoto:          v.TQuestion.CanPhoto,
			AvgCorrectRate:    v.TQuestion.AvgCorrectRate,
			AnalysisType:      v.TQuestion.AnalysisType,
			Analysis:          v.TQuestion.Analysis,
		}
		if _, ok := optionMap[v.TQuestion.Id]; ok {
			getQuestionData.Option = optionMap[v.TQuestion.Id]
		} else {
			getQuestionData.Option = []resp_model.Option{}
		}
		if len(childQuestion) > 0 {
			getQuestionData.ChildQuestionData = childQuestion
		} else {
			getQuestionData.ChildQuestionData = []resp_model.ChildQuestionData{}
		}
		data = append(data, getQuestionData)
		temp1Map[v.TQuestion.Id] = struct{}{}
	}
	return data, nil
}

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
		return data, custErr.New(custErr.USER_NOT_TRY_ORBUY_ERROR, "")
	}
	questionData, err = apiQuestion.GetQuestionData(param.PieceId)
	if err != nil {
		return data, err
	}
	for _, v := range questionData {
		if v.OptionId == 0 {
			continue
		}
		key := fmt.Sprintf(`%d%s`, v.Id, v.OptionTitle)
		if _, ok := temp3Map[key]; ok {
			continue
		}
		optionMap[v.Id] = append(optionMap[v.Id], resp_model.Option{
			Id:         v.OptionId,
			QuestionId: v.Id,
			Sequence:   v.OptionSequence,
			Title:      v.OptionTitle,
			Content:    v.OptionContent,
			Correct:    v.OptionCorrect,
		})
		temp3Map[key] = struct{}{}
	}
	for _, v := range questionData {
		var childQuestion = make([]resp_model.ChildQuestionData, 0, 0)
		if _, ok := temp1Map[v.Id]; ok {
			continue
		}
		if v.QuestionType == "COMPREHENSIVE" {
			childQuestionData, err = apiQuestion.GetChildQuestionData(v.Id, param.PieceId)
			if err != nil {
				return data, err
			}
			var childOptionMap = make(map[int64][]resp_model.Option)
			for _, v2 := range childQuestionData {
				if v2.OptionId == 0 {
					continue
				}
				key := fmt.Sprintf(`%d%s`, v2.Id, v2.OptionTitle)
				if _, ok := temp4Map[key]; ok {
					continue
				}
				childOptionMap[v2.Id] = append(childOptionMap[v2.Id], resp_model.Option{
					Id:         v2.OptionId,
					QuestionId: v2.Id,
					Sequence:   v2.OptionSequence,
					Title:      v2.OptionTitle,
					Content:    v2.OptionContent,
					Correct:    v2.OptionCorrect,
				})
				temp4Map[key] = struct{}{}
			}
			for _, v2 := range childQuestionData {
				if _, ok := temp2Map[v2.Id]; ok {
					continue
				}
				c := resp_model.ChildQuestionData{
					Id:                v2.Id,
					PieceId:           v2.PieceId,
					Sequence:          v2.Sequence,
					QuestionType:      v2.QuestionType,
					QuestionSource:    v2.QuestionSource,
					QuestionContent:   v2.QuestionContent,
					QuestionAnswer:    v2.QuestionAnswer,
					Score:             v2.Score,
					MainNodeId:        v2.MainNodeId,
					MainNodeName:      v2.MainNodeName,
					MainNodeFrequency: v2.MainNodeFrequency,
					Favorite:          v2.Favorite,
					CanPhoto:          v2.CanPhoto,
					AvgCorrectRate:    v2.AvgCorrectRate,
					AnalysisType:      v2.AnalysisType,
					Analysis:          v2.Analysis,
				}
				if _, ok := childOptionMap[v2.Id]; ok {
					c.Option = childOptionMap[v2.Id]
				} else {
					c.Option = []resp_model.Option{}
				}
				childQuestion = append(childQuestion, c)
				temp2Map[v2.Id] = struct{}{}
			}
		}
		getQuestionData := resp_model.GetQuestionData{
			Id:                v.Id,
			PieceId:           v.PieceId,
			Sequence:          v.Sequence,
			QuestionType:      v.QuestionType,
			QuestionSource:    v.QuestionSource,
			QuestionContent:   v.QuestionContent,
			QuestionAnswer:    v.QuestionAnswer,
			Score:             v.Score,
			MainNodeId:        v.MainNodeId,
			MainNodeName:      v.MainNodeName,
			MainNodeFrequency: v.MainNodeFrequency,
			Favorite:          v.Favorite,
			CanPhoto:          v.CanPhoto,
			AvgCorrectRate:    v.AvgCorrectRate,
			AnalysisType:      v.AnalysisType,
			Analysis:          v.Analysis,
		}
		if _, ok := optionMap[v.Id]; ok {
			getQuestionData.Option = optionMap[v.Id]
		} else {
			getQuestionData.Option = []resp_model.Option{}
		}
		if len(childQuestion) > 0 {
			getQuestionData.ChildQuestionData = childQuestion
		} else {
			getQuestionData.ChildQuestionData = []resp_model.ChildQuestionData{}
		}
		data = append(data, getQuestionData)
		temp1Map[v.Id] = struct{}{}
	}
	return data, nil
}

package chapter_ser

import (
	"course_learn/dao/chapter_dao"
	"course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/tools"
	custErr "course_learn/tools/error"
)

type Chapter interface {
	GetChapterList(req_model.GetChapterList) ([]resp_model.GetChapterList, error)
}

type ApiChapter struct{}

func (a ApiChapter) GetChapterList(param req_model.GetChapterList) ([]resp_model.GetChapterList, error) {
	var (
		err        error
		data       = make([]resp_model.GetChapterList, 0, 0)
		list       = make([]mysql_model.TChapterSection, 0, 0)
		chapter    = chapter_dao.ApiChapter{}
		sectionMap = make(map[int64][]resp_model.Section)
		pieceMap   = make(map[int64][]resp_model.Piece)
		temp1Map   = make(map[int64]struct{})
		temp2Map   = make(map[int64]struct{})
	)

	if param.SubjectId == tools.EMPTY_INT {
		return data, custErr.New(custErr.SUBJECT_EMPTY_ERROR, "")
	}
	list, err = chapter.GetChapterList(param)
	if err != nil {
		return data, err
	}

	for _, v := range list {
		pieceMap[v.TPiece.SectionId] = append(pieceMap[v.TPiece.SectionId], resp_model.Piece{
			Id:            v.TPiece.Id,
			Name:          v.TPiece.Name,
			QuestionCount: v.TPiece.QuestionCount,
			SectionId:     v.TPiece.SectionId,
		})
	}

	for _, v := range list {
		if _, ok := temp2Map[v.TSection.Id]; ok {
			continue
		}
		sectionMap[v.TSection.ChapterId] = append(sectionMap[v.TSection.ChapterId], resp_model.Section{
			Id:            v.TSection.Id,
			Sequence:      v.TSection.Sequence,
			Name:          v.TSection.Name,
			QuestionCount: v.TSection.QuestionCount,
			ChapterId:     v.TSection.ChapterId,
			Piece:         pieceMap[v.TSection.Id],
		})
		temp2Map[v.TSection.Id] = struct{}{}
	}
	for _, v := range list {
		if _, ok := temp1Map[v.TChapter.Id]; ok {
			continue
		}
		data = append(data, resp_model.GetChapterList{
			Id:            v.TChapter.Id,
			Sequence:      v.TChapter.Sequence,
			Name:          v.TChapter.Name,
			QuestionCount: v.TChapter.QuestionCount,
			SubjectId:     v.TChapter.SubjectId,
			Sort:          v.TChapter.Sort,
			Section:       sectionMap[v.TChapter.Id],
		})
		temp1Map[v.TChapter.Id] = struct{}{}
	}
	return data, nil
}

package manage_ser

import (
	"course_learn/common/redisx"
	"course_learn/dao/manage_dao"
	"course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/tools"
	custErr "course_learn/tools/error"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type Manage interface {
	GetManageList(param req_model.GetManageList) (resp_model.GetManageList, error)
	EditManage(param req_model.EditManage) error
	GetPackageList(param req_model.GetPackageList) (resp_model.GetPackageList, error)
	EditPackage(param req_model.EditPackage) error
	DelPackage(param req_model.DelPackage) error
	GetSubjectList(param req_model.GetSubjectsList) (resp_model.GetSubjectsList, error)
	EditSubject(param req_model.EditSubject) error
	DelSubject(param req_model.DelSubject) error
	GetChapterList(param req_model.GetChaptersList) (resp_model.GetChaptersList, error)
	GetChapterSelectList(param req_model.GetChapterSelectList) ([]resp_model.ChaptersListData, error)
	EditChapter(param req_model.EditChapter) error
	DelChapter(param req_model.DelChapter) error
	GetSectionList(param req_model.GetSectionList) (resp_model.GetSectionList, error)
	GetSectionSelectList(param req_model.GetSectionSelectList) ([]resp_model.SectionList, error)
	EditSection(param req_model.EditSection) error
	DelSection(param req_model.DelSection) error
	GetPieceList(param req_model.GetPieceList) (resp_model.GetPieceList, error)
	GetPieceSelectList(param req_model.GetPieceSelectList) ([]resp_model.PieceListData, error)
	EditPiece(param req_model.EditPiece) error
	DelPiece(param req_model.DelPiece) error
	GetQuestionList(param req_model.GetQuestionList) (resp_model.GetQuestionList, error)
	EditQuestion(param req_model.EditQuestion) error
	DelQuestion(param req_model.DelQuestion) error
	GetQuestionOptionList(param req_model.GetQuestionOptionList) ([]resp_model.GetQuestionOptionList, error)
	EditQuestionOption(param req_model.EditQuestionOption) error
	DelQuestionOption(param req_model.EditQuestionOption) error
}

type ApiManage struct{}

func (a *ApiManage) GetManageList(param req_model.GetManageList) (resp_model.GetManageList, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		total     int64
		err       error
		data      = make([]mysql_model.Manager, 0)
		list      = resp_model.GetManageList{}
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiManage.GetManageList(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, resp_model.ManageListData{
			Id:         v.Id,
			Account:    v.Account,
			Email:      v.Email,
			Password:   v.Password,
			Identifier: v.Identifier,
			Status:     v.Status,
			AddTime:    v.AddTime,
			UpdateTime: v.UpdateTime,
			DateTime:   v.DateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return list, nil
}

func (a *ApiManage) EditManage(param req_model.EditManage) error {
	var (
		now       = time.Now()
		manager   = mysql_model.Manager{}
		apiManage = manage_dao.ApiManage{}
		err       error
	)
	if param.Id <= 0 {
		// 新增
		manager, err = apiManage.GetManageByAccount(param.Account)
		if err != nil {
			return err
		}
		if manager.Id > 0 {
			return custErr.New(custErr.MANAGER_EXISTS, "")
		}
		manager.Account = param.Account
		manager.Password = param.Password
		manager.Identifier = param.Identifier
		manager.Email = param.Email
		manager.Status = 1
		manager.AddTime = now.Unix()
		manager.UpdateTime = now.Unix()
		manager.DateTime = now
		err = apiManage.AddManage(manager)
	} else {
		// 查询管理员是否存在
		manager, err = apiManage.GetManageById(param.Id)
		if err != nil {
			return err
		}
		if manager.Id <= 0 {
			return custErr.New(custErr.MANAGER_NOT_EXISTS, "")
		}
		manager.Password = param.Password
		manager.Identifier = param.Identifier
		manager.Email = param.Email
		manager.Status = param.Status
		err = apiManage.UpdateManage(manager)
	}
	if err != nil {
		return err
	}
	// 删除用户登录状态
	tokenUuid := redisx.Redis.Get(fmt.Sprintf(tools.MANAGER_LOGIN_KEY, param.Id))
	err = redisx.Redis.Del(fmt.Sprintf(tools.MANAGER_LOGIN_TOKEN_KEY, tokenUuid))
	if err != nil {
		return errors.Wrap(err, "删除用户原有Token失败")
	}
	err = redisx.Redis.Del(fmt.Sprintf(tools.MANAGER_LOGIN_KEY, param.Id))
	if err != nil {
		return errors.Wrap(err, "删除用户原有Token失败")
	}
	return nil
}

func (a *ApiManage) GetPackageList(param req_model.GetPackageList) (resp_model.GetPackageList, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		total     int64
		err       error
		data      = make([]mysql_model.TPackage, 0)
		list      = resp_model.GetPackageList{}
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiManage.GetPackageList(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, resp_model.PackageListData{
			PackageId:   v.PackageId,
			PackageName: v.PackageName,
			Status:      v.Status,
			AddTime:     v.AddTime,
			UpdateTime:  v.UpdateTime,
			DateTime:    v.DateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return list, nil
}

func (a *ApiManage) GetPackageSelectList(param req_model.GetPackageSelectList) ([]resp_model.PackageListData, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		err       error
		data      = make([]mysql_model.TPackage, 0)
		list      = make([]resp_model.PackageListData, 0)
	)
	data, err = apiManage.GetPackageSelectList(param)
	if err != nil {
		return list, err
	}
	for _, v := range data {
		list = append(list, resp_model.PackageListData{
			PackageId:   v.PackageId,
			PackageName: v.PackageName,
			AddTime:     v.AddTime,
			UpdateTime:  v.UpdateTime,
			DateTime:    v.DateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return list, nil
}

func (a *ApiManage) EditPackage(param req_model.EditPackage) error {
	var (
		packages  = mysql_model.TPackage{}
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	if param.PackageId <= 0 {
		packages.PackageName = param.PackageName
		packages.Status = 2
		packages.AddTime = now.Unix()
		packages.UpdateTime = now.Unix()
		packages.DateTime = now
		err = apiManage.AddPackage(packages)
	} else {
		packages, err = apiManage.GetPackageById(param.PackageId)
		if err != nil {
			return err
		}
		if packages.PackageId <= 0 {
			return custErr.New(custErr.PACKAGE_NOT_EXISTS, "")
		}
		packages.UpdateTime = now.Unix()
		packages.DateTime = now
		packages.PackageName = param.PackageName
		packages.Status = param.Status
		err = apiManage.UpdatePackage(packages)
	}
	return err
}

func (a *ApiManage) DelPackage(param req_model.DelPackage) error {
	var (
		packages  = mysql_model.TPackage{}
		subject   = make([]mysql_model.TSubject, 0)
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	packages, err = apiManage.GetPackageById(param.PackageId)
	if err != nil {
		return err
	}
	if packages.PackageId <= 0 {
		return custErr.New(custErr.PACKAGE_NOT_EXISTS, "")
	}
	subject, err = apiManage.GetSubjectByPackageId(param.PackageId)
	if err != nil {
		return err
	}
	if len(subject) > 0 {
		return custErr.New(custErr.PACKAGE_HAS_SUBJECT, "")
	}
	packages.Status = 3
	packages.UpdateTime = now.Unix()
	packages.DateTime = now
	return apiManage.UpdatePackage(packages)
}

func (a *ApiManage) GetSubjectList(param req_model.GetSubjectsList) (resp_model.GetSubjectsList, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		total     int64
		err       error
		data      = make([]resp_model.SubjectList, 0)
		list      = resp_model.GetSubjectsList{}
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiManage.GetSubjectList(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, resp_model.SubjectsListData{
			KnowledgeTreeId: v.KnowledgeTreeId,
			SubjectId:       v.SubjectId,
			SubjectName:     v.SubjectName,
			PackageId:       v.PackageId,
			PackageName:     v.PackageName,
			Sort:            v.Sort,
			Desc:            v.Desc,
			IsHot:           v.IsHot,
			Status:          v.Status,
			AddTime:         v.AddTime,
			UpdateTime:      v.UpdateTime,
		})
	}
	return list, nil
}

func (a *ApiManage) GetSubjectSelectList(param req_model.GetSubjectSelectList) ([]resp_model.SubjectsListData, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		err       error
		data      = make([]mysql_model.TSubject, 0)
		list      = make([]resp_model.SubjectsListData, 0)
	)
	data, err = apiManage.GetSubjectSelectList(param)
	if err != nil {
		return list, err
	}
	for _, v := range data {
		list = append(list, resp_model.SubjectsListData{
			KnowledgeTreeId: v.KnowledgeTreeId,
			SubjectId:       v.SubjectId,
			SubjectName:     v.SubjectName,
			PackageId:       v.PackageId,
			Sort:            v.Sort,
			Desc:            v.Desc,
			IsHot:           v.IsHot,
			Status:          v.Status,
			AddTime:         v.AddTime,
			UpdateTime:      v.UpdateTime,
		})
	}
	return list, nil
}

func (a *ApiManage) EditSubject(param req_model.EditSubject) error {
	var (
		subject   = mysql_model.TSubject{}
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	if param.PackageId <= 0 {
		return custErr.New(custErr.PACKAGE_NOT_EMPTY, "")
	}
	if len(param.SubjectName) <= 0 {
		return custErr.New(custErr.SUBJECT_NAME_NOT_EMPTY, "")
	}
	if param.SubjectId <= 0 {
		subject.PackageId = param.PackageId
		subject.SubjectId = param.SubjectId
		subject.SubjectName = param.SubjectName
		subject.Sort = param.Sort
		subject.Status = 2
		subject.Desc = param.Desc
		subject.IsHot = param.IsHot
		subject.AddTime = now.Unix()
		subject.UpdateTime = now.Unix()
		subject.DateTime = now
		err = apiManage.AddSubject(subject)
	} else {
		subject, err = apiManage.GetSubjectById(param.SubjectId)
		if err != nil {
			return err
		}
		if subject.SubjectId <= 0 {
			return custErr.New(custErr.SUBJECT_NOT_EXISTS, "")
		}
		subject.SubjectName = param.SubjectName
		subject.Sort = param.Sort
		subject.Status = param.Status
		subject.Desc = param.Desc
		subject.IsHot = param.IsHot
		subject.UpdateTime = now.Unix()
		subject.DateTime = now
		err = apiManage.UpdateSubject(subject)
	}
	return err
}

func (a *ApiManage) DelSubject(param req_model.DelSubject) error {
	var (
		subject   = mysql_model.TSubject{}
		chapter   = make([]mysql_model.TChapter, 0)
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	subject, err = apiManage.GetSubjectById(param.SubjectId)
	if err != nil {
		return err
	}
	if subject.SubjectId <= 0 {
		return custErr.New(custErr.SUBJECT_NOT_EXISTS, "")
	}
	chapter, err = apiManage.GetChapterBySubjectId(param.SubjectId)
	if err != nil {
		return err
	}
	if len(chapter) > 0 {
		return custErr.New(custErr.SUBJECT_HAS_CHAPTER, "")
	}
	subject.Status = 3
	subject.UpdateTime = now.Unix()
	subject.DateTime = now
	return apiManage.UpdateSubject(subject)
}

func (a *ApiManage) GetChapterList(param req_model.GetChaptersList) (resp_model.GetChaptersList, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		total     int64
		err       error
		data      = make([]resp_model.ChapterList, 0)
		list      = resp_model.GetChaptersList{}
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiManage.GetChapterList(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, resp_model.ChaptersListData{
			PackageId:     v.PackageId,
			PackageName:   v.PackageName,
			SubjectId:     v.SubjectId,
			SubjectName:   v.SubjectName,
			Id:            v.Id,
			Sequence:      v.Sequence,
			Name:          v.Name,
			QuestionCount: v.QuestionCount,
			Sort:          v.Sort,
			Status:        v.Status,
			AddTime:       v.AddTime,
			UpdateTime:    v.UpdateTime,
		})
	}
	return list, nil
}

func (a *ApiManage) GetChapterSelectList(param req_model.GetChapterSelectList) ([]resp_model.ChaptersListData, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		err       error
		data      = make([]mysql_model.TChapter, 0)
		list      = make([]resp_model.ChaptersListData, 0)
	)
	data, err = apiManage.GetChapterSelectList(param)
	if err != nil {
		return list, err
	}
	for _, v := range data {
		list = append(list, resp_model.ChaptersListData{
			SubjectId:     v.SubjectId,
			Id:            v.Id,
			Sequence:      v.Sequence,
			Name:          v.Name,
			QuestionCount: v.QuestionCount,
			Sort:          v.Sort,
			Status:        v.Status,
			AddTime:       v.AddTime,
			UpdateTime:    v.UpdateTime,
		})
	}
	return list, nil
}

func (a *ApiManage) EditChapter(param req_model.EditChapter) error {
	var (
		chapter   = mysql_model.TChapter{}
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	if param.Name == "" {
		return custErr.New(custErr.CHAPTER_NAME_NOT_EMPTY, "")
	}
	if param.Sequence == "" {
		return custErr.New(custErr.CHAPTER_SEQUENCE_NOT_EMPTY, "")
	}
	if param.Id <= 0 {
		if param.SubjectId <= 0 {
			return custErr.New(custErr.CHAPTER_SUBJECT_NOT_EMPTY, "")
		}
		chapter.SubjectId = param.SubjectId
		chapter.Sequence = param.Sequence
		chapter.Name = param.Name
		chapter.QuestionCount = param.QuestionCount
		chapter.Sort = param.Sort
		chapter.Status = 2
		chapter.AddTime = now.Unix()
		chapter.UpdateTime = now.Unix()
		chapter.DateTime = now
		err = apiManage.AddChapter(chapter)
	} else {
		chapter, err = apiManage.GetChapterById(param.Id)
		if err != nil {
			return err
		}
		if chapter.Id <= 0 {
			return custErr.New(custErr.CHAPTER_NOT_EXISTS, "")
		}
		chapter.Sequence = param.Sequence
		chapter.Name = param.Name
		chapter.QuestionCount = param.QuestionCount
		chapter.Sort = param.Sort
		chapter.Status = param.Status
		chapter.UpdateTime = now.Unix()
		chapter.DateTime = now
		err = apiManage.UpdateChapter(chapter)
	}
	return err
}

func (a *ApiManage) DelChapter(param req_model.DelChapter) error {
	var (
		chapter   = mysql_model.TChapter{}
		section   = make([]mysql_model.TSection, 0)
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	chapter, err = apiManage.GetChapterById(param.Id)
	if err != nil {
		return err
	}
	if chapter.Id <= 0 {
		return custErr.New(custErr.CHAPTER_NOT_EXISTS, "")
	}
	section, err = apiManage.GetSectionByChapterId(param.Id)
	if err != nil {
		return err
	}
	if len(section) > 0 {
		return custErr.New(custErr.CHAPTER_HAS_SECTION, "")
	}
	chapter.Status = 3
	chapter.UpdateTime = now.Unix()
	chapter.DateTime = now
	return apiManage.UpdateChapter(chapter)
}

func (a *ApiManage) GetSectionList(param req_model.GetSectionList) (resp_model.GetSectionList, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		total     int64
		err       error
		data      = make([]resp_model.SectionList, 0)
		list      = resp_model.GetSectionList{}
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiManage.GetSectionList(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, v)
	}
	return list, nil
}

func (a *ApiManage) GetSectionSelectList(param req_model.GetSectionSelectList) ([]resp_model.SectionList, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		err       error
		data      = make([]mysql_model.TSection, 0)
		list      = make([]resp_model.SectionList, 0)
	)
	data, err = apiManage.GetSectionSelectList(param)
	if err != nil {
		return list, err
	}
	for _, v := range data {
		list = append(list, resp_model.SectionList{
			ChapterId:     v.ChapterId,
			Id:            v.Id,
			Sequence:      v.Sequence,
			Name:          v.Name,
			QuestionCount: v.QuestionCount,
			Status:        v.Status,
			AddTime:       v.AddTime,
			UpdateTime:    v.UpdateTime,
		})
	}
	return list, nil
}

func (a *ApiManage) EditSection(param req_model.EditSection) error {
	var (
		section   = mysql_model.TSection{}
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	if len(param.Name) <= 0 {
		return custErr.New(custErr.SECTION_NAME_NOT_EMPTY, "")
	}
	if len(param.Sequence) <= 0 {
		return custErr.New(custErr.SECTION_SEQUENCE_NOT_EMPTY, "")
	}
	if param.Id <= 0 {
		if param.ChapterId <= 0 {
			return custErr.New(custErr.SECTION_CHAPTER_NOT_EMPTY, "")
		}
		section.Sequence = param.Sequence
		section.Name = param.Name
		section.QuestionCount = param.QuestionCount
		section.ChapterId = param.ChapterId
		section.Status = 2
		section.AddTime = now.Unix()
		section.UpdateTime = now.Unix()
		section.DateTime = now
		err = apiManage.AddSection(section)
	} else {
		section, err = apiManage.GetSectionById(param.Id)
		if err != nil {
			return err
		}
		if section.Id <= 0 {
			return custErr.New(custErr.SECTION_NOT_EXISTS, "")
		}
		section.Sequence = param.Sequence
		section.Name = param.Name
		section.QuestionCount = param.QuestionCount
		section.Status = param.Status
		section.UpdateTime = now.Unix()
		section.DateTime = now
		err = apiManage.UpdateSection(section)
	}
	return err
}

func (a *ApiManage) DelSection(param req_model.DelSection) error {
	var (
		piece     = make([]mysql_model.TPiece, 0)
		section   = mysql_model.TSection{}
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	section, err = apiManage.GetSectionById(param.Id)
	if err != nil {
		return err
	}
	if section.Id <= 0 {
		return custErr.New(custErr.SECTION_NOT_EXISTS, "")
	}
	piece, err = apiManage.GetPieceBySectionId(param.Id)
	if err != nil {
		return err
	}
	if len(piece) > 0 {
		return custErr.New(custErr.SECTION_HAS_PIECE, "")
	}
	section.Status = 3
	section.UpdateTime = now.Unix()
	section.DateTime = now
	return apiManage.UpdateSection(section)
}

func (a *ApiManage) GetPieceList(param req_model.GetPieceList) (resp_model.GetPieceList, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		total     int64
		err       error
		data      = make([]resp_model.PieceListData, 0)
		list      = resp_model.GetPieceList{}
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiManage.GetPieceList(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		list.List = append(list.List, v)
	}
	return list, nil
}

func (a *ApiManage) GetPieceSelectList(param req_model.GetPieceSelectList) ([]resp_model.PieceListData, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		err       error
		data      = make([]mysql_model.TPiece, 0)
		list      = make([]resp_model.PieceListData, 0)
	)
	data, err = apiManage.GetPieceSelectList(param)
	if err != nil {
		return list, err
	}
	for _, v := range data {
		list = append(list, resp_model.PieceListData{
			Id:            v.Id,
			Name:          v.Name,
			QuestionCount: v.QuestionCount,
			SectionId:     v.SectionId,
			Status:        v.Status,
			AddTime:       v.AddTime,
			UpdateTime:    v.UpdateTime,
		})
	}
	return list, nil
}

func (a *ApiManage) EditPiece(param req_model.EditPiece) error {
	var (
		piece     = mysql_model.TPiece{}
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)

	if len(param.Name) <= 0 {
		return custErr.New(custErr.PIECE_NAME_NOT_EMPTY, "")
	}
	if param.Id <= 0 {
		if param.SectionId <= 0 {
			return custErr.New(custErr.PIECE_SECTION_NOT_EMPTY, "")
		}
		piece.SectionId = param.SectionId
		piece.Name = param.Name
		piece.QuestionCount = param.QuestionCount
		piece.Status = 2
		piece.AddTime = now.Unix()
		piece.UpdateTime = now.Unix()
		piece.DateTime = now
		err = apiManage.AddPiece(piece)
	} else {
		piece, err = apiManage.GetPieceById(param.Id)
		if err != nil {
			return err
		}
		if piece.Id <= 0 {
			return custErr.New(custErr.SECTION_NOT_EXISTS, "")
		}
		piece.Name = param.Name
		piece.QuestionCount = param.QuestionCount
		piece.Status = param.Status
		piece.UpdateTime = now.Unix()
		piece.DateTime = now
		err = apiManage.UpdatePiece(piece)
	}
	return err
}

func (a *ApiManage) DelPiece(param req_model.DelPiece) error {
	var (
		question  = make([]mysql_model.TQuestion, 0)
		piece     = mysql_model.TPiece{}
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	piece, err = apiManage.GetPieceById(param.Id)
	if err != nil {
		return err
	}
	if piece.Id <= 0 {
		return custErr.New(custErr.PIECE_NOT_EXISTS, "")
	}
	question, err = apiManage.GetQuestionByPieceId(param.Id)
	if err != nil {
		return err
	}
	if len(question) > 0 {
		return custErr.New(custErr.PIECE_HAS_QUESTION, "")
	}
	piece.Status = 3
	piece.UpdateTime = now.Unix()
	piece.DateTime = now
	return apiManage.UpdatePiece(piece)
}

func (a *ApiManage) GetQuestionList(param req_model.GetQuestionList) (resp_model.GetQuestionList, error) {
	var (
		apiManage      = manage_dao.ApiManage{}
		total          int64
		err            error
		data           = make([]resp_model.GetQuestionsData, 0)
		child          = make([]mysql_model.TChildQuestion, 0)
		options        = make([]mysql_model.TOption, 0)
		childOptions   = make([]mysql_model.TOption, 0)
		childMap       = make(map[int64][]mysql_model.TChildQuestion)
		optionMap      = make(map[int64][]mysql_model.TOption)
		childOptionMap = make(map[int64][]mysql_model.TOption)
		list           = resp_model.GetQuestionList{}
		ids            = make([]int64, 0)
		cids           = make([]int64, 0)
	)
	if param.Size <= 0 {
		param.Size = 20
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	data, total, err = apiManage.GetQuestionList(param)
	if err != nil {
		return list, err
	}
	list.Count = total
	for _, v := range data {
		ids = append(ids, v.Id)
	}
	// 获取子题目信息
	child, err = apiManage.GetChildQuestionByPid(ids)
	if err != nil {
		return list, err
	}
	for _, v := range child {
		childMap[v.Pid] = append(childMap[v.Pid], v)
		cids = append(cids, v.Id)
	}
	// 获取选项信息
	options, err = apiManage.GetQuestionOption(ids)
	if err != nil {
		return list, err
	}
	for _, o := range options {
		optionMap[o.QuestionId] = append(optionMap[o.QuestionId], o)
	}
	// 获取子题目选项信息
	childOptions, err = apiManage.GetQuestionOption(cids)
	if err != nil {
		return list, err
	}
	for _, o := range childOptions {
		childOptionMap[o.QuestionId] = append(childOptionMap[o.QuestionId], o)
	}

	for _, v := range data {
		childData := childQuestion(v.Id, childMap, childOptionMap)
		optionData := questionOption(v.Id, optionMap)
		list.List = append(list.List, resp_model.QuestionListData{
			PackageId:         v.PackageId,
			PackageName:       v.PackageName,
			SubjectId:         v.SubjectId,
			SubjectName:       v.SubjectName,
			ChapterId:         v.ChapterId,
			ChapterName:       v.ChapterName,
			SectionId:         v.SectionId,
			SectionName:       v.SectionName,
			PieceId:           v.PieceId,
			PieceName:         v.PieceName,
			Id:                v.Id,
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
			Status:            v.Status,
			AddTime:           v.AddTime,
			UpdateTime:        v.UpdateTime,
			Child:             childData,
			Option:            optionData,
		})
	}
	return list, nil
}

func childQuestion(id int64, childMap map[int64][]mysql_model.TChildQuestion,
	childOptionMap map[int64][]mysql_model.TOption) []resp_model.ChildQuestionListData {
	childData := make([]resp_model.ChildQuestionListData, 0)
	if questions, ok := childMap[id]; ok {
		for _, que := range questions {
			optionData := questionOption(que.Id, childOptionMap)
			childData = append(childData, resp_model.ChildQuestionListData{
				Id:                que.Id,
				Sequence:          que.Sequence,
				QuestionType:      que.QuestionType,
				QuestionSource:    que.QuestionSource,
				QuestionContent:   que.QuestionContent,
				QuestionAnswer:    que.QuestionAnswer,
				Score:             que.Score,
				MainNodeId:        que.MainNodeId,
				MainNodeName:      que.MainNodeName,
				MainNodeFrequency: que.MainNodeFrequency,
				Favorite:          que.Favorite,
				CanPhoto:          que.CanPhoto,
				AvgCorrectRate:    que.AvgCorrectRate,
				AnalysisType:      que.AnalysisType,
				Analysis:          que.Analysis,
				Option:            optionData,
			})
		}
	}
	return childData
}

func questionOption(id int64, optionMap map[int64][]mysql_model.TOption) []resp_model.Option {
	optionData := make([]resp_model.Option, 0)
	if option, ok := optionMap[id]; ok {
		for _, op := range option {
			optionData = append(optionData, resp_model.Option{
				Id:         op.Id,
				QuestionId: op.QuestionId,
				Sequence:   op.Sequence,
				Title:      op.Title,
				Content:    op.Content,
				Correct:    op.Correct,
			})
		}
	}
	return optionData
}

func (a *ApiManage) EditQuestion(param req_model.EditQuestion) error {
	var (
		question  = mysql_model.TQuestion{}
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	if param.Sequence <= 0 {
		return custErr.New(custErr.SEQUENCE_NOT_EMPTY, "")
	}
	if len(param.QuestionType) <= 0 {
		return custErr.New(custErr.QUESTION_TYPE_NOT_EMPTY, "")
	}
	if len(param.QuestionContent) <= 0 {
		return custErr.New(custErr.QUESTION_CONTENT_NOT_EMPTY, "")
	}
	if param.Id <= 0 {
		var id int64
		id, err = apiManage.GetquestionMaxId()
		if err != nil {
			return err
		}
		if param.PieceId <= 0 {
			return custErr.New(custErr.PIECE_NOT_EXISTS, "")
		}
		question.Id = id
		question.PieceId = param.PieceId
		question.Sequence = param.Sequence
		question.QuestionType = param.QuestionType
		question.QuestionContent = param.QuestionContent
		question.QuestionAnswer = param.QuestionAnswer
		question.Score = param.Score
		question.Analysis = param.Analysis
		question.Status = param.Status
		question.AddTime = now.Unix()
		question.UpdateTime = now.Unix()
		question.DateTime = now
		err = apiManage.AddQuestion(question)
	} else {
		question.Id = param.Id
		question.PieceId = param.PieceId
		question.Sequence = param.Sequence
		question.QuestionType = param.QuestionType
		question.QuestionContent = param.QuestionContent
		question.QuestionAnswer = param.QuestionAnswer
		question.Score = param.Score
		question.Analysis = param.Analysis
		question.Status = param.Status
		question.UpdateTime = now.Unix()
		question.DateTime = now
		err = apiManage.UpdateQuestion(question)
	}
	return err
}

func (a *ApiManage) DelQuestion(param req_model.DelQuestion) error {
	var (
		question  = mysql_model.TQuestion{}
		apiManage = manage_dao.ApiManage{}
		err       error
		now       = time.Now()
	)
	question, err = apiManage.GetQuestionById(param.Id, param.PieceId)
	if err != nil {
		return err
	}
	if question.Id <= 0 {
		return custErr.New(custErr.QUESTION_EMPTY, "")
	}
	question.Status = 3
	question.UpdateTime = now.Unix()
	question.DateTime = now
	return apiManage.UpdateQuestion(question)
}

func (a *ApiManage) GetQuestionOptionList(param req_model.GetQuestionOptionList) ([]resp_model.GetQuestionOptionList, error) {
	var (
		apiManage = manage_dao.ApiManage{}
		err       error
		data      = make([]mysql_model.TOption, 0)
		list      = make([]resp_model.GetQuestionOptionList, 0)
	)
	data, err = apiManage.GetQuestionOptionList(param)
	if err != nil {
		return list, err
	}
	for _, v := range data {
		list = append(list, resp_model.GetQuestionOptionList{
			Id:         v.Id,
			QuestionId: v.QuestionId,
			Sequence:   v.Sequence,
			Title:      v.Title,
			Content:    v.Content,
			Correct:    v.Correct,
		})
	}
	return list, nil
}

func (a *ApiManage) EditQuestionOption(param req_model.EditQuestionOption) error {
	var (
		option    = mysql_model.TOption{}
		apiManage = manage_dao.ApiManage{}
		err       error
	)
	if len(param.Title) <= 0 {
		return custErr.New(custErr.QUESTION_OPTION_TITLE_NOT_EMPTY, "")
	}
	if len(param.Content) <= 0 {
		return custErr.New(custErr.QUESTION_OPTION_CONTENT_NOT_EMPTY, "")
	}
	if param.Id <= 0 {
		if param.QuestionId <= 0 {
			return custErr.New(custErr.QUESTION_EMPTY, "")
		}
		option.QuestionId = param.QuestionId
		option.Sequence = param.Sequence
		option.Title = param.Title
		option.Content = param.Content
		option.Correct = param.Correct
		err = apiManage.AddOption(option)
	} else {
		option, err = apiManage.GetOptionById(param.Id)
		if err != nil {
			return err
		}
		option.QuestionId = param.QuestionId
		option.Sequence = param.Sequence
		option.Title = param.Title
		option.Content = param.Content
		option.Correct = param.Correct
		err = apiManage.UpdateOption(option)
	}
	return err
}

func (a *ApiManage) DelQuestionOption(param req_model.EditQuestionOption) error {
	var (
		option    = mysql_model.TOption{}
		apiManage = manage_dao.ApiManage{}
		err       error
	)
	option, err = apiManage.GetOptionById(param.Id)
	if err != nil {
		return err
	}
	if option.Id <= 0 {
		return custErr.New(custErr.OPTION_NOT_EXISTS, "")
	}
	return apiManage.DelQuestionOption(option.Id)
}

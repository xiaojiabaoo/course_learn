package manage_dao

import (
	"course_learn/common/gormx"
	"course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

type Manage interface {
	GetManageList(param req_model.GetManageList) ([]mysql_model.Manager, int64, error)
	GetManageById(id int64) (mysql_model.Manager, error)
	GetManageByAccount(account string) (mysql_model.Manager, error)
	AddManage(manage mysql_model.Manager) error
	UpdateManage(manage mysql_model.Manager) error
	GetPackageList(param req_model.GetPackageList) ([]mysql_model.TPackage, int64, error)
	GetPackageSelectList(param req_model.GetPackageSelectList) ([]mysql_model.TPackage, error)
	AddPackage(packages mysql_model.TPackage) error
	UpdatePackage(packages mysql_model.TPackage) error
	GetPackageById(id int64) (mysql_model.TPackage, error)
	GetSubjectByPackageId(packageId int64) ([]mysql_model.TSubject, error)
	GetSubjectList(param req_model.GetSubjectsList) ([]resp_model.SubjectList, int64, error)
	GetSubjectSelectList(param req_model.GetSubjectSelectList) ([]mysql_model.TSubject, error)
	AddSubject(subject mysql_model.TSubject) error
	UpdateSubject(subject mysql_model.TSubject) error
	GetSubjectById(id int64) (mysql_model.TSubject, error)
	GetChapterBySubjectId(subjectId int64) ([]mysql_model.TChapter, error)
	GetChapterList(param req_model.GetChaptersList) ([]resp_model.ChapterList, int64, error)
	GetChapterSelectList(param req_model.GetChapterSelectList) ([]mysql_model.TChapter, error)
	GetChapterById(id int64) (mysql_model.TChapter, error)
	AddChapter(chapter mysql_model.TChapter) error
	UpdateChapter(chapter mysql_model.TChapter) error
	GetSectionByChapterId(chapterId int64) ([]mysql_model.TSection, error)
	GetSectionList(param req_model.GetSectionList) ([]resp_model.SectionList, int64, error)
	GetSectionSelectList(param req_model.GetSectionSelectList) ([]mysql_model.TSection, error)
	GetSectionById(id int64) (mysql_model.TSection, error)
	AddSection(section mysql_model.TSection) error
	UpdateSection(section mysql_model.TSection) error
	GetPieceBySectionId(sectionId int64) ([]mysql_model.TPiece, error)
	GetPieceList(param req_model.GetPieceList) ([]resp_model.PieceListData, int64, error)
	GetPieceSelectList(param req_model.GetPieceSelectList) ([]mysql_model.TPiece, error)
	GetPieceById(id int64) (mysql_model.TPiece, error)
	AddPiece(piece mysql_model.TPiece)
	UpdatePiece(piece mysql_model.TPiece)
	GetQuestionByPieceId(pieceId int64) ([]mysql_model.TQuestion, error)
	GetQuestionList(param req_model.GetQuestionList) ([]resp_model.GetQuestionsData, int64, error)
	GetChildQuestionByPid(pids []int64) ([]mysql_model.TChildQuestion, error)
	GetQuestionOption(ids []int64) ([]mysql_model.TOption, error)
	GetquestionMaxId()
	AddQuestion(question mysql_model.TQuestion)
	UpdateQuestion(question mysql_model.TQuestion)
	GetQuestionById(id, pieceId int64) (mysql_model.TQuestion, error)
	GetOptionById(id int64)
	GetQuestionOptionList(param req_model.GetQuestionOptionList)
	AddOption(question mysql_model.TOption)
	UpdateOption(option mysql_model.TOption)
	DelQuestionOption(id int64)
}

type ApiManage struct{}

func (a *ApiManage) GetManageList(param req_model.GetManageList) ([]mysql_model.Manager, int64, error) {
	var (
		total int64
		err   error
		data  = make([]mysql_model.Manager, 0)
	)
	db := gormx.DB.Debug().Table("manager")
	if len(param.Account) > 0 {
		db = db.Where("account like ?", "%"+param.Account+"%")
	}
	db = db.Where("status=1")
	db = db.Order("update_time desc")
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询管理员列表失败")
	}
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询管理员列表失败")
	}
	return data, total, nil
}

func (a *ApiManage) GetManageById(id int64) (mysql_model.Manager, error) {
	var (
		manage = mysql_model.Manager{}
		err    error
	)
	err = gormx.DB.Debug().Table("manager").Where("id=? and status=1", id).First(&manage).Error
	if err != nil && err.Error() != "record not found" {
		return manage, errors.Wrap(err, "查询管理员失败")
	}
	return manage, nil
}

func (a *ApiManage) GetManageByAccount(account string) (mysql_model.Manager, error) {
	var (
		manage = mysql_model.Manager{}
		err    error
	)
	err = gormx.DB.Debug().Table("manager").Where("account=? and status=1", account).First(&manage).Error
	if err != nil && err.Error() != "record not found" {
		return manage, errors.Wrap(err, "查询管理员失败")
	}
	return manage, nil
}

func (a *ApiManage) AddManage(manage mysql_model.Manager) error {
	err := gormx.DB.Debug().Table("manager").Create(&manage).Error
	if err != nil {
		return errors.Wrap(err, "添加管理员失败")
	}
	return nil
}

func (a *ApiManage) UpdateManage(manage mysql_model.Manager) error {
	err := gormx.DB.Debug().Table("manager").Where("id=?", manage.Id).Updates(&manage).Error
	if err != nil {
		return errors.Wrap(err, "修改管理员失败")
	}
	return nil
}

func (a *ApiManage) GetPackageList(param req_model.GetPackageList) ([]mysql_model.TPackage, int64, error) {
	var (
		total int64
		err   error
		data  = make([]mysql_model.TPackage, 0)
	)
	db := gormx.DB.Debug().Table("t_package")
	if len(param.PackageName) > 0 {
		db = db.Where("package_name like ?", "%"+param.PackageName+"%")
	}
	if param.Status > 0 {
		db = db.Where("status = ?", param.Status)
	} else {
		db = db.Where("status in (1,2)")
	}
	db = db.Order("update_time desc")
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询专业列表失败")
	}
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询专业列表失败")
	}
	return data, total, nil
}

func (a *ApiManage) GetPackageSelectList(param req_model.GetPackageSelectList) ([]mysql_model.TPackage, error) {
	var (
		err  error
		data = make([]mysql_model.TPackage, 0)
	)
	db := gormx.DB.Debug().Table("t_package")
	db = db.Where("status in (1,2)")
	db = db.Order("update_time desc")
	err = db.Find(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询专业列表失败")
	}
	return data, nil
}

func (a *ApiManage) AddPackage(packages mysql_model.TPackage) error {
	err := gormx.DB.Debug().Table("t_package").Create(&packages).Error
	if err != nil {
		return errors.Wrap(err, "添加专业失败")
	}
	return nil
}

func (a *ApiManage) UpdatePackage(packages mysql_model.TPackage) error {
	err := gormx.DB.Debug().Table("t_package").Where("package_id=?", packages.PackageId).Updates(&packages).Error
	if err != nil {
		return errors.Wrap(err, "修改专业失败")
	}
	return nil
}

func (a *ApiManage) GetPackageById(id int64) (mysql_model.TPackage, error) {
	var (
		packages = mysql_model.TPackage{}
		err      error
	)
	err = gormx.DB.Debug().Table("t_package").Where("package_id=? and status in (1,2)", id).First(&packages).Error
	if err != nil && err.Error() != "record not found" {
		return packages, errors.Wrap(err, "查询专业失败")
	}
	return packages, nil
}

func (a *ApiManage) GetSubjectByPackageId(packageId int64) ([]mysql_model.TSubject, error) {
	var (
		subject = make([]mysql_model.TSubject, 0)
		err     error
	)
	err = gormx.DB.Debug().Table("t_subject").Where("package_id=? and status in (1,2)", packageId).
		Find(&subject).Error
	if err != nil {
		return subject, errors.Wrap(err, "查询专业失败")
	}
	return subject, nil
}

func (a *ApiManage) GetSubjectList(param req_model.GetSubjectsList) ([]resp_model.SubjectList, int64, error) {
	var (
		total int64
		err   error
		data  = make([]resp_model.SubjectList, 0)
	)
	db := gormx.DB.Debug().Table("t_subject s").Select("s.*,p.package_name").
		Joins("left join t_package p on s.package_id = p.package_id")
	if param.PackageId > 0 {
		db = db.Where("s.package_id = ?", param.PackageId)
	}
	if len(param.SubjectName) > 0 {
		db = db.Where("s.subject_name like ?", "%"+param.SubjectName+"%")
	}
	if param.Status > 0 {
		db = db.Where("s.status = ?", param.Status)
	} else {
		db = db.Where("s.status in (1,2)")
	}
	db = db.Where("p.status in (1,2)")
	db = db.Order("s.update_time desc")
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询课程列表失败")
	}
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询课程列表失败")
	}
	return data, total, nil
}

func (a *ApiManage) GetSubjectSelectList(param req_model.GetSubjectSelectList) ([]mysql_model.TSubject, error) {
	var (
		err  error
		data = make([]mysql_model.TSubject, 0)
	)
	db := gormx.DB.Debug().Table("t_subject")
	if param.PackageId >= 0 {
		db = db.Where("package_id = ?", param.PackageId)
	}
	db = db.Where("status in (1,2)")
	db = db.Order("update_time desc")
	err = db.Find(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询课程列表失败")
	}
	return data, nil
}

func (a *ApiManage) AddSubject(subject mysql_model.TSubject) error {
	err := gormx.DB.Debug().Table("t_subject").Create(&subject).Error
	if err != nil {
		return errors.Wrap(err, "添加课程失败")
	}
	return nil
}

func (a *ApiManage) UpdateSubject(subject mysql_model.TSubject) error {
	err := gormx.DB.Debug().Table("t_subject").Where("subject_id=?", subject.SubjectId).Updates(&subject).Error
	if err != nil {
		return errors.Wrap(err, "修改课程失败")
	}
	return nil
}

func (a *ApiManage) GetSubjectById(id int64) (mysql_model.TSubject, error) {
	var (
		subject = mysql_model.TSubject{}
		err     error
	)
	err = gormx.DB.Debug().Table("t_subject").Where("subject_id=? and status in (1,2)", id).First(&subject).Error
	if err != nil && err.Error() != "record not found" {
		return subject, errors.Wrap(err, "查询课程失败")
	}
	return subject, nil
}

func (a *ApiManage) GetChapterBySubjectId(subjectId int64) ([]mysql_model.TChapter, error) {
	var (
		chapter = make([]mysql_model.TChapter, 0)
		err     error
	)
	err = gormx.DB.Debug().Table("t_chapter").Where("subject_id=? and status in (1,2)", subjectId).
		Find(&chapter).Error
	if err != nil {
		return chapter, errors.Wrap(err, "查询专业失败")
	}
	return chapter, nil
}

func (a *ApiManage) GetChapterList(param req_model.GetChaptersList) ([]resp_model.ChapterList, int64, error) {
	var (
		total int64
		err   error
		data  = make([]resp_model.ChapterList, 0)
	)
	db := gormx.DB.Debug().Table("t_chapter c").
		Select("c.*,p.package_name,p.package_id,s.subject_name").
		Joins("left join t_subject s on c.subject_id = s.subject_id").
		Joins("left join t_package p on s.package_id = p.package_id")
	if param.PackageId > 0 {
		db = db.Where("p.package_id = ?", param.PackageId)
	}
	if param.SubjectId > 0 {
		db = db.Where("c.subject_id = ?", param.SubjectId)
	}
	if len(param.ChapterName) > 0 {
		db = db.Where("c.name like ?", "%"+param.ChapterName+"%")
	}
	if param.Status > 0 {
		db = db.Where("c.status = ?", param.Status)
	} else {
		db = db.Where("c.status in (1,2)")
	}
	db = db.Where("p.status in (1,2)")
	db = db.Where("s.status in (1,2)")
	db = db.Order("c.update_time desc")
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询学习章列表失败")
	}
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询学习章列表失败")
	}
	return data, total, nil
}

func (a *ApiManage) GetChapterSelectList(param req_model.GetChapterSelectList) ([]mysql_model.TChapter, error) {
	var (
		err  error
		data = make([]mysql_model.TChapter, 0)
	)
	db := gormx.DB.Debug().Table("t_chapter")
	db = db.Where("subject_id = ?", param.SubjectId)
	db = db.Where("status in (1,2)")
	db = db.Order("update_time desc")
	err = db.Find(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询学习章列表失败")
	}
	return data, nil
}

func (a *ApiManage) GetChapterById(id int64) (mysql_model.TChapter, error) {
	var (
		chapter = mysql_model.TChapter{}
		err     error
	)
	err = gormx.DB.Debug().Table("t_chapter").Where("id=? and status in (1,2)", id).First(&chapter).Error
	if err != nil && err.Error() != "record not found" {
		return chapter, errors.Wrap(err, "查询学习章失败")
	}
	return chapter, nil
}

func (a *ApiManage) AddChapter(chapter mysql_model.TChapter) error {
	err := gormx.DB.Debug().Table("t_chapter").Create(&chapter).Error
	if err != nil {
		return errors.Wrap(err, "添加学习章失败")
	}
	return nil
}

func (a *ApiManage) UpdateChapter(chapter mysql_model.TChapter) error {
	err := gormx.DB.Debug().Table("t_chapter").Where("id=?", chapter.Id).Updates(&chapter).Error
	if err != nil {
		return errors.Wrap(err, "修改学习章失败")
	}
	return nil
}

func (a *ApiManage) GetSectionByChapterId(chapterId int64) ([]mysql_model.TSection, error) {
	var (
		section = make([]mysql_model.TSection, 0)
		err     error
	)
	err = gormx.DB.Debug().Table("t_section").Where("chapter_id=? and status in (1,2)", chapterId).
		Find(&section).Error
	if err != nil {
		return section, errors.Wrap(err, "查询专业失败")
	}
	return section, nil
}

func (a *ApiManage) GetSectionList(param req_model.GetSectionList) ([]resp_model.SectionList, int64, error) {
	var (
		total int64
		err   error
		data  = make([]resp_model.SectionList, 0)
	)
	db := gormx.DB.Debug().Table("t_section se").
		Select("se.*,p.package_name,p.package_id,s.subject_id,s.subject_name,c.name chapter_name").
		Joins("left join t_chapter c on c.id = se.chapter_id").
		Joins("left join t_subject s on c.subject_id = s.subject_id").
		Joins("left join t_package p on s.package_id = p.package_id")
	if param.PackageId > 0 {
		db = db.Where("p.package_id = ?", param.PackageId)
	}
	if param.SubjectId > 0 {
		db = db.Where("s.subject_id = ?", param.SubjectId)
	}
	if param.ChapterId > 0 {
		db = db.Where("se.chapter_id = ?", param.ChapterId)
	}
	if len(param.SectionName) > 0 {
		db = db.Where("se.name like ?", "%"+param.SectionName+"%")
	}
	if param.Status > 0 {
		db = db.Where("se.status = ?", param.Status)
	} else {
		db = db.Where("se.status in (1,2)")
	}
	db = db.Where("p.status in (1,2)")
	db = db.Where("s.status in (1,2)")
	db = db.Where("c.status in (1,2)")
	db = db.Order("se.update_time desc")
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询学习章列表失败")
	}
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询学习章列表失败")
	}
	return data, total, nil
}

func (a *ApiManage) GetSectionSelectList(param req_model.GetSectionSelectList) ([]mysql_model.TSection, error) {
	var (
		err  error
		data = make([]mysql_model.TSection, 0)
	)
	db := gormx.DB.Debug().Table("t_section se")
	db = db.Where("se.chapter_id = ?", param.ChapterId)
	db = db.Where("se.status in (1,2)")
	db = db.Order("se.update_time desc")
	err = db.Find(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询学习章列表失败")
	}
	return data, nil
}

func (a *ApiManage) GetSectionById(id int64) (mysql_model.TSection, error) {
	var (
		section = mysql_model.TSection{}
		err     error
	)
	err = gormx.DB.Debug().Table("t_section").Where("id=? and status in (1,2)", id).First(&section).Error
	if err != nil && err.Error() != "record not found" {
		return section, errors.Wrap(err, "查询学习节失败")
	}
	return section, nil
}

func (a *ApiManage) AddSection(section mysql_model.TSection) error {
	err := gormx.DB.Debug().Table("t_section").Create(&section).Error
	if err != nil {
		return errors.Wrap(err, "添加学习节失败")
	}
	return nil
}

func (a *ApiManage) UpdateSection(section mysql_model.TSection) error {
	err := gormx.DB.Debug().Table("t_section").Where("id=?", section.Id).Updates(&section).Error
	if err != nil {
		return errors.Wrap(err, "修改学习节失败")
	}
	return nil
}

func (a *ApiManage) GetPieceBySectionId(sectionId int64) ([]mysql_model.TPiece, error) {
	var (
		piece = make([]mysql_model.TPiece, 0)
		err   error
	)
	err = gormx.DB.Debug().Table("t_piece").Where("section_id=? and status in (1,2)", sectionId).
		Find(&piece).Error
	if err != nil {
		return piece, errors.Wrap(err, "查询学习块失败")
	}
	return piece, nil
}

func (a *ApiManage) GetPieceList(param req_model.GetPieceList) ([]resp_model.PieceListData, int64, error) {
	var (
		total int64
		err   error
		data  = make([]resp_model.PieceListData, 0)
	)
	db := gormx.DB.Debug().Table("t_piece pi").
		Select("pi.*,p.package_name,p.package_id,s.subject_id,s.subject_name,c.id chapter_id,c.name chapter_name," +
			"se.id section_id,se.name section_name").
		Joins("left join t_section se on se.id = pi.section_id").
		Joins("left join t_chapter c on c.id = se.chapter_id").
		Joins("left join t_subject s on c.subject_id = s.subject_id").
		Joins("left join t_package p on s.package_id = p.package_id")
	if param.PackageId > 0 {
		db = db.Where("p.package_id = ?", param.PackageId)
	}
	if param.SubjectId > 0 {
		db = db.Where("s.subject_id = ?", param.SubjectId)
	}
	if param.ChapterId > 0 {
		db = db.Where("c.id = ?", param.ChapterId)
	}
	if param.SectionId > 0 {
		db = db.Where("pi.section_id = ?", param.SectionId)
	}
	if len(param.PieceName) > 0 {
		db = db.Where("pi.name like ?", "%"+param.PieceName+"%")
	}
	if param.Status > 0 {
		db = db.Where("pi.status = ?", param.Status)
	} else {
		db = db.Where("pi.status in (1,2)")
	}
	db = db.Where("p.status in (1,2)")
	db = db.Where("s.status in (1,2)")
	db = db.Where("c.status in (1,2)")
	db = db.Where("se.status in (1,2)")
	db = db.Order("pi.update_time desc")
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询学习块列表失败")
	}
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询学习块列表失败")
	}
	return data, total, nil
}

func (a *ApiManage) GetPieceSelectList(param req_model.GetPieceSelectList) ([]mysql_model.TPiece, error) {
	var (
		err  error
		data = make([]mysql_model.TPiece, 0)
	)
	db := gormx.DB.Debug().Table("t_piece pi")
	db = db.Where("pi.section_id = ?", param.SectionId)
	db = db.Where("pi.status in (1,2)")
	db = db.Order("pi.update_time desc")
	err = db.Find(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询学习块列表失败")
	}
	return data, nil
}

func (a *ApiManage) GetPieceById(id int64) (mysql_model.TPiece, error) {
	var (
		piece = mysql_model.TPiece{}
		err   error
	)
	err = gormx.DB.Debug().Table("t_piece").Where("id=? and status in (1,2)", id).First(&piece).Error
	if err != nil && err.Error() != "record not found" {
		return piece, errors.Wrap(err, "查询学习节失败")
	}
	return piece, nil
}

func (a *ApiManage) AddPiece(piece mysql_model.TPiece) error {
	err := gormx.DB.Debug().Table("t_piece").Create(&piece).Error
	if err != nil {
		return errors.Wrap(err, "添加学习块失败")
	}
	return nil
}

func (a *ApiManage) UpdatePiece(piece mysql_model.TPiece) error {
	err := gormx.DB.Debug().Table("t_piece").Where("id=?", piece.Id).Updates(&piece).Error
	if err != nil {
		return errors.Wrap(err, "添加学习块失败")
	}
	return nil
}

func (a *ApiManage) GetQuestionByPieceId(pieceId int64) ([]mysql_model.TQuestion, error) {
	var (
		question = make([]mysql_model.TQuestion, 0)
		err      error
	)
	err = gormx.DB.Debug().Table("t_question").Where("piece_id=? and status in (1,2)", pieceId).
		Find(&question).Error
	if err != nil {
		return question, errors.Wrap(err, "查询题目失败")
	}
	return question, nil
}

func (a *ApiManage) GetQuestionList(param req_model.GetQuestionList) ([]resp_model.GetQuestionsData, int64, error) {
	var (
		total int64
		err   error
		data  = make([]resp_model.GetQuestionsData, 0)
	)
	db := gormx.DB.Debug().Table("t_question q").
		Select("q.*,p.package_name,p.package_id,s.subject_id,s.subject_name,c.id chapter_id,c.name chapter_name," +
			"se.id section_id,se.name section_name,pi.id piece_id,pi.name piece_name").
		Joins("left join t_piece pi on pi.id = q.piece_id").
		Joins("left join t_section se on se.id = pi.section_id").
		Joins("left join t_chapter c on c.id = se.chapter_id").
		Joins("left join t_subject s on c.subject_id = s.subject_id").
		Joins("left join t_package p on s.package_id = p.package_id")
	if param.PackageId > 0 {
		db = db.Where("p.package_id = ?", param.PackageId)
	}
	if param.SubjectId > 0 {
		db = db.Where("s.subject_id = ?", param.SubjectId)
	}
	if param.ChapterId > 0 {
		db = db.Where("c.id = ?", param.ChapterId)
	}
	if param.SectionId > 0 {
		db = db.Where("se.id = ?", param.SectionId)
	}
	if param.PieceId > 0 {
		db = db.Where("q.piece_id = ?", param.PieceId)
	}
	if len(param.QuestionContent) > 0 {
		db = db.Where("q.question_content like ?", "%"+param.QuestionContent+"%")
	}
	if len(param.QuestionType) > 0 {
		questionType := make([]string, 0)
		for _, v := range strings.Split(param.QuestionType, ",") {
			questionType = append(questionType, v)
		}
		db = db.Where("q.question_type in ?", questionType)
	}
	if param.Status > 0 {
		db = db.Where("q.status = ?", param.Status)
	} else {
		db = db.Where("q.status in (1,2)")
	}
	db = db.Where("p.status in (1,2)")
	db = db.Where("s.status in (1,2)")
	db = db.Where("c.status in (1,2)")
	db = db.Where("se.status in (1,2)")
	db = db.Where("q.status in (1,2)")
	db = db.Order("q.update_time desc")
	db.Session(&gorm.Session{})
	err = db.Count(&total).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询题目列表失败")
	}
	err = db.Limit(param.Size).Offset((param.Page - 1) * param.Size).Find(&data).Error
	if err != nil {
		return data, total, errors.Wrap(err, "查询题目列表失败")
	}
	return data, total, nil
}

func (a *ApiManage) GetChildQuestionByPid(pids []int64) ([]mysql_model.TChildQuestion, error) {
	var (
		err  error
		data = make([]mysql_model.TChildQuestion, 0)
	)
	err = gormx.DB.Debug().Table("t_child_question").Where("pid in ?", pids).Find(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询题目列表失败")
	}
	return data, nil
}

func (a *ApiManage) GetQuestionOption(ids []int64) ([]mysql_model.TOption, error) {
	var (
		err  error
		data = make([]mysql_model.TOption, 0)
	)
	err = gormx.DB.Debug().Table("t_option").Where("question_id in ?", ids).Order("sequence asc").
		Find(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询题目选项失败")
	}
	return data, nil
}

func (a *ApiManage) GetquestionMaxId() (int64, error) {
	var (
		err   error
		maxId int64
	)
	err = gormx.DB.Debug().Select("max(id)+1 id").Table("t_question").Scan(&maxId).Error
	if err != nil {
		return maxId, errors.Wrap(err, "查询题目失败")
	}
	return maxId, nil
}

func (a *ApiManage) AddQuestion(question mysql_model.TQuestion) error {
	err := gormx.DB.Debug().Table("t_question").Create(&question).Error
	if err != nil {
		return errors.Wrap(err, "添加题目失败")
	}
	return nil
}

func (a *ApiManage) UpdateQuestion(question mysql_model.TQuestion) error {
	err := gormx.DB.Debug().Table("t_question").Where("id=? and piece_id=?", question.Id, question.PieceId).Updates(&question).Error
	if err != nil {
		return errors.Wrap(err, "修改题目失败")
	}
	return nil
}

func (a *ApiManage) GetQuestionById(id, pieceId int64) (mysql_model.TQuestion, error) {
	var (
		err  error
		data = mysql_model.TQuestion{}
	)
	err = gormx.DB.Debug().Table("t_question").Where("id = ? and piece_id = ?", id, pieceId).First(&data).Error
	if err != nil && err.Error() != "record not found" {
		return data, errors.Wrap(err, "查询题目失败")
	}
	return data, nil
}

func (a *ApiManage) GetOptionById(id int64) (mysql_model.TOption, error) {
	var (
		err  error
		data = mysql_model.TOption{}
	)
	err = gormx.DB.Debug().Table("t_option").Where("id = ?", id).First(&data).Error
	if err != nil && err.Error() != "record not found" {
		return data, errors.Wrap(err, "查询题目失败")
	}
	return data, nil
}

func (a *ApiManage) GetQuestionOptionList(param req_model.GetQuestionOptionList) ([]mysql_model.TOption, error) {
	var (
		err  error
		data = make([]mysql_model.TOption, 0)
	)
	db := gormx.DB.Debug().Table("t_option")
	db = db.Where("question_id = ?", param.QuestionId)
	err = db.Find(&data).Error
	if err != nil {
		return data, errors.Wrap(err, "查询选项列表失败")
	}
	return data, nil
}

func (a *ApiManage) AddOption(question mysql_model.TOption) error {
	err := gormx.DB.Debug().Table("t_option").Create(&question).Error
	if err != nil {
		return errors.Wrap(err, "添加题目失败")
	}
	return nil
}

func (a *ApiManage) UpdateOption(option mysql_model.TOption) error {
	err := gormx.DB.Debug().Table("t_option").Where("id=?", option.Id).Updates(&option).Error
	if err != nil {
		return errors.Wrap(err, "修改选项失败")
	}
	return nil
}

func (a *ApiManage) DelQuestionOption(id int64) error {
	option := mysql_model.TOption{}
	err := gormx.DB.Debug().Table("t_option").Where("id=?", id).Delete(&option).Error
	if err != nil {
		return errors.Wrap(err, "删除选项失败")
	}
	return nil
}

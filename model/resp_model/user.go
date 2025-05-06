package resp_model

import "course_learn/model/mysql_model"

type GetUserLeanRecord struct{}

type GetUserData struct {
	NickName      string `json:"nick_name"`
	Gender        int64  `json:"gender"`
	Logo          string `json:"logo"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Area          string `json:"area"`
	AddTime       int64  `json:"add_time"`
	AccountStatus int64  `json:"account_status"`
}

type GetLoginRecordData struct {
	Count int64             `json:"count"`
	List  []LoginRecordData `json:"list"`
}

type LoginRecordData struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	LoginTime   int64  `json:"login_time"`
	LoginIp     string `json:"login_ip"`
	LoginArea   string `json:"login_area"`
	LoginMethod int64  `json:"login_method"`
	NickName    string `json:"nick_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

type GetOrderListData struct {
	Count int64           `json:"count"`
	List  []OrderListData `json:"list"`
}

type OrderListData struct {
	Id             int64   `json:"id"`
	OrderNo        string  `json:"order_no"`
	Type           int64   `json:"type"`
	IsPayment      int64   `json:"is_payment"`
	Amount         float64 `json:"amount"`
	BuyType        int64   `json:"buy_type"`
	IsValid        int64   `json:"is_valid"`
	ExpirationTime int64   `json:"expiration_time"`
	UserId         int64   `json:"user_id"`
	SubjectId      int64   `json:"subject_id"`
	SubjectName    string  `json:"subject_name"`
	AddTime        int64   `json:"add_time"`
	NickName       string  `json:"nick_name"`
	Phone          string  `json:"phone"`
	Email          string  `json:"email"`
}

type GetLearnRecordData struct {
	Count int64             `json:"count"`
	List  []LearnRecordData `json:"list"`
}

type LearnRecordData struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	NickName    string `json:"nick_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	PackageId   int64  `json:"package_id"`
	PackageName string `json:"package_name"`
	SubjectId   int64  `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	ChapterId   int64  `json:"chapter_id"`
	ChapterName string `json:"chapter_name"`
	SectionId   int64  `json:"section_id"`
	SectionName string `json:"section_name"`
	PieceId     int64  `json:"piece_id"`
	PieceName   string `json:"piece_name"`
	AddTime     int64  `json:"add_time"`
}

type GetUserListData struct {
	Count int64          `json:"count"`
	List  []UserListData `json:"list"`
}

type UserListData struct {
	Id            int64  `json:"id"`
	NickName      string `json:"nick_name"`
	Gender        int64  `json:"gender"`
	Logo          string `json:"logo"`
	AreaCode      string `json:"area_code"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Area          string `json:"area"`
	AddTime       int64  `json:"add_time"`
	UpdateTime    int64  `json:"update_time"`
	AccountStatus int64  `json:"account_status"`
}

type GetManageList struct {
	Count int64            `json:"count"`
	List  []ManageListData `json:"list"`
}

type ManageListData struct {
	Id         int64  `json:"id"`
	Account    string `json:"account"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Identifier string `json:"identifier"`
	Status     int64  `json:"status"`
	AddTime    int64  `json:"add_time"`
	UpdateTime int64  `json:"update_time"`
	DateTime   string `json:"date_time"`
}

type GetPackageList struct {
	Count int64             `json:"count"`
	List  []PackageListData `json:"list"`
}

type PackageListData struct {
	PackageId   int64  `json:"package_id"`
	PackageName string `json:"package_name"`
	Status      int64  `json:"status"`
	AddTime     int64  `json:"add_time"`
	UpdateTime  int64  `json:"update_time"`
	DateTime    string `json:"date_time"`
}

type GetSubjectsList struct {
	Count int64              `json:"count"`
	List  []SubjectsListData `json:"list"`
}

type SubjectsListData struct {
	KnowledgeTreeId int64  `json:"knowledge_tree_id"`
	PackageId       int64  `json:"package_id"`
	PackageName     string `json:"package_name"`
	SubjectId       int64  `json:"subject_id"`
	SubjectName     string `json:"subject_name"`
	Sort            int64  `json:"sort"`
	Desc            string `json:"desc"`
	IsHot           int64  `json:"is_hot"`
	Status          int64  `json:"status"`
	AddTime         int64  `json:"add_time"`
	UpdateTime      int64  `json:"update_time"`
}

type GetChaptersList struct {
	Count int64              `json:"count"`
	List  []ChaptersListData `json:"list"`
}

type ChaptersListData struct {
	PackageId     int64  `json:"package_id"`
	PackageName   string `json:"package_name"`
	SubjectId     int64  `json:"subject_id"`
	SubjectName   string `json:"subject_name"`
	Id            int64  `json:"id"`
	Sequence      string `json:"sequence"`
	Name          string `json:"name"`
	QuestionCount int64  `json:"question_count"`
	Sort          int64  `json:"sort"`
	Status        int64  `json:"status"`
	AddTime       int64  `json:"add_time"`
	UpdateTime    int64  `json:"update_time"`
}

type SubjectList struct {
	mysql_model.TSubject
	PackageName string `json:"package_name"`
}

type ChapterList struct {
	mysql_model.TChapter
	SubjectName string `json:"subject_name"`
	PackageId   int64  `json:"package_id"`
	PackageName string `json:"package_name"`
}

type GetSectionList struct {
	Count int64         `json:"count"`
	List  []SectionList `json:"list"`
}

type SectionList struct {
	PackageId     int64  `json:"package_id"`
	PackageName   string `json:"package_name"`
	SubjectId     int64  `json:"subject_id"`
	SubjectName   string `json:"subject_name"`
	ChapterId     int64  `json:"chapter_id"`
	ChapterName   string `json:"chapter_name"`
	Id            int64  `json:"id"`
	Sequence      string `json:"sequence"`
	Name          string `json:"name"`
	QuestionCount int64  `json:"question_count"`
	Status        int64  `json:"status"`
	AddTime       int64  `json:"add_time"`
	UpdateTime    int64  `json:"update_time"`
}

type GetPieceList struct {
	Count int64           `json:"count"`
	List  []PieceListData `json:"list"`
}

type PieceListData struct {
	PackageId     int64  `json:"package_id"`
	PackageName   string `json:"package_name"`
	SubjectId     int64  `json:"subject_id"`
	SubjectName   string `json:"subject_name"`
	ChapterId     int64  `json:"chapter_id"`
	ChapterName   string `json:"chapter_name"`
	SectionId     int64  `json:"section_id"`
	SectionName   string `json:"section_name"`
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	QuestionCount int64  `json:"question_count"`
	Status        int64  `json:"status"`
	AddTime       int64  `json:"add_time"`
	UpdateTime    int64  `json:"update_time"`
}

type GetQuestionList struct {
	Count int64              `json:"count"`
	List  []QuestionListData `json:"list"`
}

type QuestionListData struct {
	PackageId         int64                   `json:"package_id"`
	PackageName       string                  `json:"package_name"`
	SubjectId         int64                   `json:"subject_id"`
	SubjectName       string                  `json:"subject_name"`
	ChapterId         int64                   `json:"chapter_id"`
	ChapterName       string                  `json:"chapter_name"`
	SectionId         int64                   `json:"section_id"`
	SectionName       string                  `json:"section_name"`
	PieceId           int64                   `json:"piece_id"`
	PieceName         string                  `json:"piece_name"`
	Id                int64                   `json:"id"`
	Sequence          int64                   `json:"sequence"`
	QuestionType      string                  `json:"question_type"`
	QuestionSource    string                  `json:"question_source"`
	QuestionContent   string                  `json:"question_content"`
	QuestionAnswer    string                  `json:"question_answer"`
	Score             string                  `json:"score"`
	MainNodeId        int64                   `json:"main_node_id"`
	MainNodeName      string                  `json:"main_node_name"`
	MainNodeFrequency string                  `json:"main_node_frequency"`
	Favorite          int64                   `json:"favorite"`
	CanPhoto          int64                   `json:"can_photo"`
	AvgCorrectRate    string                  `json:"avg_correct_rate"`
	AnalysisType      int64                   `json:"analysis_type"`
	Analysis          string                  `json:"analysis"`
	Status            int64                   `json:"status"`
	AddTime           int64                   `json:"add_time"`
	UpdateTime        int64                   `json:"update_time"`
	Child             []ChildQuestionListData `json:"child"`
	Option            []Option                `json:"option"`
}

type ChildQuestionListData struct {
	Id                int64    `json:"id"`
	Sequence          int64    `json:"sequence"`
	QuestionType      string   `json:"question_type"`
	QuestionSource    string   `json:"question_source"`
	QuestionContent   string   `json:"question_content"`
	QuestionAnswer    string   `json:"question_answer"`
	Score             string   `json:"score"`
	MainNodeId        int64    `json:"main_node_id"`
	MainNodeName      string   `json:"main_node_name"`
	MainNodeFrequency string   `json:"main_node_frequency"`
	Favorite          int64    `json:"favorite"`
	CanPhoto          int64    `json:"can_photo"`
	AvgCorrectRate    string   `json:"avg_correct_rate"`
	AnalysisType      int64    `json:"analysis_type"`
	Analysis          string   `json:"analysis"`
	Option            []Option `json:"option"`
}

type GetQuestionsData struct {
	PackageId         int64  `json:"package_id"`
	PackageName       string `json:"package_name"`
	SubjectId         int64  `json:"subject_id"`
	SubjectName       string `json:"subject_name"`
	ChapterId         int64  `json:"chapter_id"`
	ChapterName       string `json:"chapter_name"`
	SectionId         int64  `json:"section_id"`
	SectionName       string `json:"section_name"`
	PieceId           int64  `json:"piece_id"`
	PieceName         string `json:"piece_name"`
	Id                int64  `json:"id"`
	Sequence          int64  `json:"sequence"`
	QuestionType      string `json:"question_type"`
	QuestionSource    string `json:"question_source"`
	QuestionContent   string `json:"question_content"`
	QuestionAnswer    string `json:"question_answer"`
	Score             string `json:"score"`
	MainNodeId        int64  `json:"main_node_id"`
	MainNodeName      string `json:"main_node_name"`
	MainNodeFrequency string `json:"main_node_frequency"`
	Favorite          int64  `json:"favorite"`
	CanPhoto          int64  `json:"can_photo"`
	AvgCorrectRate    string `json:"avg_correct_rate"`
	AnalysisType      int64  `json:"analysis_type"`
	Analysis          string `json:"analysis"`
	Status            int64  `json:"status"`
	AddTime           int64  `json:"add_time"`
	UpdateTime        int64  `json:"update_time"`
}

type GetQuestionOptionList struct {
	Id         int64  `json:"id"`
	QuestionId int64  `json:"question_id"`
	Sequence   int64  `json:"sequence"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Correct    int64  `json:"correct"`
}

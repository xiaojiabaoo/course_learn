package req_model

type GetUserLeanRecord struct {
	UserId int64 `json:"user_id"` //后台使用，客户端无需理会该字段
}

type AddUserLeanRecord struct {
	UserId  int64 `json:"user_id"`  //后台使用，客户端无需理会该字段
	PieceId int64 `json:"piece_id"` //学习节块ID
}

type GetUserData struct {
	UserId int64 `json:"user_id"` //后台使用，客户端无需理会该字段
}

type SetUserData struct {
	UserId   int64  `json:"user_id"` //后台使用，客户端无需理会该字段
	NickName string `json:"nick_name"`
	Gender   int64  `json:"gender"`
	Logo     string `json:"logo"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Area     string `json:"area"`
}

type ChangePassword struct {
	UserId          int64  `json:"user_id"` //后台使用，客户端无需理会该字段
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
	SecurityCode    int64  `json:"security_code"` //验证码
}

type GetLoginRecord struct {
	UserId   int64  `form:"user_id"` //后台使用，客户端无需理会该字段
	Page     int    `form:"page"`
	Size     int    `form:"size"`
	Action   string `form:"action"`    // MANAGE_LOGIN.管理员查询；USER_LOGIN.用户查询
	NickName string `form:"nick_name"` // 后台使用
	Phone    string `form:"phone"`     // 后台使用
	Email    string `form:"email"`     // 后台使用
}

type GetOrderList struct {
	UserId    int64  `form:"user_id"` //后台使用，客户端无需理会该字段
	Page      int    `form:"page"`
	Size      int    `form:"size"`
	SubjectId int64  `form:"subject_id"`
	Action    string `form:"action"`    // MANAGE_LOGIN.管理员查询；USER_LOGIN.用户查询
	NickName  string `form:"nick_name"` // 后台使用
	Phone     string `form:"phone"`     // 后台使用
	Email     string `form:"email"`     // 后台使用
}

type GetLearnRecord struct {
	UserId    int64   `form:"user_id"`   // 后台使用，客户端无需理会该字段
	UserIds   []int64 `form:"user_ids"`  // 后台使用，客户端无需理会该字段
	NickName  string  `form:"nick_name"` // 后台使用
	Phone     string  `form:"phone"`     // 后台使用
	Email     string  `form:"email"`     // 后台使用
	Page      int     `form:"page"`
	Size      int     `form:"size"`
	PackageId int64   `form:"package_id"`
	SubjectId int64   `form:"subject_id"`
	ChapterId int64   `form:"chapter_id"`
	SectionId int64   `form:"section_id"`
	PieceId   int64   `form:"piece_id"`
	Action    string  `form:"action"` // MANAGE_LOGIN.管理员查询；USER_LOGIN.用户查询
}

type GetUserList struct {
	ManageId int64  `form:"manage_id"` //后台使用，客户端无需理会该字段
	UserName string `form:"user_name"`
	Email    string `form:"email"`
	Phone    string `form:"phone"`
	Gender   string `form:"gender"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}

type EditUserData struct {
	UserId        int64 `json:"user_id"`
	AccountStatus int64 `json:"account_status"`
}

type DelUserData struct {
	UserId int64 `json:"user_id"`
}

type GetManageList struct {
	ManageId int64  `form:"manage_id"`
	Account  string `form:"account"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}

type EditManage struct {
	Id         int64  `json:"id"`
	Account    string `json:"account"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Identifier string `json:"identifier"`
	Status     int64  `json:"status"`
}

type GetPackageList struct {
	ManageId    int64  `form:"manage_id"`
	PackageName string `form:"package_name"`
	Status      int    `form:"status"`
	Page        int    `form:"page"`
	Size        int    `form:"size"`
}

type GetPackageSelectList struct {
	ManageId int64 `json:"manage_id"`
}

type EditPackage struct {
	PackageId   int64  `json:"package_id"`
	PackageName string `json:"package_name"`
	Status      int64  `form:"status"` // 新增时无效
}

type DelPackage struct {
	PackageId int64 `json:"package_id"`
}

type GetSubjectsList struct {
	PackageId   int64  `form:"package_id"`
	SubjectName string `form:"subject_name"`
	IsHot       int64  `form:"is_hot"`
	Status      int64  `form:"status"`
	Page        int    `form:"page"`
	Size        int    `form:"size"`
}

type GetSubjectSelectList struct {
	PackageId int64 `json:"package_id"`
}

type EditSubject struct {
	PackageId   int64  `json:"package_id"` // 修改时无效
	SubjectId   int64  `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	Sort        int64  `json:"sort"`
	Status      int64  `json:"status"` // 新增时无效
	Desc        string `json:"desc"`
	IsHot       int64  `json:"is_hot"`
}

type DelSubject struct {
	SubjectId int64 `json:"subject_id"`
}

type GetChaptersList struct {
	PackageId   int64  `form:"package_id"`
	SubjectId   int64  `form:"subject_id"`
	ChapterName string `form:"chapter_name"`
	Status      int64  `form:"status"`
	Page        int    `form:"page"`
	Size        int    `form:"size"`
}

type GetChapterSelectList struct {
	SubjectId int64 `json:"subject_id"`
}

type EditChapter struct {
	SubjectId     int64  `json:"subject_id"` // 修改时无效
	Id            int64  `json:"id"`
	Sequence      string `json:"sequence"`
	Name          string `json:"name"`
	QuestionCount int64  `json:"question_count"`
	Sort          int64  `json:"sort"`
	Status        int64  `json:"status"` // 新增时无效
}

type DelChapter struct {
	Id int64 `json:"id"`
}

type GetSectionList struct {
	PackageId   int64  `form:"package_id"`
	SubjectId   int64  `form:"subject_id"`
	ChapterId   int64  `form:"chapter_id"`
	SectionName string `form:"section_name"`
	Status      int64  `form:"status"`
	Page        int    `form:"page"`
	Size        int    `form:"size"`
}

type GetSectionSelectList struct {
	ChapterId int64 `json:"chapter_id"`
}

type EditSection struct {
	ChapterId     int64  `json:"chapter_id"` // 修改时无效
	Id            int64  `json:"id"`
	Sequence      string `json:"sequence"`
	Name          string `json:"name"`
	QuestionCount int64  `json:"question_count"`
	Status        int64  `json:"status"` // 新增时无效
}

type DelSection struct {
	Id int64 `json:"id"`
}

type GetPieceList struct {
	PackageId int64  `form:"package_id"`
	SubjectId int64  `form:"subject_id"`
	ChapterId int64  `form:"chapter_id"`
	SectionId int64  `form:"section_id"`
	PieceName string `form:"name"`
	Status    int64  `form:"status"`
	Page      int    `form:"page"`
	Size      int    `form:"size"`
}

type GetPieceSelectList struct {
	SectionId int64 `json:"section_id"`
}

type EditPiece struct {
	SectionId     int64  `json:"section_id"` // 修改时无效
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	QuestionCount int64  `json:"question_count"`
	Status        int64  `json:"status"` // 新增时无效
}

type DelPiece struct {
	Id int64 `json:"id"`
}

type GetQuestionList struct {
	PackageId       int64  `form:"package_id"`
	SubjectId       int64  `form:"subject_id"`
	ChapterId       int64  `form:"chapter_id"`
	SectionId       int64  `form:"section_id"`
	PieceId         int64  `form:"piece_id"`
	Status          int64  `form:"status"`
	QuestionContent string `form:"question_content"`
	QuestionType    string `form:"question_type"`
	Page            int    `form:"page"`
	Size            int    `form:"size"`
}

type EditQuestion struct {
	Id              int64  `json:"id"`
	PieceId         int64  `json:"piece_id"`
	Sequence        int64  `json:"sequence"`
	QuestionType    string `json:"question_type"`
	QuestionContent string `json:"question_content"`
	QuestionAnswer  string `json:"question_answer"`
	Score           string `json:"score"`
	Analysis        string `json:"analysis"`
	Status          int64  `json:"status"`
}

type DelQuestion struct {
	Id      int64 `json:"id"`
	PieceId int64 `json:"piece_id"`
}

type GetQuestionOptionList struct {
	QuestionId int64 `form:"question_id"`
}

type EditQuestionOption struct {
	Id         int64  `json:"id"`
	QuestionId int64  `json:"question_id"`
	Sequence   int64  `json:"sequence"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Correct    int64  `json:"correct"`
}

type DelQuestionOption struct {
	Id int64 `json:"id"`
}

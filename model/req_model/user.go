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
	UserId int64 `form:"user_id"` //后台使用，客户端无需理会该字段
	Page   int   `form:"page"`
	Size   int   `form:"size"`
}

type GetOrderList struct {
	UserId    int64 `form:"user_id"` //后台使用，客户端无需理会该字段
	Page      int   `form:"page"`
	Size      int   `form:"size"`
	SubjectId int64 `form:"subject_id"`
}

type GetLearnRecord struct {
	UserId    int64 `form:"user_id"` //后台使用，客户端无需理会该字段
	Page      int   `form:"page"`
	Size      int   `form:"size"`
	PackageId int64 `form:"package_id"`
	SubjectId int64 `form:"subject_id"`
	ChapterId int64 `form:"chapter_id"`
	SectionId int64 `form:"section_id"`
	PieceId   int64 `form:"piece_id"`
}

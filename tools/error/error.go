package error

type CustomError struct {
	err  string
	Code int64
}

func (m *CustomError) Error() string {
	return m.err
}

func New(code int64, msg string) error {
	if msg == "" {
		msg = GetErrMsg(code)
	}
	return &CustomError{
		err:  msg,
		Code: code,
	}
}

const (
	//系统相关错误码
	SUCCES       = 0
	SERVER_ERROR = 99999
	//参数验证相关错误码
	INVALID_PARAMS       = 20000
	PARAM_ERROR          = 20001
	TOKEN_ERROR          = 20002
	SUBJECT_EMPTY_ERROR  = 20003
	PIECE_EMPTY_ERROR    = 20004
	IMAGE_EMPTY_ERROR    = 20005
	SECURITY_EMPTY_ERROR = 20006
	//账号问题相关错误码
	ACCOUNT_EXISTS           = 30000
	ACCOUNT_NOT_EXISTS       = 30001
	ACCOUNT_EXCEPTION        = 30002
	ACCOUNT_FORMAT_ERROR     = 30003
	ACCOUNT_PASS_ERROR       = 30004
	ACCOUNT_BLOCKED_ERROR    = 30005
	ACCOUNT_NOT_EMPTY        = 30006
	ACCOUNT_LOGIN_TYPE_ERROR = 30007
	ACCOUNT_LOGIN_ERROR      = 30008
	MANAGER_IDENTIFIER_ERROR = 30009
	//验证码问题相关错误码
	SECURITY_CODE_EXISTS = 40000
	SMS_SEND_ERROR       = 40001
	SECURITY_CODE_ERROR  = 40002
	SEND_TYPE_ERROR      = 40003
	//用户相关错误码
	USER_TRY_MAX_ERROR          = 50000
	USER_TRY_SAME_ERROR         = 50001
	USER_ID_EMPTY_ERROR         = 50002
	USER_PIECEID_EMPTY_ERROR    = 50003
	USER_NOT_TRY_ORBUY_ERROR    = 50004
	USER_CONFIRM_PASSWORD_ERROR = 50005
	USER_PASSWORD_ERROR         = 50006
	//管理员相关错误码
	MANAGER_NOT_EXISTS = 60000
	MANAGER_EXISTS     = 60001
	//题库相关错误码
	PACKAGE_NOT_EXISTS         = 70000
	PACKAGE_HAS_SUBJECT        = 70001
	SUBJECT_NOT_EXISTS         = 70002
	SUBJECT_HAS_CHAPTER        = 70003
	CHAPTER_NOT_EXISTS         = 70004
	CHAPTER_HAS_SECTION        = 70005
	SECTION_NOT_EXISTS         = 70006
	SECTION_HAS_PIECE          = 70007
	PIECE_NOT_EXISTS           = 70008
	PIECE_HAS_QUESTION         = 70009
	PACKAGE_NOT_EMPTY          = 70010
	SUBJECT_NAME_NOT_EMPTY     = 70011
	CHAPTER_NAME_NOT_EMPTY     = 70012
	CHAPTER_SEQUENCE_NOT_EMPTY = 70013
	CHAPTER_SUBJECT_NOT_EMPTY  = 70014

	SECTION_NAME_NOT_EMPTY     = 70015
	SECTION_SEQUENCE_NOT_EMPTY = 70016
	SECTION_CHAPTER_NOT_EMPTY  = 70017

	PIECE_NAME_NOT_EMPTY              = 70018
	PIECE_SECTION_NOT_EMPTY           = 70019
	SEQUENCE_NOT_EMPTY                = 70020
	QUESTION_TYPE_NOT_EMPTY           = 70021
	QUESTION_CONTENT_NOT_EMPTY        = 70022
	QUESTION_EMPTY                    = 70023
	QUESTION_OPTION_TITLE_NOT_EMPTY   = 70024
	QUESTION_OPTION_CONTENT_NOT_EMPTY = 70025
	OPTION_NOT_EXISTS                 = 70026
)

var ErrText = map[int64]string{
	//系统相关错误码
	SUCCES:       "操作成功",
	SERVER_ERROR: "服务器异常，请稍后再试！",
	//参数验证相关错误码
	INVALID_PARAMS:       "参数非法或缺失!",
	PARAM_ERROR:          "参数输入有误，请重新输入!",
	TOKEN_ERROR:          "token错误或不存在",
	SUBJECT_EMPTY_ERROR:  "请输入课程ID",
	PIECE_EMPTY_ERROR:    "请输入块ID",
	IMAGE_EMPTY_ERROR:    "上传的图片不能为空",
	SECURITY_EMPTY_ERROR: "请填写验证码",
	//账号问题相关错误码
	ACCOUNT_EXISTS:           "账号已存在",
	ACCOUNT_NOT_EXISTS:       "账号不存在",
	ACCOUNT_NOT_EMPTY:        "账号不能为空",
	ACCOUNT_EXCEPTION:        "账号异常",
	ACCOUNT_FORMAT_ERROR:     "电话或电子邮件格式不正确",
	ACCOUNT_PASS_ERROR:       "密码填写错误",
	ACCOUNT_BLOCKED_ERROR:    "账号已被拉黑，请联系管理员处理！",
	ACCOUNT_LOGIN_TYPE_ERROR: "用户登录类型错误",
	ACCOUNT_LOGIN_ERROR:      "请登录账号",
	MANAGER_IDENTIFIER_ERROR: "识别码错误",
	//验证码问题相关错误码
	SECURITY_CODE_EXISTS: "验证码已存在，请勿重复操作",
	SECURITY_CODE_ERROR:  "验证码输入错误",
	SMS_SEND_ERROR:       "短息发送失败",
	SEND_TYPE_ERROR:      "短息发送类型错误",
	//用户相关错误码
	USER_TRY_MAX_ERROR:          "试用机会已用完，通过积分可兑换试用机会或购买本课程学习",
	USER_TRY_SAME_ERROR:         "课程正在试用中，不能重复试用",
	USER_ID_EMPTY_ERROR:         "用户ID不能为空",
	USER_PIECEID_EMPTY_ERROR:    "节块ID不能为空",
	USER_NOT_TRY_ORBUY_ERROR:    "试用或购买该课程后才可以学习",
	USER_CONFIRM_PASSWORD_ERROR: "新密码和确认密码不一致",
	USER_PASSWORD_ERROR:         "原密码输入错误",
	//管理员相关错误码
	MANAGER_NOT_EXISTS:         "管理员不存在",
	MANAGER_EXISTS:             "管理员已存在",
	PACKAGE_NOT_EXISTS:         "专业不存在",
	PACKAGE_NOT_EMPTY:          "专业不能为空",
	PACKAGE_HAS_SUBJECT:        "专业下已经存在课程，请先删除课程",
	SUBJECT_NOT_EXISTS:         "课程不存在",
	SUBJECT_HAS_CHAPTER:        "专业下已经存在学习章，请先删除学习章",
	CHAPTER_NOT_EXISTS:         "学习章不存在",
	CHAPTER_HAS_SECTION:        "学习章下已经存在学习节，请先删除学习节",
	SECTION_NOT_EXISTS:         "学习节不存在",
	SECTION_HAS_PIECE:          "学习节下已经存在学习块，请先删除学习块",
	PIECE_NOT_EXISTS:           "学习块不存在",
	PIECE_HAS_QUESTION:         "学习块下已经存在题目，请先删除题目",
	SUBJECT_NAME_NOT_EMPTY:     "课程名称不能为空",
	CHAPTER_NAME_NOT_EMPTY:     "学习章名称不能为空",
	CHAPTER_SEQUENCE_NOT_EMPTY: "学习章名顺序不能为空",
	CHAPTER_SUBJECT_NOT_EMPTY:  "学习章所属课程不能为空",
	SECTION_NAME_NOT_EMPTY:     "学习节名称不能为空",
	SECTION_SEQUENCE_NOT_EMPTY: "学习节名顺序不能为空",
	SECTION_CHAPTER_NOT_EMPTY:  "学习节所属学习章不能为空",

	PIECE_NAME_NOT_EMPTY:              "学习块名称不能为空",
	PIECE_SECTION_NOT_EMPTY:           "学习块所属学习节名称不能为空",
	SEQUENCE_NOT_EMPTY:                "顺序不能为空",
	QUESTION_TYPE_NOT_EMPTY:           "题目类型不能为空",
	QUESTION_CONTENT_NOT_EMPTY:        "题目内容不能为空",
	QUESTION_EMPTY:                    "题目不存在",
	QUESTION_OPTION_TITLE_NOT_EMPTY:   "选项不能为空",
	QUESTION_OPTION_CONTENT_NOT_EMPTY: "选项内容不能为空",
	OPTION_NOT_EXISTS:                 "选项不存在",
}

func GetErrMsg(code int64) string {
	if v, ok := ErrText[code]; ok {
		return v
	}
	return "服务器繁忙，稍后再试！"
}

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
}

func GetErrMsg(code int64) string {
	if v, ok := ErrText[code]; ok {
		return v
	}
	return "服务器繁忙，稍后再试！"
}

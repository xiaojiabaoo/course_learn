package tools

const (
	LOGIN_VERIFY      = "【xiaoben内部学习系统】你正在进行登录，验证码是%s，请于10分内完成验证，打死你都不要告诉别人哦！"
	REGISTER_VERIFY   = "【xiaoben内部学习系统】你正在注册账号，验证码是%s，请于10分内完成验证，打死你都不要告诉别人哦！"
	RESET_VERIFY      = "【xiaoben内部学习系统】你正在重置密码，验证码是%s，请于10分内完成验证，打死你都不要告诉别人哦！"
	BIND_PHONE_VERIFY = "【xiaoben内部学习系统】你正在绑定手机，验证码是%s，请于10分内完成验证，打死你都不要告诉别人哦！"
	BIND_EMAIL_VERIFY = "【xiaoben内部学习系统】你正在绑定邮箱，验证码是%s，请于10分内完成验证，打死你都不要告诉别人哦！"
	MODIFY_VERIFY     = "【xiaoben内部学习系统】你正在修改密码，验证码是%s，请于10分内完成验证，打死你都不要告诉别人哦！"
)

const (
	SECURITY_CODE_EXPIRE   = 10 * 60               //验证码时间
	USER_ONLINE_EXPIRE     = 36000                 //用户登录的过期时间
	SECURITY_CODE_KEY      = "SECURITY_CODE:%s"    //验证码key
	USER_LOGIN_KEY         = "USER_LOGIN:%d"       //用户登录token
	USER_LOGIN_TOKEN_KEY   = "USER_LOGIN_TOKEN:%s" //用户登录token
	USER_DEFAULT_PASS      = "%s+12138"            //用户在验证码登录注册时的默认密码
	USER_DEFAULT_TRY_COUNT = 5                     //用户注册时默认的试用次数
	USER_DEFAULT_TRY_TIME  = 15                    //用户注册时每次试用的时间 单位：天
	USER_EXPIRE_DAY        = 5 * 86400             //用户试用/购买的产品还剩多少天提示即将过期

	MANAGER_ONLINE_EXPIRE   = 36000                    //管理员登录的过期时间
	MANAGER_LOGIN_KEY       = "MANAGER_LOGIN:%d"       //管理员登录token
	MANAGER_LOGIN_TOKEN_KEY = "MANAGER_LOGIN_TOKEN:%s" //管理员登录token

	USER_LOGIN    = "USER_LOGIN"
	MANAGER_LOGIN = "MANAGER_LOGIN"
)

const (
	EMPTY_INT = 0  //int类型空值
	EMPTY_STR = "" //string类型空值

	LOGIN_TYPE_PHONE = 1 //登录方式 手机
	LOGIN_TYPE_EMAIL = 2 //登录方式 邮箱

	LOGIN_TYPE_PASS   = 2 //密码登录
	LOGIN_TYPE_VERIFY = 1 //验证码登录：登录或注册
	LOGIN_WAY_MANAGER = 3 //管理员登录

	VERIFY_CODE_REGISTER   = 1 //验证码发送类型：注册
	VERIFY_CODE_RESET      = 2 //验证码发送类型：重置密码
	VERIFY_CODE_BIND_PHONE = 3 //验证码发送类型：绑定手机
	VERIFY_CODE_BIND_EMAIL = 4 //验证码发送类型：绑定邮箱
	VERIFY_CODE_LOGIN      = 5 //验证码发送类型：登录
	VERIFY_CODE_MODIFY     = 6 //验证码发送类型：修改密码

)

package req_model

type Verification struct {
	Account  string `json:"account" validate:"required"`   //手机号码/邮箱地址
	SendType int64  `json:"send_type" validate:"required"` //发送信息类型 1：注册，2：重置密码，3：绑定手机，4：绑定邮箱，5：登录，6：修改密码，如果是3、4，需要传递token
	AreaCode string `json:"area_code" desc:"区号"`           //区号，格式如：+86 该区号主要用于发送，如果是手机注册或绑定手机等操作，需要填写该区号字段，重置密码时区号可以直接和手机拼接组合
}

type Login struct {
	LoginWay     int64  `json:"login_way" validate:"required"`  //用户登录方式：1.手机号码登录；2.邮箱地址登录；3.管理员登录
	LoginType    int64  `json:"login_type" validate:"required"` //用户登录类型：1.验证码登录；2.密码登录
	Account      string `json:"account" validate:"required"`    //用户注册手机号码/邮箱地址
	SecurityCode int64  `json:"security_code"`                  //验证码
	Password     string `json:"password"`                       //密码
	Identifier   string `json:"identifier"`                     //识别码：管理员登录时需要
	AreaCode     string `json:"area_code"`                      //用户手机区号，非手机注册时该字段无效
	IP           string `json:"ip"`                             //用户注册IP地址
}

type Upload struct {
	Id       int64  `form:"img_id"`
	FileType string `form:"file_type"`
}

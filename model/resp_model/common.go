package resp_model

type Login struct {
	Token    string `json:"token"`     //登录信息
	UserId   int64  `json:"user_id"`   //用户id
	Account  string `json:"account"`   //用户账号名称
	NickName string `json:"nick_name"` //用户昵称
	Logo     string `json:"logo"`      //用户头像
}

type Upload struct {
	Src string `json:"src"`
}

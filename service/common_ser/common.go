package common_ser

import (
	"course_learn/common/redisx"
	"course_learn/dao/common_dao"
	"course_learn/dao/user_dao"
	"course_learn/model/mysql_model"
	"course_learn/model/req_model"
	"course_learn/model/resp_model"
	"course_learn/tools"
	custErr "course_learn/tools/error"
	"course_learn/tools/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"mime/multipart"
	"time"
)

type Common interface {
	Login(param req_model.Login) (resp_model.Login, error)
	Verification(param req_model.Verification) error
	GetUserInfo(c *gin.Context) (int, error)
	Upload(file *multipart.FileHeader) (resp_model.Upload, error)
}

type ApiCommon struct{}

func (a *ApiCommon) Login(param req_model.Login) (resp_model.Login, error) {
	var (
		err       error
		code      string
		now       = time.Now()
		uuid      = util.CreateUUID()
		apiUser   = user_dao.ApiUser{}
		login     = resp_model.Login{}
		user      = &mysql_model.Users{}
		apiCommon = common_dao.ApiCommon{}
	)
	switch param.LoginType {
	case tools.LOGIN_TYPE_PASS:
		//查询登录的账号信息
		user, err = apiCommon.GetUserDataByAccount(param.Account)
		if err != nil {
			return login, err
		}
		if user.Id == tools.EMPTY_INT {
			return login, custErr.New(custErr.ACCOUNT_NOT_EXISTS, "")
		}
		//查询密码是否输入正确
		if user.Password != param.Password {
			return login, custErr.New(custErr.ACCOUNT_PASS_ERROR, "")
		}
		//登录成功，把token信息存入redis
		err = a.login(user.Id, uuid)
		if err != nil {
			return login, err
		}
	case tools.LOGIN_TYPE_VERIFY:
		//检测验证码是否正确
		code = redisx.Redis.Get(fmt.Sprintf(tools.SECURITY_CODE_KEY, param.Account))
		if err != nil {
			return login, errors.Wrap(err, "获取验证码失败")
		}
		if param.SecurityCode != int64(util.Str2Int(code)) {
			//验证码错误
			return login, custErr.New(custErr.SECURITY_CODE_ERROR, "")
		}
		err = redisx.Redis.Del(fmt.Sprintf(tools.SECURITY_CODE_KEY, param.Account))
		if err != nil {
			return login, errors.Wrap(err, "删除验证码失败")
		}
		//验证码通过后查询账号信息
		user, err = apiCommon.GetUserDataByAccount(param.Account)
		if err != nil {
			return login, err
		}
		if user.Id == tools.EMPTY_INT {
			//没有账号-注册
			users := &mysql_model.Users{
				Email:         param.Account,
				Password:      fmt.Sprintf(tools.USER_DEFAULT_PASS, param.Account),
				AddTime:       now.Unix(),
				UpdateTime:    now.Unix(),
				DateTime:      now,
				AccountStatus: 0,
			}
			err = apiUser.AddUsers(users)
			if err != nil {
				return login, err
			}
			//添加用户参数
			userParam := &mysql_model.UserParam{
				UserId:   users.Id,
				TryCount: tools.USER_DEFAULT_TRY_COUNT,
				TryTime:  tools.USER_DEFAULT_TRY_TIME,
				AddTime:  now.Unix(),
				UpTime:   now.Unix(),
				DateTime: now,
			}
			err = apiUser.AddUserParam(userParam)
			if err != nil {
				return login, err
			}
			user.Id = users.Id
		}
		//登录
		err = a.login(user.Id, uuid)
		if err != nil {
			return login, err
		}
	default:
		return login, custErr.New(custErr.ACCOUNT_LOGIN_TYPE_ERROR, "")
	}
	//添加登录日志
	logs := &mysql_model.UserLoginLogs{
		UserId:      user.Id,
		LoginTime:   now.Unix(),
		LoginIp:     "127.0.0.1",
		LoginArea:   "中国大陆",
		LoginMethod: param.LoginWay,
		DateTime:    now,
	}
	err = apiUser.AddUserLoginLogs(logs)
	if err != nil {
		return login, err
	}
	login.UserId = user.Id
	login.Account = param.Account
	login.NickName = user.NickName
	login.Logo = user.Logo
	login.Token = uuid
	return login, nil
}

func (a *ApiCommon) login(userId int64, uuid string) error {
	//先把原有的登录状态信息删除
	err := redisx.Redis.Del(fmt.Sprintf(tools.USER_LOGIN_KEY, userId))
	if err != nil {
		return errors.Wrap(err, "删除用户原有Token失败")
	}
	err = redisx.Redis.Del(fmt.Sprintf(tools.USER_LOGIN_TOKEN_KEY, uuid))
	if err != nil {
		return errors.Wrap(err, "删除用户原有Token失败")
	}
	//添加新的登录信息
	err = redisx.Redis.Set(fmt.Sprintf(tools.USER_LOGIN_KEY, userId), uuid, tools.USER_ONLINE_EXPIRE)
	if err != nil {
		return errors.Wrap(err, "用户登录redis存储token失败")
	}
	err = redisx.Redis.Set(fmt.Sprintf(tools.USER_LOGIN_TOKEN_KEY, uuid), userId, tools.USER_ONLINE_EXPIRE)
	if err != nil {
		return errors.Wrap(err, "用户登录redis存储token失败")
	}
	return nil
}

func (a *ApiCommon) Verification(param req_model.Verification) error {
	var (
		body    string
		err     error
		codeNum = util.RandNum(100000, 999999) //生成验证码
	)
	if param.Account == "" {
		return custErr.New(custErr.ACCOUNT_NOT_EMPTY, "")
	}
	if !util.VerifyEmail(param.Account) {
		return custErr.New(custErr.ACCOUNT_FORMAT_ERROR, "")
	}

	switch param.SendType {
	case tools.VERIFY_CODE_REGISTER:
		body = fmt.Sprintf(tools.REGISTER_VERIFY, codeNum)
	case tools.VERIFY_CODE_LOGIN:
		body = fmt.Sprintf(tools.LOGIN_VERIFY, codeNum)
	case tools.VERIFY_CODE_MODIFY:
		body = fmt.Sprintf(tools.MODIFY_VERIFY, codeNum)
	default:
		return custErr.New(custErr.SECURITY_CODE_ERROR, "")
	}
	err = tools.SendMail([]string{param.Account}, body, codeNum)
	if err != nil {
		return err
	}
	err = redisx.Redis.Set(fmt.Sprintf(tools.SECURITY_CODE_KEY, param.Account), codeNum, tools.SECURITY_CODE_EXPIRE)
	if err != nil {
		return errors.Wrap(err, "发送验证码redis存储验证码失败")
	}
	return nil
}

func (a *ApiCommon) GetUserInfo(c *gin.Context) (int64, error) {
	token := c.GetHeader("token")
	if token == "" {
		return 0, custErr.New(custErr.TOKEN_ERROR, "")
	}
	userId := util.Str2Int(redisx.Redis.Get(fmt.Sprintf(tools.USER_LOGIN_TOKEN_KEY, token)))
	if userId == 0 {
		return 0, custErr.New(custErr.ACCOUNT_LOGIN_ERROR, "")
	}
	return int64(userId), nil
}

func (a *ApiCommon) Upload(file *multipart.FileHeader) (resp_model.Upload, error) {
	upload := resp_model.Upload{}
	files, err := file.Open()
	if err != nil {
		return upload, err
	}
	defer files.Close()
	/*fileByte, err := ioutil.ReadAll(files) //获取上传文件字节流
	if err != nil {
		return upload, err
	}*/
	return upload, nil

}

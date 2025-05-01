package router

import (
	"course_learn/controller/chapter"
	"course_learn/controller/common"
	"course_learn/controller/question"
	"course_learn/controller/subject"
	"course_learn/controller/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func Run(addr string) {
	router := gin.Default()
	router.Use(Cors())
	commons := router.Group("/common")
	{
		commons.POST("/login", common.Login)         // 用户登录接口（和注册接口共用）
		commons.POST("/verify", common.Verification) // 验证码发送接口
		commons.POST("/upload", common.Upload)       // 图片上传接口
	}
	subjects := router.Group("/subject")
	{
		get := subjects.Group("/get")
		{
			get.POST("/list", subject.GetSubjectList)     // 课程列表
			get.POST("/detail", subject.GetSubjectDetail) // 课程详情
		}
		set := subjects.Group("/set")
		{
			set.POST("/try", subject.TrySubject) // 课程试用
		}
	}
	chapters := router.Group("/chapter")
	{
		get := chapters.Group("/get")
		{
			get.POST("/list", chapter.GetChapterList) // 节块列表
		}
	}
	questions := router.Group("/question")
	{
		get := questions.Group("/get")
		{
			get.POST("/list", question.GetQuestionData)
		}
	}
	users := router.Group("/user")
	{
		get := users.Group("/get")
		{
			get.POST("/lean/record", user.GetUserLeanRecord)
			get.POST("/data", user.GetUserData)
			get.POST("/login_record", user.GetLoginRecord)
			get.POST("/order_list", user.GetOrderList)
			get.POST("/learn_record", user.GetLearnRecord)
		}
		set := users.Group("/set")
		{
			set.POST("/lean/record", user.AddUserLeanRecord)
			set.POST("/data", user.SetUserData)
			set.POST("/change_password", user.ChangePassword)
		}
	}
	InitWebRouter(router)
	if err := router.Run(addr); err != nil {
		panic(err)
	}

}

package response

import (
	custom "course_learn/tools/error"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 返回响应
func ResponseData(c *gin.Context, err error, data interface{}) {
	if data == "" || data == nil {
		data = make([]int, 0)
	}
	resp := &Response{Data: data, Code: 0, Msg: custom.ErrText[custom.SUCCES]}
	if err != nil { // 这里是错误模式
		if tem, ok := err.(*custom.CustomError); ok {
			//var msg = custom.ErrText[tem.Code]
			resp = &Response{Code: tem.Code, Msg: tem.Error(), Data: data}
		} else {
			// 包装一下
			stack := fmt.Sprintf("stack error trace:\n%+v\n", err) //错误的堆栈
			fmt.Println(stack)
			//返回给前端
			resp = &Response{Code: custom.SERVER_ERROR, Msg: custom.ErrText[custom.SERVER_ERROR], Data: make([]int, 0)}
		}
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Current-Time", strconv.Itoa(int(time.Now().Unix())))
	c.JSON(200, resp)
	c.Abort()
}

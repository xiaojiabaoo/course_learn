package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitWebRouter(engine *gin.Engine) {
	//engine.Use(common.TranslationMiddleware())
	engine.Static("/static/css", "./common/static/css")
	engine.Static("/static/js", "./common/static/js")
	engine.Static("/static/image", "./common/static/image")
	engine.Static("/subject", "./common/static/html/subject")
	engine.Static("/about", "./common/static/html/about")
	engine.Static("/activity", "./common/static/html/activity")
	engine.Static("/integral", "./common/static/html/integral")
	engine.Static("/topic", "./common/static/html/topic")
	engine.Static("/user_html", "./common/static/html/user")
	engine.LoadHTMLFiles(
		"./common/static/html/main.html", "./common/static/html/login.html",
		"./common/static/html/404.html", "./common/static/html/500.html")
	engine.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", gin.H{})
	})
	engine.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{})
	})
	engine.GET("/500", func(c *gin.Context) {
		c.HTML(http.StatusOK, "500.html", gin.H{})
	})
	engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
}

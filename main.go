package main

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Frontend Views
func indexWeb(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"web_title": "Talk Bar",
		"title":     "Welcome to TalkBar",
		"context1":  "Anything can say",
		"context2":  "You're freedom to communite sex,relation,etc",
	})
}

func context(c *gin.Context) {
	c.HTML(http.StatusOK, "context.html", gin.H{
		"web_title": "Community",
	})
}

// Backend Manager
func userLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"web_title":   "Manger Login",
		"main_title":  "Talkger Login",
		"manager_id":  "Manger ID",
		"manager_pwd": "Manger Password",
	})
}

func loginSuccessful(c *gin.Context) {
	_name := c.DefaultPostForm("user_name", "GUEST")
	_pwd := c.DefaultPostForm("user_pwd", "None")
	fmt.Println(reflect.TypeOf(_name))
	fmt.Println(reflect.TypeOf(_pwd))
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"ID":     _name,
		"pwd":    _pwd,
	})
}

// For Test
func testGin(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", gin.H{
		"web_title": "For Test",
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	views := router.Group("/views")
	{
		views.GET("/index", indexWeb)
		views.POST("/context", context)
	}
	api := router.Group("/api")
	{
		api.GET("/login", userLogin)
		api.POST("/loginSuccessful", loginSuccessful)
	}
	fortest := router.Group("fortest")
	{
		fortest.GET("/index", testGin)
	}
	router.Run(":1450")
}

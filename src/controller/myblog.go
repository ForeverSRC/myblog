package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyblogGet(c *gin.Context){
	isLogin:=checkAccount(c)//通过验证Cookie判断是否已经登录
	c.HTML(http.StatusOK,"myblog.html",gin.H{
		"IsLogin":isLogin,
	})
}

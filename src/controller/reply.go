package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/src/models"
	"net/http"
)

func AddReplyPost(c *gin.Context){
	tid:=c.PostForm("tid")
	nickname:=c.PostForm("nickname")
	content:=c.PostForm("content")
	models.AddReply(tid,nickname,content)
	c.Redirect(http.StatusFound,"/topic/view/"+tid)
}

func DeleteReply(c *gin.Context){
	//检查是否登录
	if !checkAccount(c){
		c.Redirect(http.StatusFound,"/login")
		return
	}
	rid:=c.Query("rid")
	tid:=c.Query("tid")
	models.DeleteReply(rid,tid)
	c.Redirect(http.StatusFound,"/topic/view/"+tid)
}


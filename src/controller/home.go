package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/src/models"
	"net/http"
)

//访问首页
func HomeGet(c *gin.Context){
	isLogin:=checkAccount(c)//通过验证Cookie判断是否已经登录
	cid:=c.Query("cid")
	label:=c.Query("label")
	var topics []*models.Topic
	if len(cid)==0&&len(label)==0{
		topics=models.GetAllTopics(true)
	}else if len(cid)!=0&&len(label)==0{
		topics=models.GetTopicByCate(cid)
	}else{
		topics=models.GetTopicByLabel(label)
	}

	cates:=models.GetAllCategories()
	c.HTML(http.StatusOK,"home.html",gin.H{
		"IsHome":true,
		"IsLogin":isLogin,
		"topics":topics,
		"Categories":cates,
	})
}

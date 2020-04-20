package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/src/models"
	"github.com/src/mylog"
	"net/http"
)

func TopicGet(c *gin.Context){
	topics:=models.GetAllTopics(false)
	c.HTML(http.StatusOK,"topic.html",gin.H{
		"IsTopic":true,
		"IsLogin":checkAccount(c),
		"Topics":topics,
	})
}

func TopicPost(c *gin.Context) {
	if IsAdm(c) {
		//解析提交的表单
		//1.获取表单数据
		title := c.PostForm("title")
		content := c.PostForm("content")
		cid := c.PostForm("cid")
		labels := c.PostForm("labels")
		//2.判断tid
		tid := c.PostForm("tid")

		if len(tid) == 0 {
			//没有tid，新增文章
			//2.提交到数据库
			models.AddTopic(title, cid, labels, content)
			mylog.Logger.Printf("Administrator add a topic: [%s].",title)
		} else {
			//tid存在，修改文章
			models.ModifyTopic(tid, title, cid, labels, content)
			mylog.Logger.Printf("Administrator modified the topic: [tid=%s].",tid)
		}
		//重定向
		c.Redirect(http.StatusFound, "/topic")
	}else{
		mylog.Logger.Println("An visitor attempted to add/modify a topic.")
	}
}

func TopicAdd(c *gin.Context) {
	if IsAdm(c){
		//获取所有现存分类
		mylog.Logger.Println("Administrator begin to add a topic.")
		categories := models.GetAllCategories()
		c.HTML(http.StatusOK, "topic_add.html", gin.H{
			"IsTopic":    true,
			"IsLogin":    checkAccount(c),
			"Categories": categories,
		})
	}else{
		mylog.Logger.Println("A visitor attempted to add a topic.")
	}
}

func TopicView(c *gin.Context){
	//获取ID
	tid:=c.Param("tid")
	topic,err:=models.GetTopic(tid)
	if err!=nil{
		c.Redirect(http.StatusFound,"/")
		return
	}
	replies:=models.GetAllReplies(tid)
	c.HTML(http.StatusOK,"topic_view.html",gin.H{
		"IsTopic":true,
		"IsLogin":checkAccount(c),
		"Topic":topic,
		"Replies":replies,
	})
}

func TopicModify(c *gin.Context) {
	if IsAdm(c){
		tid := c.Query("tid")
		mylog.Logger.Printf("Administrator begins to modify the topic: [tid=%s].",tid)
		topic, err := models.GetTopic(tid)
		if err != nil {
			c.Redirect(http.StatusFound, "/")
			return
		}
		//获取所有现存分类
		categories := models.GetAllCategories()

		c.HTML(http.StatusOK, "topic_modify.html", gin.H{
			"IsTopic":    true,
			"IsLogin":    checkAccount(c),
			"Topic":      topic,
			"Tid":        tid,
			"Categories": categories,
			"CurCID":     topic.CID,
		})
	}else{
		mylog.Logger.Println("A visitor attempted to modify a topic.")
	}
}

func TopicDelete(c *gin.Context){
	//检查是否为管理员登录
	IsAdm(c)
	models.DeleteTopic(c.Query("tid"))
	c.Redirect(http.StatusFound,"/topic")
}

func ImgUpPost(c *gin.Context){
	pic,_:=c.FormFile("editormd-image-file")
	dst:=fmt.Sprintf("./statics/upload_img/%s", pic.Filename)
	c.SaveUploadedFile(pic, dst)
	c.JSON(http.StatusOK, gin.H{
		"success":1,
		"message":"上传成功",
		"url":"/static/upload_img/"+pic.Filename,
	})
}

func Download_pdf(c *gin.Context){
	tid:=c.Param("tid")
	topic,err:=models.GetTopic(tid)
	if err!=nil{
		c.Redirect(http.StatusFound,"/")
		return
	}
	c.HTML(http.StatusOK,"topic_download_pdf.html",gin.H{
		"Topic":topic,
	})
}
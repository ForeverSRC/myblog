package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/src/models"
	"github.com/src/mylog"
	"net/http"
)

func CategoryGet(c *gin.Context){
	cates:=models.GetAllCategories()
	c.HTML(http.StatusOK,"category.html",gin.H{
		"IsCategory":true,
		"IsLogin":checkAccount(c),
		"Categories":cates,
	})
}

func CategoryPost(c *gin.Context){
	if IsAdm(c){
		name:=c.PostForm("name")
		if len(name)!=0{
			models.AddCategory(name)
		}
		//重定向以刷新分类列表
		c.Redirect(http.StatusFound,"/category")
		mylog.Logger.Println("Administrator add a category.")
	}else{
		mylog.Logger.Println("A visitor attempted to add a category.")
	}
	return
}

func CategoryDel(c *gin.Context){
	if IsAdm(c){
		id:=c.Param("id")
		if len(id)!=0{
			models.DelCategory(id)
		}
		c.Redirect(http.StatusFound,"/category")
		mylog.Logger.Println("Administrator deleted a category.")
	}else{
		mylog.Logger.Println("A visitor attempted to delete a category.")
	}
	return
}
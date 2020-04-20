package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/src/controller"
	"github.com/src/functions"
	"github.com/src/mylog"
	"html/template"
	"io"
	"github.com/src/models"
	"net/http"
	"os"
	"os/signal"
	"time"
)



func init(){
	models.RegisterDB()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(mylog.Writer)
	gin.Logger()
	models.Db.LogMode(true)
	models.Db.SetLogger(mylog.Logger)

}
func main() {
	r:=gin.Default()
	//静态文件服务
	r.Static("/static","./statics")

	//相册服务
	r.Static("/images","./albums")

	//定义模板函数
	r.SetFuncMap(template.FuncMap{
		"decodeLabels":functions.DecodingLabels,
		"decodeLabelsIterable":functions.DecodingLabelsIterable,
		"unescaped":functions.Unescaped,
	})

	//处理模板文件
	r.LoadHTMLGlob("./templates/*.*")

	//路由
	r.GET("/myblog",controller.MyblogGet)
	r.GET("/", controller.HomeGet)

	r.GET("/login", controller.LoginGet)
	r.GET("login/exit",controller.LoginExit)
	r.POST("/login",controller.LoginPost)


	r.GET("/category",controller.CategoryGet)
	r.POST("/category/add",controller.CategoryPost)
	r.GET("/category/del/:id",controller.CategoryDel)

	r.GET("/topic",controller.TopicGet)
	r.POST("/topic",controller.TopicPost)
	r.Any("/topic/add",controller.TopicAdd)
	r.POST("/topic/imgup",controller.ImgUpPost)

	r.GET("/topic/view/:tid",controller.TopicView)
	r.GET("topic/download_pdf/:tid",controller.Download_pdf)
	r.GET("/topic/modify",controller.TopicModify)

	r.GET("/topic/delete",controller.TopicDelete)

	r.POST("/reply/add",controller.AddReplyPost)
	r.GET("/reply/delete",controller.DeleteReply)

	r.GET("/todolist",controller.TodoList)

	r.GET("/albums",controller.AlbumGet)
	r.GET("/albums/data",controller.AlbumGetAll)
	r.POST("/albums/data",controller.AlbumPostNew)
	r.DELETE("/albums/data",controller.AlbumDelete)
	r.GET("/albums/album/:id",controller.AlbumPhotoGet)

	r.POST("/photo/get",controller.PhotoGetAll)
	r.POST("/photo/data",controller.PhotoPostNew)
	r.DELETE("/photo/data",controller.PhotoDelete)

	todoGroup:=r.Group("/todolist")
	{
		//添加待办事项
		todoGroup.POST("/todo",controller.AddTodo)
		//查看所有待办事项
		todoGroup.GET("/todo",controller.ViewAllTodo)

		//修改待办事项状态
		todoGroup.PUT("/todo",controller.ModifyTodo)
		//修改待办事项内容
		todoGroup.PUT("/todo/modify",controller.SetTodoContent)
		//删除待办事项
		todoGroup.DELETE("/todo/",controller.DeleteTodo)

		//编辑某一待办事项的具体内容
		todoGroup.GET("/todo/modify/:id",controller.EditTodo)
	}
	srv := &http.Server{
		Addr:    ":80",
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mylog.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	mylog.Logger.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		mylog.Logger.Fatal("Server Shutdown:", err)
	}
	mylog.Logger.Println("Server exiting")
	defer models.Db.Close()
}

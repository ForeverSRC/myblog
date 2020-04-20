package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/src/models"
	"net/http"
)

func TodoList(c *gin.Context){
	IsAdm(c)
	//返回页面
	c.HTML(http.StatusOK,"todo.html",gin.H{
		"IsTodo":true,
		"IsLogin":checkAccount(c),
	})
}

func AddTodo(c *gin.Context){
	if !IsAdm(c){
		c.JSON(http.StatusNotFound,gin.H{
			"code":233,//写入数据库失败
			"msg":"非法访问",
		})
		return
	}
	//前端填写待办事项，提交到后端
	//取数据
	var todo models.Todo;
	c.BindJSON(&todo);
	//存入数据库
	err:=models.AddTodo(&todo);
	//返回响应
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,//写入数据库成功
			"msg":"添加成功",
		});
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":233,//写入数据库失败
			"msg":"添加失败，错误：\n"+err.Error(),
		})
	}

}

func ViewAllTodo(c *gin.Context){
	if !IsAdm(c){
		c.JSON(http.StatusNotFound,gin.H{
			"code":233,//写入数据库失败
			"msg":"非法访问",
		})
		return
	}
	var list []models.Todo;
	err:=models.GetAllTodo(&list);
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,//写入数据库成功
			"msg":"success",
			"data":list,
		});
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":234,//获取失败
			"msg":err.Error(),
		})
	}
}

func ModifyTodo(c *gin.Context){
	if !IsAdm(c){
		c.JSON(http.StatusNotFound,gin.H{
			"code":233,//写入数据库失败
			"msg":"非法访问",
		})
		return
	}
	var todo models.Todo;
	c.BindJSON(&todo);
	err:=models.ChangeStatus(&todo);
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,//写入数据库成功
			"msg":"事项设置成功",
		});
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":235,//获取失败
			"msg":"错误："+err.Error(),
		})
	}
}

func SetTodoContent(c *gin.Context){
	if !IsAdm(c){
		c.JSON(http.StatusNotFound,gin.H{
			"code":233,//写入数据库失败
			"msg":"非法访问",
		})
		return
	}
	var todo models.Todo;
	c.BindJSON(&todo);
	err:=models.SetContent(&todo);
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,//写入数据库成功
			"msg":"编辑成功",
		});
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":235,//获取失败
			"msg":"错误："+err.Error(),
		})
	}
}

func DeleteTodo(c *gin.Context){
	if !IsAdm(c){
		c.JSON(http.StatusNotFound,gin.H{
			"code":233,//写入数据库失败
			"msg":"非法访问",
		})
		return
	}
	var todo models.Todo;
	c.BindJSON(&todo);
	err:=models.DeleteTodo(&todo)
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,//写入数据库成功
			"msg":"删除成功",
		});
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":235,//获取失败
			"msg":"错误："+err.Error(),
		})
	}
}

func EditTodo(c *gin.Context){
	if IsAdm(c){
		id:=c.Param("id")
		todo:=models.GetTodo(id)
		c.HTML(http.StatusOK,"todo_edit.html",gin.H{
			"IsTodo":true,
			"IsLogin":checkAccount(c),
			"Todo":todo,
		})
	}
}
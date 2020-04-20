package controller

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//管理员密码
func md5pwd() string{
	h := md5.New()
	h.Write([]byte("src1234"))
	cipherStr := h.Sum(nil)
	dst:=fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
	return dst
}

func md5fy(str string)string{
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	dst:=fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
	return dst
}

//登录页面
func LoginGet(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",nil)
}

//退出登录
func LoginExit(c *gin.Context){
	//重新设置Cookie
	c.SetCookie("uname", "", -1, "/", "", false, true)
	c.SetCookie("pwd", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound,"/")
}

//登录请求
func LoginPost(c *gin.Context){
	username:=c.PostForm("uname")
	password:=c.PostForm("pwd")
	autoLogin:=c.PostForm("autologin")=="on"

	//验证登录
	//此处采用硬编码模式，后期改为数据库模式
	password=md5fy(password)//md5加密
	if username=="src"&&password==md5pwd(){
		maxAge:=0
		if autoLogin{
			maxAge=1<<20
		}
		c.SetCookie("uname",username,maxAge,"/","",false,true)
		c.SetCookie("pwd",password,maxAge,"/","",false,true)
	}
	//重定向
	c.Redirect(http.StatusFound,"/")
}

//验证cookie
func checkAccount(c *gin.Context) bool{
	//获取cookie
	uname,err:=c.Cookie("uname")
	if err!=nil{
		return false
	}

	pwd,err:=c.Cookie("pwd")
	if err!=nil{
		return false
	}

	//补充：验证结果提示
	return uname=="src"&&pwd==md5pwd()
}
package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAdm(c *gin.Context) bool{
	if !checkAccount(c){
		c.Redirect(http.StatusFound,"/login")
		return false
	}
	return true
}

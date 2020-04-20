package mylog

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"log"
	"time"
)

var Logger *log.Logger

var Writer *rotatelogs.RotateLogs

func init(){
	//日志文件
	path:="./log/ginlog"
	Writer,_=rotatelogs.New(path+"%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(168)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	Logger=log.New(Writer,"[System]",log.LstdFlags|log.Lshortfile)
}
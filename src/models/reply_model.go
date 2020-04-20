package models

import (
	"strconv"
	"time"
)

//评论
type Comment struct{
	ID int64
	Tid int64
	Name string
	Content string `gorm:"size:5000"`
	Created time.Time
}

func AddReply(tid,nickname,content string){
	tidNum,_:=strconv.ParseInt(tid,10,64)
	reply:=Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	Db.Create(&reply)
	ModifyReplyCount(tid,true,1)
}


func GetAllReplies(tid string) []*Comment{
	tidNum,_:=strconv.ParseInt(tid,10,64)
	replies:=make([]*Comment,0)

	//SELECT * FROM comments WHERE tid=tidNum ORDER BY Created desc
	Db.Where(&Comment{Tid:tidNum}).Order("created desc").Find(&replies)
	return replies
}

func DeleteReply(rid,tid string){
	ridNum,_:=strconv.ParseInt(rid,10,64)
	reply:=Comment{ID:ridNum}
	Db.Unscoped().Delete(&reply)
	//一定要删完再处理回复数及最新回复时间
	ModifyReplyCount(tid,false,1)
}

func GetLatestReply(tid string) *Comment{
	//获得最新的评论
	var reply Comment
	Db.Where("tid",tid).Order("created desc").First(&reply)
	return &reply
}
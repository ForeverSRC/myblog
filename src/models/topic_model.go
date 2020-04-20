package models

import (
	"github.com/src/functions"
	"strconv"
	"time"
)

type Topic struct {
	ID int64
	Uid int64
	Title string
	CID int64
	Labels string
	Content string `gorm:"type:longtext"`
	Attachment string
	Created time.Time
	Updated time.Time
	Views int64
	Author string
	ReplyTime time.Time
	ReplyCount int64 `gorm:"DEFAULT(0)"`
	ReplyLastUserId int64
}

func AddTopic(title,cid,label,content string){
	//处理标签
	label=functions.EncodingLabels(label)

	cidNum,_:=strconv.ParseInt(cid,10,64)
	topic:=&Topic{
		Title:           title,
		CID:             cidNum,
		Labels:          label,
		Content:         content,
		Created:         time.Now(),
		Updated:         time.Now(),
	}
	if Db.NewRecord(topic){
		Db.Create(topic)
		ModifyTopicCount(cid,true,1)
	}
}
func GetAllTopics(isDesc bool) []*Topic{
	topics:=make([]*Topic,0)
	if isDesc{
		Db.Order("updated desc").Find(&topics)
	}else{
		Db.Find(&topics)
	}
	return topics
}

func GetTopicByCate(cid string) []*Topic{
	cidNum,_:=strconv.ParseInt(cid,10,64)
	topics:=make([]*Topic,0)
	//SELECT * FROM topics WHERE cid=cid
	Db.Where(&Topic{CID:cidNum}).Find(&topics)
	return topics
}

func GetTopicByLabel(label string) []*Topic{
	topics:=make([]*Topic,0)
	//SELECT * FROM topics WHERE cid=cid
	Db.Where("labels LIKE ?","%$"+label+"#%").Find(&topics)
	return topics
}

func GetTopic(tid string) (*Topic,error){
	tidNum,err:=strconv.ParseInt(tid,10,64)
	if err!=nil{
		return nil,err
	}
	topic:=new(Topic)
	Db.Where(&Topic{ID:tidNum}).First(topic)
	topic.Views++
	Db.Model(topic).Update(Topic{Views:topic.Views})
	return topic,nil
}

func ModifyTopic(tid,title,cid,label,content string){
	tidNum,_:=strconv.ParseInt(tid,10,64)
	cidNum,_:=strconv.ParseInt(cid,10,64)
	topic:=new(Topic)
	Db.Where(&Topic{ID:tidNum}).First(topic)
	//之前的cid
	cid_pre:=topic.CID
	//处理标签
	label=functions.EncodingLabels(label)
	//更新文章
	Db.Model(topic).Update(Topic{
		Title:title,
		CID:cidNum,
		Labels:label,
		Content:content,
		Updated:time.Now(),
	})
	//如果cid更改
	if cidNum!=cid_pre{
		//上一个Category的TopicCount减1
		cid_p_s:=strconv.FormatInt(cid_pre,10)
		ModifyTopicCount(cid_p_s,false,1)
		//新的Category的TopicCount加1
		ModifyTopicCount(cid,true,1)
	}
}

func DeleteTopic(tid string){
	topic,_:=GetTopic(tid)
	//更新对应分类下的计数
	ModifyTopicCount(strconv.FormatInt(topic.CID,10) ,false,1)
	//删除文章
	Db.Unscoped().Delete(&topic)


}

func ModifyReplyCount(tid string,cal bool,num int64){
	//获取Topic指针
	topic,_:=GetTopic(tid)
	//计算
	if cal{
		//增加
		topic.ReplyCount+=num
		topic.ReplyTime=time.Now()
	}else{
		//减少
		topic.ReplyCount-=num
		//获取最新评论
		reply:=GetLatestReply(tid)
		if reply.ID==0{
			//没有评论了
			topic.ReplyTime,_=time.Parse("2006/01/02 15:04:05","2006/01/02 15:04:05")
		}else{
			topic.ReplyTime=reply.Created
		}

	}
	//更新数据库
	Db.Model(topic).Update(&topic)
}
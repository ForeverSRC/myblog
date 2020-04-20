package models

import (
	"strconv"
	"time"
)

type Category struct {
	ID int64
	Title string `gorm:"UNIQUE"`
	Created time.Time
	Views int64
	TopicTime time.Time
	TopicCount int64
	TopicLastUserID int64
}

//添加Category
func AddCategory(name string){
	cate:=Category{
		Title:name,
		Created:time.Now(),
	}
	//name没有被占用则创建，name被占用则返回
	Db.Create(&cate)
}

//获取所有Category
func GetAllCategories() []*Category{
	cates:=make([]*Category,0)
	Db.Find(&cates)
	return cates
}

//获取某个Category
func GetCategory(id string) *Category{
	cid,_:=strconv.ParseInt(id,10,64)
	cate:=new(Category)
	Db.Where(&Category{ID:cid}).First(cate)
	return cate
}

//删除某个Category
func DelCategory(id string) error{
	cid,err:=strconv.ParseInt(id,10,64)
	if err!=nil{
		return err
	}
	cate:=Category{ID:cid}
	Db.Unscoped().Delete(&cate)
	return nil
}

//修改某个Category的TopicCount
func ModifyTopicCount(cid string,cal bool,num int64){
	//获取Category指针
	cate:=GetCategory(string(cid))
	//计算
	if cal{
		//增加
		cate.TopicCount+=num
	}else{
		//减少
		cate.TopicCount-=num
	}
	//更新数据库
	Db.Model(cate).Update(Category{TopicCount:cate.TopicCount})
}
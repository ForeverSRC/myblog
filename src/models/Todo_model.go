package models

import (
	"strconv"
	"time"
)

type Todo struct{
	ID int64 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content" gorm:"size:5000" gorm:"default:''"`
	Created time.Time `json:"created"`
	Status bool `json:"status"`

}

func AddTodo(todo *Todo) (err error){
	todo.Created=time.Now();
	return Db.Create(todo).Error;
}

func GetAllTodo(list *[]Todo) (err error){
	return Db.Find(list).Error;
}

func ChangeStatus(todo *Todo) (err error){
	return Db.Model(todo).Where("id=?",todo.ID).Update("status",todo.Status).Error;
}

func SetContent(todo *Todo)(err error){
	return Db.Model(todo).Where("id=?",todo.ID).Update("content",todo.Content).Error;
}

func DeleteTodo(todo *Todo)(err error){
	return Db.Unscoped().Delete(todo).Error
}

func GetTodo(id string)(*Todo){
	idNum,_:=strconv.ParseInt(id,10,64)
	todo:=new(Todo)
	Db.Where(&Todo{ID:idNum}).First(todo)
	return todo
}
package models

import (
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

type Album struct {
	ID int64          `json:"id"`
	Name string       `json:"name" gorm:"unique"`
	FrontPage string  `json:"frontpage"`
	URL string        `json:"url"`
	Numbers int64     `json:"numbers"`
}

const DefaultFrontPate="http://gss0.baidu.com/-4o3dSag_xI4khGko9WTAnF6hhy/zhidao/wh%3D450%2C600/sign=acd65d7b8018367aaddc77d91b43a7e2/bba1cd11728b4710a5956d9ec5cec3fdfc03232d.jpg"

type Photo struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	AID int64 	`json:"aid"`
	Path string `json:"path"`
}

func GetAllAlbums() ([]*Album,error){
	albums:=make([]*Album,0)
	err:=Db.Find(&albums).Error
	return albums,err
}
func CreateAlbum(album *Album)(err error){
	album.URL="/albums/album/"+album.Name
	album.FrontPage=DefaultFrontPate
	//建立服务器上相册文件夹
	album.Numbers=0
	err=Db.Create(album).Error
	if err!=nil{
		return err
	}
	path:="./albums/"+album.Name
	err=os.Mkdir(path,0755)
	return err
}

func DeleteAlbum(album *Album)(err error){
	//删除数据库中相册内所有照片
	temp:=Photo{AID:album.ID}
	err=Db.Unscoped().Delete(temp).Error
	if err!=nil{
		return err
	}
	//删除相册文件夹
	path:="./albums/"+album.Name
	err = os.RemoveAll(path)
	if err!=nil{
		return err
	}
	//删除相册
	err=Db.Unscoped().Delete(album).Error
	return err
}

func GetAlbumName(id string)(name string,err error){
	var aid int64
	aid,err=strconv.ParseInt(id,10,64)
	if err!=nil{
		name=""
		return
	}
	var temp=Album{ID:aid}
	err=Db.Find(&temp).Error
	if err!=nil{
		name=""
		return
	}
	name=temp.Name
	return
}

func AddPhoto(photo *Photo)(error){
	//新建图片
	err:=Db.Create(photo).Error
	if err!=nil{
		return err
	}
	//增加当前相册图片计数
	return Db.Model(Album{}).Where("id=?",photo.AID).Select("numbers").Update(map[string]interface{}{"numbers": gorm.Expr("numbers + ?", 1)}).Error
}

func GetAllPhotos(aid int64)([]*Photo,error){
	totals:=make([]*Photo,0)
	err:=Db.Where("a_id=?",aid).Find(&totals).Error
	return totals,err

}

func DeletePhoto(photo *Photo)(error){
	//减少当前相册图片计数
	err:=Db.Model(Album{}).Where("id=?",photo.AID).Select("numbers").Update(map[string]interface{}{"numbers": gorm.Expr("numbers - ?", 1)}).Error
	if err!=nil{
		return err
	}
	//删除图片
	return Db.Unscoped().Delete(photo).Error
}

func GetPhoto(photo *Photo)(error){
	return Db.First(photo).Error
}
package controller

import (
	"github.com/src/mylog"
	"github.com/gin-gonic/gin"
	"github.com/src/models"
	"net/http"
	"os"
	"strconv"
)

func AlbumGet(c *gin.Context){
	c.HTML(http.StatusOK,"album.html",gin.H{
		"IsAlbum":true,
		"IsLogin":checkAccount(c),
	})
}

func AlbumGetAll(c *gin.Context){
	//获得所有相册
	albums,err:=models.GetAllAlbums()
	islog:=checkAccount(c)
	//传回
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,
			"msg":"success",
			"nums":len(albums),
			"data":albums,
			"islog":islog,
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
			"nums":0,
			"data":nil,
			"islog":islog,
		})
	}
}

func AlbumPostNew(c *gin.Context){
	var album models.Album
	if !checkAccount(c){
		c.JSON(http.StatusOK,gin.H{
			"code":235,
			"msg":"非法的用户操作",
		})
		mylog.Logger.Println("A visitor attempted to add an album!")
		return
	}
	c.BindJSON(&album)
	err:=models.CreateAlbum(&album)
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,
			"msg":"success",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
		})
	}
}

func AlbumDelete(c *gin.Context){
	if !checkAccount(c){
		c.JSON(http.StatusOK,gin.H{
			"code":235,
			"msg":"非法的用户操作",
		})
		mylog.Logger.Println("A visitor attempted to delete an album!")
		return
	}
	var album models.Album
	c.BindJSON(&album)
	err:=models.DeleteAlbum(&album)
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,
			"msg":"success",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
		})
	}
}

func AlbumPhotoGet(c *gin.Context){
	c.HTML(http.StatusOK,"photo.html",gin.H{
		"IsAlbum":true,
		"IsLogin":checkAccount(c),
	})
}

func PhotoPostNew(c *gin.Context){
	if !checkAccount(c){
		c.JSON(http.StatusOK,gin.H{
			"code":235,
			"msg":"非法的用户操作",
		})
		mylog.Logger.Println("A visitor attempted to add a photo!")
		return
	}
	//获取文件
	file,err:=c.FormFile("pic")
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
		})
		return
	}
	//获取相册id
	albumid:=c.PostForm("albumid")
	//存图片文件
	//根据id返回相册名
	albumname,err:=models.GetAlbumName(albumid)
	path:="./albums/"+albumname+"/"+file.Filename
	err=c.SaveUploadedFile(file,path)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
		})
		mylog.Logger.Printf("%v\n",err)
		return
	}
	//写数据库
	aid,_:=strconv.ParseInt(albumid,10,64)
	photo:=models.Photo{Name:file.Filename,Path:"/images/"+albumname+"/"+file.Filename,AID:aid}
	err=models.AddPhoto(&photo)
	//返回
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,
			"msg":"success",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
		})
	}
}

func PhotoGetAll(c *gin.Context){
	var album models.Album
	c.BindJSON(&album)
	paths,err:=models.GetAllPhotos(album.ID)
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,
			"msg":"success",
			"nums":len(paths),
			"data":paths,
			"islog":checkAccount(c),
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
			"nums":0,
			"data":nil,
			"islog":checkAccount(c),
		})
	}
}

func PhotoDelete(c *gin.Context){
	if !checkAccount(c){
		c.JSON(http.StatusOK,gin.H{
			"code":235,
			"msg":"非法的用户操作",
		})
		mylog.Logger.Println("A visitor attempted to delete a photo!")
		return
	}
	var photo models.Photo
	c.BindJSON(&photo)
	err:=models.GetPhoto(&photo)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
		})
		return
	}
	err=models.DeletePhoto(&photo)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
		})
		return
	}
	//删除文件
	aname,err:=models.GetAlbumName(strconv.FormatInt(photo.AID,10))
	path:="./albums/"+aname+"/"+photo.Name
	err=os.Remove(path)
	if err==nil{
		c.JSON(http.StatusOK,gin.H{
			"code":666,
			"msg":"success",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":233,
			"msg":err.Error(),
		})
	}
}
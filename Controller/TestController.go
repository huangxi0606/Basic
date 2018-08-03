package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Basic/Db"
)


func Hhx(ctx *gin.Context){
   	var hhz= Db.Link()
   	if hhz == false{
		ctx.JSON(http.StatusSeeOther,gin.H{
			"code":404,
			"message":"数据库连接错误",
		})
		return
	}


}
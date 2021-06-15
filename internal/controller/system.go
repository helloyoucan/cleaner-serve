package controller

import (
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func AdminLogin(c *gin.Context) {

}

func Upload(c *gin.Context)  {
	file, err  := c.FormFile("file")
	if err != nil {
		util.RespJSON(c,gin.H{"err":err.Error()})
		return
	}
	log.Println(file.Filename)
	newFileName := strconv.FormatInt(time.Now().Unix(),10) + strconv.Itoa(rand.Intn(999999-100000)+10000)+"-" + file.Filename
	err=c.SaveUploadedFile(file, "./upload/"+newFileName)
	if err != nil {
		util.RespJSON(c,gin.H{"err":err.Error()})
		return
	}
	util.RespJSON(c,gin.H{"data":"/upload/"+newFileName})
}
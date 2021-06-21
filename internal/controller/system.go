package controller

import (
	"cleaner-serve/internal/logic"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {

}

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		util.RespJSON(c, gin.H{"err": err.Error()})
		return
	}
	url, err := logic.UploadFileToQiNiu(file)
	if err != nil {
		util.RespJSON(c, gin.H{"err": err.Error()})
		return
	}
	util.RespJSON(c, gin.H{"data": url})
}

func GetLocation(c *gin.Context) {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	data, err := logic.GetLocationInfoByBaiduMap(reqIP)
	if err != nil {
		util.RespJSON(c, gin.H{"err": err.Error()})
		return
	}
	util.RespJSON(c, gin.H{"data": data,})
}

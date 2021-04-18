package util

import (
	"cleaner-serve/internal/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespJSON(c *gin.Context, json gin.H) {
	var code = json["code"]
	if json["err"] != nil && code == nil {
		code = errcode.DBError
	}
	if code == nil {
		code = errcode.Success
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": json["data"],
		"msg":  json["msg"],
		"err":  json["err"],
	})
}

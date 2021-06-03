package util

import (
	"cleaner-serve/internal/errcode"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"io"
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

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

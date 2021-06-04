package util

import (
	"cleaner-serve/internal/errcode"
	"cleaner-serve/internal/models"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
)

// 统一返回请求的数据格式
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

// 处理获取的分页参数
func Paginate(pages *models.Pages) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pages.Page == 0 {
			pages.Page = 1
		}
		switch {
		case pages.PageSize > 100:
			pages.PageSize = 100
			break
		case pages.PageSize <= 0:
			pages.PageSize = 10
			break
		}
		offset := (pages.Page - 1) * pages.PageSize
		return db.Offset(offset).Limit(pages.PageSize)
	}
}

// 处理需要返回的分页参数数据
func HandlePages(pages *models.Pages, total int64) (err error) {
	pages.Total, err = strconv.Atoi(strconv.FormatInt(total, 10))
	if err != nil {
		return err
	}
	pages.TotalPage = pages.Total / pages.PageSize
	if pages.Total%pages.PageSize != 0 {
		pages.TotalPage += 1
	}
	return
}

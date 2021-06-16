package util

import (
	"cleaner-serve/internal/errcode"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	mathRand "math/rand"
	"net/http"
	"strconv"
	"time"
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

// int64转 int
func Int64ToInt(num int64) (intNum int) {
	intNum, err := strconv.Atoi(strconv.FormatInt(num, 10))
	if err != nil {
		fmt.Errorf("int64转 int=>" + err.Error())
		return 0
	}
	return
}
func StrToUint(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

// 计算总页数
func CalcTotalPage(total int, pageSize int) (totalPage int) {
	totalPage = total / pageSize
	if total%pageSize != 0 {
		totalPage += 1
	}
	return
}

// 处理获取的分页参数
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// 处理查询创建时间范围
func QueryCreated(startTime uint, endTime uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startTime != 0 {
			db.Where("created >= (?)", startTime)
		}
		if endTime != 0 {
			db.Where("created =< (?)", endTime)
		}
		return db
	}
}

func getTimeTick64() int64 {
	return time.Now().UnixNano() / 1e6
}

func getFormatTime(time time.Time) string {
	return time.Format("20060102")
}

/**
生成订单编号:日期20191025时间戳1571987125435+3位随机数
 */
func GenerateOrderNum()string {
	date := getFormatTime(time.Now())
	r := mathRand.Intn(100)
	return date+ strconv.FormatInt(getTimeTick64(),10)+ strconv.Itoa(r)
}

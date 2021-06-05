package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

type APIUserCoupon struct {
	CouponId string `json:"coupon_id"`
	UserId      string   `json:"user_id"`
}
func CreateAUserCoupon(c *gin.Context) {
	var apiUserCoupon APIUserCoupon
	c.BindJSON(&apiUserCoupon)
	if apiUserCoupon.UserId==""{
		util.RespJSON(c, gin.H{
			"err": "user_id 无效",
		})
		return
	}
	if apiUserCoupon.CouponId==""{
		util.RespJSON(c, gin.H{
			"err": "coupon_id 无效",
		})
		return
	}
	var userCoupon models.UserCoupon
	userCoupon.UserId = apiUserCoupon.UserId
	userCoupon.CouponId = apiUserCoupon.CouponId
	userCoupon.ID=util.UniqueId()
	coupon,err :=dao.GetACouponById(userCoupon.CouponId)
	if coupon.ID ==""{
		util.RespJSON(c, gin.H{
			"err": "优惠券不存在",
		})
		return
	}
	if err !=nil{
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	userCoupon.Name=coupon.Name
	userCoupon.StartTime=coupon.StartTime
	userCoupon.EndTime=coupon.EndTime
	userCoupon.Description=coupon.Description
	userCoupon.Status=8 //枚举
	err = dao.CreateAUserCoupon(&userCoupon)
	util.RespJSON(c, gin.H{
		"err": err,
	})

}
func GetUserCouponByUseId(c *gin.Context) {
	userId,isOk :=c.Params.Get("userId")
	if !isOk ||userId=="" {
		util.RespJSON(c, gin.H{
			"err": "user_id 无效",
		})
		return
	}
	// todo
	userCouponList, err := dao.GetUserCouponByUseId(userId)
	util.RespJSON(c, gin.H{
		"err":  err,
		"data": userCouponList,
	})
}

func UpdateAUserCoupon(c *gin.Context)  {
	id := c.Query("id")
	if id=="" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	userCoupon,err:=dao.GetAUserCouponById(id)
	if err!=nil {
		util.RespJSON(c,gin.H{
			"err":err.Error(),
		})
		return
	}
	err=c.BindJSON(&userCoupon)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err=dao.UpdateAUserCoupon(userCoupon)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{"data":userCoupon,})
}
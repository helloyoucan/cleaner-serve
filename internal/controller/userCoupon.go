package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

func CreateAUserCoupon(c *gin.Context) {
	var userCoupon models.UserCoupon
	c.BindJSON(&userCoupon)
	if userCoupon.UserId==""{
		util.RespJSON(c, gin.H{
			"err": "userId 无效",
		})
		return
	}
	if userCoupon.CouponId==""{
		util.RespJSON(c, gin.H{
			"err": "CouponId 无效",
		})
		return
	}
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
	userId := c.Query("userId")
	if userId=="" {
		util.RespJSON(c, gin.H{
			"err": "userId无效",
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
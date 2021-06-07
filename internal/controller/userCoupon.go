package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/logic"
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
	err:=logic.CreateAUserCoupon(apiUserCoupon.UserId,apiUserCoupon.CouponId)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{})

}
func GetUserCouponByUseId(c *gin.Context) {
	userId,isOk :=c.Params.Get("userId")
	if !isOk ||userId=="" {
		util.RespJSON(c, gin.H{
			"err": "user_id 无效",
		})
		return
	}
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
package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

func CreateACoupon(c *gin.Context) {
	var coupon models.Coupon
	err:=c.BindJSON(&coupon)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	coupon.ID=util.UniqueId()
	err = dao.CreateACoupon(&coupon)
	util.RespJSON(c, gin.H{
		"err": err,
	})
}
func GetAllCoupon(c *gin.Context) {
	couponList, err := dao.GetAllCoupon()
	util.RespJSON(c, gin.H{
		"err":  err,
		"data": couponList,
	})
}
func UpdateACoupon(c *gin.Context) {
	id := c.Query("id")
	if id=="" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	coupon, err := dao.GetACouponById(id)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err=c.BindJSON(&coupon)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err = dao.UpdateACoupon(coupon)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{})
}
func DeleteACoupon(c *gin.Context) {
	id := c.Query("id")
	if id=="" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	err := dao.DeleteACoupon(id)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{})
}

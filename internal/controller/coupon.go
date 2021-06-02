package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateACoupon(c *gin.Context) {
	var coupon models.Coupon
	c.BindJSON(&coupon)
	fmt.Println("---------------------"+coupon.Name)
	err := dao.CreateACoupon(&coupon)
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
	id, ok := c.Params.Get("id")
	if !ok {
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
	c.BindJSON(&coupon)
	err = dao.UpdateACoupon(coupon)
	util.RespJSON(c, gin.H{
		"err": err.Error(),
	})
}
func DeleteACoupon(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	err := dao.DeleteACoupon(id)
	util.RespJSON(c, gin.H{
		"err": err.Error(),
	})
}

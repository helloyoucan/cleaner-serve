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
			"err": err.Error(),
		})
		return
	}
	err = dao.CreateACoupon(&coupon)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{})
}
func GetCouponByPages(c *gin.Context) {
	var query =new(models.CouponQuery)
	err:=c.ShouldBind(&query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	couponList, total, err := dao.GetCouponByPages(query)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	var pages models.Pages
	pages.CalcPagesData(query.Page,query.PageSize,total)
	util.RespJSON(c, gin.H{
		"data": gin.H{
			"list":couponList,
			"pages":pages,
		},
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

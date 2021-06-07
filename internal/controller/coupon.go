package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
	"strconv"
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
	var pages =new (models.Pages)
	pages.Page,_ = strconv.Atoi(c.Query("page"))
	pages.PageSize,_ = strconv.Atoi(c.Query("page_size"))
	couponList, err := dao.GetCouponByPages(pages)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
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

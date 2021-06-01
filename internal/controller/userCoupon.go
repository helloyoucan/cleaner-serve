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
	err := dao.CreateAUserCoupon(&userCoupon)
	util.RespJSON(c, gin.H{
		"err": err,
	})
}
func GetUserCouponByUseId(c *gin.Context) {
	userId, ok := c.Params.Get("userId")
	if !ok {
		util.RespJSON(c, gin.H{
			"err": "userId无效",
		})
		return
	}
	couponList, err := dao.GetUserCouponByUseId(userId)
	util.RespJSON(c, gin.H{
		"err":  err,
		"data": couponList,
	})
}

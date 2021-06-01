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

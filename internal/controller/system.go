package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminLogin(c *gin.Context) {

}

func CreateACoupon(c *gin.Context) {
	var coupon models.Coupon
	c.Bind(&coupon)
	err := dao.CreateACoupon(&coupon)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":0,
			"data":nil,
			"msg":nil,
		})
	} else {

	}
}
func GetAllCoupon(c *gin.Context) {

}
func UpdateACoupon(c *gin.Context) {

}
func DeleteACoupon(c *gin.Context) {

}

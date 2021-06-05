package controller

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

type  APIOrder struct {
	models.Order
	ExtraService struct{
		ID string `json:"id"`
		models.BaseExtraService
	} `json:"extra_service"`
}
func CreateAOrder(c *gin.Context)  {
	var apiOrder APIOrder
	err:=c.BindJSON(apiOrder)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	// todo
	// 1.创建订单使用的额外服务
	// 2.创建订单
}
package controller

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

type ExtraService struct {
	ID string `json:"id"`
	models.BaseExtraService
}
type OrderCoupon struct {
	ID string `json:"id"`
	models.BaseOrderCoupon
}
type  APIOrder struct {
	ExtraServiceList []*ExtraService `json:"extra_service_list" gorm:"default:[]"`
	OrderCouponList []*OrderCoupon `json:"order_coupon_list" gorm:"default:[]"`
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
	// 1.创建订单使用的额外服务(多个)
	// todo
	//var extraService = new(models.ExtraService) //需要创建切片
	//extraService.ID=util.UniqueId()
	//extraService.Name = apiOrder.ExtraService.Name
	//extraService.UnitPrice = apiOrder.ExtraService.UnitPrice
	//extraService.Discount = apiOrder.ExtraService.Discount
	//extraService.Description = apiOrder.ExtraService.Description
	// 2.创建订单使用的优惠券(多个)
	// 2.创建订单
}
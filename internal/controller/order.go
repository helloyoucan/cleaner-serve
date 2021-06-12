package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

//type APIOrder struct {
//	Order            models.Order                `json:"order"`
//	ExtraServiceList []*models.OrderExtraService `json:"extra_service_list" gorm:"default:[]"`
//	OrderCouponList  []*models.OrderCoupon       `json:"order_coupon_list" gorm:"default:[]"`
//}

func CreateAOrder(c *gin.Context) {
	var order models.Order
	err := c.BindJSON(&order)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	// 创建订单
	err = dao.CreateAOrder(&order)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	//err = logic.CreateExtraServiceList(apiOrder.Order.ID, apiOrder.ExtraServiceList)
	//if err != nil {
	//	util.RespJSON(c, gin.H{
	//		"err": err,
	//	})
	//	return
	//}
	//err = logic.CreateOrderCouponList(apiOrder.Order.ID, apiOrder.OrderCouponList)
	//if err != nil {
	//	util.RespJSON(c, gin.H{
	//		"err": err,
	//	})
	//	return
	//}
	util.RespJSON(c, gin.H{})
}

func GetOrderByPages(c *gin.Context) {
	var query =new(models.OrderQuery)
	err:=c.ShouldBind(&query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	orderList, total, err := dao.GetOrderByPages(query)
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
			"list":  orderList,
			"pages": pages,
		},
	})
}
func GetOrderByUserByPages(c *gin.Context) {
	var query =new(models.OrderQuery)
	err:=c.ShouldBind(&query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	if query.UserId == "" {
		util.RespJSON(c, gin.H{
			"err": "userId 无效",
		})
		return
	}
	orderList, total,err := dao.GetOrderByUserByPages(query)
	var pages models.Pages
	pages.CalcPagesData(query.Page,query.PageSize,total)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{
		"data": gin.H{
			"list":  orderList,
			"pages": pages,
		},
	})
}

func GetAOrder(c *gin.Context) {
	orderId, isOk := c.Params.Get("orderId")
	if !isOk || orderId == "" {
		util.RespJSON(c, gin.H{
			"err": "orderId 无效",
		})
		return
	}
	order, err := dao.GetAOrderById(orderId)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{
		"data": order,
	})
}
func UpdateAOrder(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	order, err := dao.GetAOrderById(id)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err = c.BindJSON(&order)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err = dao.UpdateAOrder(order)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{"data": order,})
}
func DeleteAOrder(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	err := dao.DeleteAOrder(id)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{})
}

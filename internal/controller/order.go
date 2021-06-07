package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/logic"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type APIOrder struct {
	Order            models.Order                `json:"order"`
	ExtraServiceList []*models.OrderExtraService `json:"extra_service_list" gorm:"default:[]"`
	OrderCouponList  []*models.OrderCoupon       `json:"order_coupon_list" gorm:"default:[]"`
}

func CreateAOrder(c *gin.Context) {
	var apiOrder APIOrder
	err := c.BindJSON(&apiOrder)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	// 创建订单
	apiOrder.Order.ID = util.UniqueId()
	err = dao.CreateAOrder(&apiOrder.Order)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	err = logic.CreateExtraServiceList(apiOrder.Order.ID, apiOrder.ExtraServiceList)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	// 创建订单使用的附加服务
	err = logic.CreateOrderCouponList(apiOrder.Order.ID, apiOrder.OrderCouponList)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	util.RespJSON(c, gin.H{})
}

func GetOrderByPages(c *gin.Context) {
	var pages = new(models.Pages)
	pages.Page, _ = strconv.Atoi(c.Query("page"))
	pages.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	orderList, err := dao.GetOrderByPages(pages)
	util.RespJSON(c, gin.H{
		"err": err,
		"data": gin.H{
			"list":  orderList,
			"pages": pages,
		},
	})
}
func GetOrderByUserByPages(c *gin.Context) {
	var pages = new(models.Pages)
	pages.Page, _ = strconv.Atoi(c.Query("page"))
	pages.PageSize, _ = strconv.Atoi(c.Query("page_size"))
	userId, isOk := c.Params.Get("userId")
	if !isOk || userId == "" {
		util.RespJSON(c, gin.H{
			"err": "userId 无效",
		})
		return
	}
	orderList, err := dao.GetOrderByUserByPages(pages, userId)
	util.RespJSON(c, gin.H{
		"err": err,
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

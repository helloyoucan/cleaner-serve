package logic

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
)

func CreateOrderCouponList(orderId string,orderCouponList []*models.OrderCoupon)(err error)  {
	for _, orderCoupon := range orderCouponList {
		orderCoupon.ID=util.UniqueId()
		orderCoupon.OrderId=orderId
	}
	return dao.CreateOrderCouponList(orderCouponList)
}
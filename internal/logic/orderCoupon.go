package logic

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
)

func CreateOrderCouponList(orderId string,orderCouponList []*models.OrderCoupon)(err error)  {
	for _, orderCoupon := range orderCouponList {
		orderCoupon.OrderID=orderId
	}
	return dao.CreateOrderCouponList(orderCouponList)
}
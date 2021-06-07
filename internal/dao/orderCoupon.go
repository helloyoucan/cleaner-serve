package dao

import "cleaner-serve/internal/models"

/**
订单使用的优惠券
*/
func CreateOrderCouponList(orderCouponList []*models.OrderCoupon) (err error) {
	return DB.Create(&orderCouponList).Error
}
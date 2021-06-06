package dao

import "cleaner-serve/internal/models"

/**
订单使用的优惠券
*/
func CreateAOrderExtraService(coupon *models.Coupon) (err error) {
	return DB.Create(&coupon).Error
}
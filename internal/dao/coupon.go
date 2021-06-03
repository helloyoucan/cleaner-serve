package dao

import (
	"cleaner-serve/internal/models"
)

/**
优惠券的操作
*/
func CreateACoupon(coupon *models.Coupon) (err error) {
	return DB.Create(&coupon).Error
}
func GetAllCoupon() (couponList []*models.Coupon, err error) {
	err= DB.Find(&couponList).Error
	if err != nil {
		return nil, err
	}
	return
}

// 根据多个id获取优惠券
func GetAllCouponByCouponIds(ids []uint) (couponList []*models.Coupon, err error) {
	if len(ids) == 0 {
		return couponList, nil
	}
	var db = DB.Where("id = ?", ids[0])
	for _, v := range ids[1:] {
		db = db.Or("id = ?", v)
	}
	err = db.Find(&couponList).Error
	return
}
func GetACouponById(id string) (coupon *models.Coupon, err error) {
	coupon = new(models.Coupon) //不通过new关键字实例化就会报错
	err = DB.Where("id = ?", id).Find(&coupon).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateACoupon(coupon *models.Coupon) (err error) {
	err = DB.Save(&coupon).Error
	return
}
func DeleteACoupon(id string) (err error) {
	err = DB.Where("id=?", id).Delete(&models.Coupon{}).Error
	return
}

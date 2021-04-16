package dao

import "cleaner-serve/internal/models"



func CreateAUserCoupon(userCoupon *models.UserCoupon) (err error) {
	return DB.Create(&userCoupon).Error
}
// 通过用户ID获取多个优惠券
func GetUserCouponByUseId(userId string) (couponList []*models.Coupon, err error) {
	var userCouponList []*models.UserCoupon
	err = DB.Find(&userCouponList).Error
	if err != nil {
		return nil, err
	}
	var ids []uint
	for _, v := range userCouponList {
		ids = append(ids, v.ID)
	}
	couponList, err = getAllCouponByCouponIds(ids)
	if err != nil {
		return nil, err
	}
	return

}

func UpdateAUserCoupon(userCoupon *models.UserCoupon)(err error)  {
	err=DB.Save(&userCoupon).Error
	return
}

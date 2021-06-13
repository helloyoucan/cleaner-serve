package dao

import "cleaner-serve/internal/models"

// 用户领取优惠券
func CreateAUserCoupon(userCoupon *models.UserCoupon) (err error) {
	return DB.Create(&userCoupon).Error
}

// 通过用户ID获取多个优惠券
func GetUserCouponByUseId(userId string) (userCouponList []*models.UserCoupon, err error) {
	err = DB.Where("user_id=?",userId).Order("created desc").Find(&userCouponList).Error
	if err != nil {
		return nil, err
	}
	return

}
func GetAUserCouponById(id string) (userCoupon *models.UserCoupon, err error) {
	userCoupon = new(models.UserCoupon)
	err = DB.Where("id=?",id).Find(&userCoupon).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateAUserCoupon(userCoupon *models.UserCoupon) (err error) {
	err = DB.Save(&userCoupon).Error
	return
}

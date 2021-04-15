package dao

import "cleaner-serve/internal/models"

func CreateACoupon(coupon *models.Coupon) (err error){
	return DB.Create(&coupon).Error
}
func GetAllCoupon()(couponList []*models.Coupon,err error)  {
	err = DB.Find(&couponList).Error
	if  err != nil {
		return nil, err
	}
	return 
}
func GetUserCouponByUseId(userId string)(couponList []*models.Coupon,err error)  {
	var userCouponList []*models.UserCoupon
	err =DB.Find(&userCouponList).Error
	if  err != nil {
		return nil, err
	}
	couponList,err = getCouponsByCouponIds(userCouponList)
	if  err != nil {
		return nil, err
	}
	return

}
func getCouponsByCouponIds(userCouponList []*models.UserCoupon)(couponList []*models.Coupon,err error)  {
	return
}
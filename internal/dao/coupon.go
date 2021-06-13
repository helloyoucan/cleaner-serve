package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"errors"
)

/**
优惠券的操作
*/
func CreateACoupon(coupon *models.Coupon) (err error) {
	return DB.Create(&coupon).Error
}

func GetCouponByPages(query *models.CouponQuery) (couponList []*models.Coupon, total int64, err error) {
	db := DB.Scopes(util.QueryCreated(query.CreatedStartTime, query.CreatedEndTime)).Order("created desc")
	if query.Name != "" {
		db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.StartTime != 0 {
		db.Where("start_time >= (?)", query.StartTime)
	}
	if query.EndTime != 0 {
		db.Where("end_time =< (?)", query.EndTime)
	}
	db.Model(&models.Coupon{}).Count(&total)
	err = db.Scopes(util.Paginate(query.Page, query.PageSize)).Find(&couponList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

//// 根据多个id获取优惠券
//func GetAllCouponByCouponIds(ids []string) (couponList []*models.Coupon, err error) {
//	if len(ids) == 0 {
//		return couponList, nil
//	}
//	var db = DB.Where("id = ?", ids[0])
//	for _, v := range ids[1:] {
//		db = db.Or("id = ?", v)
//	}
//	err = db.Find(&couponList).Error
//	return
//}
func GetACouponById(id string) (coupon *models.Coupon, err error) {
	coupon = new(models.Coupon) //不通过new关键字实例化就会报错
	err = DB.Where("id = ?", id).Find(&coupon).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateACoupon(coupon *models.Coupon) (err error) {
	result := DB.Where("id=?", coupon.ID).Select("*").Updates(&coupon)
	if result.RowsAffected == 0 {
		return errors.New("数据不存在")
	}
	return result.Error
}
func DeleteACoupon(id string) (err error) {
	err = DB.Where("id=?", id).Delete(&models.Coupon{}).Error
	return
}

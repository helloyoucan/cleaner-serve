package logic

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"errors"
)

func CreateAUserCoupon(userId string,couponId string)(err error)  {
	var userCoupon models.UserCoupon
	userCoupon.UserId = userId
	userCoupon.CouponId = couponId
	userCoupon.ID=util.UniqueId()
	coupon,err :=dao.GetACouponById(userCoupon.CouponId)
	if err !=nil{
		return err
	}
	if coupon.ID ==""{
		return errors.New("优惠券不存在")
	}
	userCoupon.Name=coupon.Name
	userCoupon.StartTime=coupon.StartTime
	userCoupon.EndTime=coupon.EndTime
	userCoupon.Description=coupon.Description
	userCoupon.Status=8 //枚举
	err = dao.CreateAUserCoupon(&userCoupon)
	return  err
}
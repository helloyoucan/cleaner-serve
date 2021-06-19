package logic

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/enum"
	"cleaner-serve/internal/models"
	"errors"
)

func CreateAUserCoupon(userId string,couponId string)(err error)  {
	var userCoupon models.UserCoupon
	userCoupon.UserId = userId
	userCoupon.CouponId = couponId
	coupon,err :=dao.GetACouponById(userCoupon.CouponId)
	if err !=nil{
		return err
	}
	if coupon.ID ==""{
		return errors.New("优惠券不存在")
	}
	userCoupon.Name=coupon.Name
	userCoupon.Type = coupon.Type
	userCoupon.TypeValue = coupon.TypeValue
	userCoupon.ThresholdType = coupon.ThresholdType
	userCoupon.ThresholdValue = coupon.ThresholdValue
	userCoupon.ExpiryType = coupon.ExpiryType
	userCoupon.ExpiryTypeValue = coupon.ExpiryTypeValue
	userCoupon.StartTime=coupon.StartTime
	userCoupon.EndTime=coupon.EndTime
	userCoupon.Description=coupon.Description
	userCoupon.Status=enum.UserCouponStatusBeUse
	err = dao.CreateAUserCoupon(&userCoupon)
	return  err
}
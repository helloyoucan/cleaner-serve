package enum

const (
	BranchStatusClose=0 //关闭
	BranchStatusOpen=1//营业中
	BranchStatusRest=2//休息中

	//优惠券门槛
	CouponThresholdTypeNone=0
	CouponThresholdTypeFixedAmount=1
	CouponThresholdTypeFirstOrder=2

	// 用户领取的优惠券
	UserCouponStatusBeUse = 0 //未使用
	UserCouponStatusUsed = 1 //已使用
)
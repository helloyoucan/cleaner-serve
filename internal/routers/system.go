package routers

import (
	"cleaner-serve/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter(r *gin.Engine) {
	systemGroup := r.Group("api/system")
	systemGroup.GET("/login", controller.AdminLogin)
	// 优惠券
	systemGroup.POST("/coupon", controller.CreateACoupon)
	systemGroup.GET("/coupon", controller.GetCouponByPages)
	systemGroup.PUT("/coupon", controller.UpdateACoupon)
	systemGroup.DELETE("/coupon", controller.DeleteACoupon)
	//服务网店
	systemGroup.POST("/branch", controller.CreateABranch)
	systemGroup.GET("/branch", controller.GetBranchByPages)
	systemGroup.PUT("/branch", controller.UpdateABranch)
	systemGroup.DELETE("/branch", controller.DeleteABranch)
	// 用户领取的优惠券
	systemGroup.POST("/user/coupon", controller.CreateAUserCoupon)
	systemGroup.GET("/user/coupon", controller.GetUserCouponByUseId)
	systemGroup.PUT("/user/coupon", controller.UpdateAUserCoupon)
}

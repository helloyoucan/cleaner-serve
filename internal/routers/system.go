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
	systemGroup.GET("/coupon/pages", controller.GetCouponByPages)
	systemGroup.PUT("/coupon", controller.UpdateACoupon)
	systemGroup.DELETE("/coupon", controller.DeleteACoupon)
	//服务网店
	systemGroup.POST("/branch", controller.CreateABranch)
	systemGroup.GET("/branch/pages", controller.GetBranchByPages)
	systemGroup.PUT("/branch", controller.UpdateABranch)
	systemGroup.DELETE("/branch", controller.DeleteABranch)
	// 用户领取的优惠券
	systemGroup.POST("/user/coupon", controller.CreateAUserCoupon)
	systemGroup.GET("/user/coupon/:userId", controller.GetUserCouponByUseId)
	systemGroup.PUT("/user/coupon", controller.UpdateAUserCoupon)
	// 附加服务
	systemGroup.POST("/extraService", controller.CreateAExtraService)
	systemGroup.GET("/extraService/pages", controller.GetExtraServiceByPages)
	systemGroup.GET("/extraService/active", controller.GetAllActiveExtraService)
	systemGroup.PUT("/extraService", controller.UpdateAExtraService)
	systemGroup.DELETE("/extraService", controller.DeleteAExtraService)
	// 订单
	systemGroup.POST("/order", controller.CreateAOrder)
}

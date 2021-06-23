package routers

import (
	"cleaner-serve/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitWxRouter(r *gin.Engine) {
	wxGroup := r.Group("wx")
	wxGroup.GET("/session", controller.GetUserInfo)

	//服务网点
	wxGroup.GET("/branch", controller.GetAllBranch)
	// 领取优惠券
	wxGroup.POST("/user/coupon", controller.CreateAUserCoupon)
	// 附加服务
	wxGroup.GET("/extraService", controller.GetAllActiveExtraService)
	// 订单
	wxGroup.POST("/order", controller.CreateAOrder)
	wxGroup.GET("/order/:userId/pages", controller.GetOrderByUserByPages)
	wxGroup.GET("/order/user/:orderId", controller.GetAOrder)
	wxGroup.PUT("/order", controller.UpdateAOrder)
}

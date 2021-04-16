package routers

import (
	"cleaner-serve/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitAdminRouter(r *gin.Engine) {
	systemGroup := r.Group("api/system")
	// 优惠券
	systemGroup.GET("/login", controller.AdminLogin)
	systemGroup.POST("/coupon", controller.CreateACoupon)
	systemGroup.GET("/coupon", controller.GetAllCoupon)
	systemGroup.PUT("/coupon", controller.UpdateACoupon)
	systemGroup.DELETE("/coupon", controller.DeleteACoupon)
}

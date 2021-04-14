package routers

import (
	"cleaner-serve/internal/controller"
	"github.com/gin-gonic/gin"
)

func InitWxRouter(r *gin.Engine) {
	wxGroup := r.Group("wx")
	wxGroup.GET("/session", controller.GetUserInfo)
}

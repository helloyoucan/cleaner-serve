package routers

import "github.com/gin-gonic/gin"

func SetupRouter() (r *gin.Engine) {
	r = gin.Default()
	InitWxRouter(r)
	InitAdminRouter(r)
	return
}

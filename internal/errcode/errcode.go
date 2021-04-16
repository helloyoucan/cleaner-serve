package errcode

import "github.com/gin-gonic/gin"

var (
	Success = gin.H{
		"code":    0,
		"msg": nil,
	}
)

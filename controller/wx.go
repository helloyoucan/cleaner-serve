package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const APPID = "wxcd39857e98e1162d"
const SECRET = "b5e7be2013df3fe008aa71f65bff00d2"

func GetUserInfo(c *gin.Context) {
	wxCode := c.Param("code")
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + wxCode + "&grant_type=authorization_code")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", resp)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

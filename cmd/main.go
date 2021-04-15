package main

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/routers"
	"fmt"
)

func main() {
	err:=dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer  dao.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Coupon{})
	r := routers.SetupRouter()
	r.Run(":8080")
	fmt.Println("serve run 8080")
}

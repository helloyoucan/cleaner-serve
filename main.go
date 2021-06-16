package main

import (
	"cleaner-serve/configs"
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/routers"
	"fmt"
)


func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	// 模型绑定
	dao.DB.AutoMigrate(&models.Coupon{})
	dao.DB.AutoMigrate(&models.UserCoupon{})
	dao.DB.AutoMigrate(&models.OrderCoupon{})
	dao.DB.AutoMigrate(&models.ExtraService{})
	dao.DB.AutoMigrate(&models.OrderExtraService{})
	dao.DB.AutoMigrate(&models.Order{})
	dao.DB.AutoMigrate(&models.Warrior{})
	dao.DB.AutoMigrate(&models.Branch{})
	r := routers.SetupRouter()
	r.Static("/upload", "./upload")
	//r.StaticFS("/more_static", http.Dir("my_file_system"))
	//r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	r.Run(":"+configs.Port)
	fmt.Println("serve run 8080")
}

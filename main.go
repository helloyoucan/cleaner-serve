package main

import (
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
	defer dao.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Coupon{})
	dao.DB.AutoMigrate(&models.UserCoupon{})
	dao.DB.AutoMigrate(&models.ClientInfo{})
	dao.DB.AutoMigrate(&models.Branch{})
	dao.DB.AutoMigrate(&models.ExtraService{})
	dao.DB.AutoMigrate(&models.WarriorInfo{})
	dao.DB.AutoMigrate(&models.Order{})
	r := routers.SetupRouter()
	r.Run(":8080")
	fmt.Println("serve run 8080")
}

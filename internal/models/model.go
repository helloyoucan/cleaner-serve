package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 洗衣机信息
type MachineInfo struct {
	gorm.Model
	Brand string `json:"brand"`
	Type string `json:type`
	Mode string `json:"mode"`
	Photos []string `json:"photo"`
	Remark string `json:"remark"`
}

// 优惠券
type Coupon struct {
	gorm.Model
	Name string `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Status uint8 `json:"status"`
}
// 订单
type Order struct {
	gorm.Model
	Status uint8 `json:"status"`
	ClientInfo *ClientInfo `json:"client_info"`
	Branch *Branch `json:"branch"`
	BaseCost float32 `json:"base_cost"`
	ExtraService *ExtraService `json:"extra_service"`
	ExtraDistance *ExtraDistance `json:"extra_distance"`
	CouponId uint64 `json:"coupon_id"`
	TotalPrice float64 `json:"total_price"`
}
// 附加服务
type ExtraService struct {
	gorm.Model
	Name string `json:"name"`
	Cost float32 `json:"cost"`
	Discount float32 `json:"discount"`
}
// 额外的距离
type ExtraDistance struct {
	UnitPrice float32 `json:"unit_price"`
	distance float64 `json:"distance"`
}
// 订单的客户信息
type ClientInfo struct {
	Name string `json:"name"`
	Phone uint16 `json:"phone"`
	Address struct {
		Province string `json:"province"`
		City string `json:"city"`
		Area string `json:"area"`
		Detail string `json:"detail"`
	}`json:"address"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
}
// 服务网店
type Branch struct {
	gorm.Model
	Name string `json："name"`
	Latitude int32 `json:"latitude"`
	Longitude int32 `json:"longitude"`
	ContactPerson string `json:"contact_person"`
	ContactPhone uint16 `json:"contact_phone"`
}
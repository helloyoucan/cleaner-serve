package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BaseCoupon struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `json:"name"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Description string `json:"description"`
}

// 优惠券
type Coupon struct {
	BaseCoupon
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// 用户拥有的优惠券
type UserCoupon struct {
	gorm.Model
	CouponId uint  `json:"coupon_id"`
	UserId   uint  `json:"user_id"`
	Status   uint8 `json:"status"`
}

//机器信息
type Machine struct {
	gorm.Model
	Brand  string `json:"brand"` // 品牌
	Type   string `json:"type"`
	Mode   string `json:"mode"`
	Photos string `json:"photo" gorm:"default:'[]'"`
	Remark string `json:"remark"  gorm:"default:''"`
}

// 订单
type Order struct {
	gorm.Model
	Branch       *Branch `json:"branch"` //服务网点
	ClientInfoId *uint   `json:"client_info_id"`
	Status       uint8   `json:"status"`
	StartTime    int64   `json:"start_time"`
	EndTime      int64   `json:"end_time"`
	Machine      Machine `json:"machine_info" gorm:"embedded;embeddedPrefix:machine_"`
	// ExtraServices      []baseExtraService `json:"extra_services" gorm:"default:[]ExtraServiceBase"`
	Distance          uint64 `json:"distance"`            //距离网点
	ExtraDistanceCost uint64 `json:"extra_distance_cost"` // 超出服务范围的费用
	// Coupons            []BaseCoupon `json:"coupons"`
	TotalPrice         float64 `json:"total_price"`           //总费用
	WarriorId          *uint   `json:"warrior_id"`            //接单战士id
	RefundStatus       uint8   `json:"refund_status"`         //退款状态
	RefundSArrivalTime int64   `json:"refund_s_arrival_time"` //退款到账时间
}

// 接单服务的战士信息
type WarriorInfo struct {
	gorm.Model
	Name           string  `json:"name"`
	Phone          string  `json:"phone"`
	Score          float32 `json:"score"` //评分
	Age            uint8   `json:"age"`
	Gender         uint8   `json:"gender"`
	Address        string  `json:"address"`
	JoinTime       int64   `json:"join_time" gorm:"autoCreateTime"`
	BelongBranchId *uint   `json:"belong_branch_id"`
}
type baseExtraService struct {
	Name        string  `json:"name"`
	Cost        float32 `json:"cost"`
	Discount    float32 `json:"discount"` //这个服务的折扣
	Description string  `json:"description"`
}

// 附加服务
type ExtraService struct {
	gorm.Model
	baseExtraService
}
type BaseClientInfo struct {
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    uint16 `json:"phone"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Address  string `json:"address"`
}

// 订单的客户信息
type ClientInfo struct {
	gorm.Model
	BaseClientInfo
}

// 服务网店
type Branch struct {
	gorm.Model
	Name                string  `json:"name"`
	Latitude            int32   `json:"latitude"`
	Longitude           int32   `json:"longitude"`
	ContactPerson       *string `json:"contact_person"`
	ContactPhone        *uint16 `json:"contact_phone"`
	WarriorManagerId    *uint   `json:"warrior_manager_id"`     // 管理这个店的战士
	Range               *uint   `json:"range"`                  //服务范围
	BaseCost            *uint64 `json:"base_cost"`              //这个网点的基础费用
	ExtraRangeUnitPrice *uint64 `json:"extra_range_unit_price"` //超出范围的单价
}

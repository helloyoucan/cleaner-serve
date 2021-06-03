package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 优惠券
type Coupon struct {
	ID        string `gorm:"primary_key" gorm:"<-:create"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name        string `json:"name"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Description string `json:"description"`
}

// 用户拥有的优惠券(用户关联)
type UserCoupon struct {
	gorm.Model
	UserId      uint   `json:"user_id"`
	Status      uint8  `json:"status"`
	Name        string `json:"name"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Description string `json:"description"`
}

// 订单使用的拥有的优惠券(订单表关联)
type OrderCoupon struct {
	gorm.Model
	UserCouponId uint   `json:"user_coupon_id"`
	OrderId      uint   `json:"order_id"`
	Name         string `json:"name"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Description  string `json:"description"`
}

//机器信息
type Machine struct {
	gorm.Model
	Brand         string `json:"brand"` // 品牌
	Type          string `json:"type"`
	Mode          string `json:"mode"`
	PhotosJsonStr string `json:"photos_json_str" gorm:"default:''"`
	Remark        string `json:"remark"  gorm:"default:''"`
}

// 附加服务
type ExtraService struct {
	gorm.Model
	Name        string  `json:"name"`
	UnitPrice   float32 `json:"unit_price"`
	Discount    float32 `json:"discount"` //这个服务的折扣
	Description string  `json:"description"`
}

// 订单使用的附加服务（订单关联）
type OrderExtraService struct {
	gorm.Model
	OrderId        uint    `json:"order_id"`
	ExtraServiceId uint    `json:"extra_service_id"`
	Name           string  `json:"name"`
	Cost           float32 `json:"cost"`
	Discount       float32 `json:"discount"` //这个服务的折扣
	Description    string  `json:"description"`
}

// 订单
type Order struct {
	gorm.Model
	BranchId           uint    `json:"branch_id"`      //服务网点
	ClientInfoId       uint    `json:"client_info_id"` //客户信息id
	Status             uint8   `json:"status"`
	StartTime          int64   `json:"start_time"`
	EndTime            int64   `json:"end_time"`
	MachineId          uint    `json:"machine_id"`
	Distance           uint64  `json:"distance"`              //距离网点
	TotalPrice         float64 `json:"total_price"`           //总费用
	WarriorId          *uint   `json:"warrior_id"`            //接单战士id
	RefundStatus       uint8   `json:"refund_status"`         //退款状态
	RefundSArrivalTime int64   `json:"refund_s_arrival_time"` //退款到账时间
}

// 接单服务的战士信息
type WarriorInfo struct {
	gorm.Model
	Name           string  `json:"name"`
	Phone          uint64  `json:"phone"`
	Score          float32 `json:"score"` //评分
	Age            uint8   `json:"age"`
	Gender         uint8   `json:"gender"`
	Address        string  `json:"address"`
	JoinTime       int64   `json:"join_time" gorm:"autoCreateTime"`
	BelongBranchId *uint   `json:"belong_branch_id"`
}

// 订单的客户信息（地址簿）
type ClientInfo struct {
	ID        string `gorm:"primary_key" gorm:"<-:create"`
	CreatedAt time.Time
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    uint64 `json:"phone"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Address  string `json:"address"`
}

// 服务网店
type Branch struct {
	ID        string `gorm:"primary_key" gorm:"<-:create"`
	CreatedAt time.Time
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Name                string  `json:"name"`
	Latitude            int32   `json:"latitude"`
	Longitude           int32   `json:"longitude"`
	ContactPerson       *string `json:"contact_person"` //联系人
	ContactPhone        *uint64  `json:"contact_phone"`  // 联系人电话
	WarriorManagerId    *uint   `json:"warrior_manager_id"`     // 管理这个店的战士
	Range               *uint   `json:"range"`                  //服务范围
	BaseCost            *uint64 `json:"base_cost"`              //这个网点的基础费用（单位：分）
	ExtraRangeUnitPrice *uint64 `json:"extra_range_unit_price"` //超出范围的单价（单位：分）
}

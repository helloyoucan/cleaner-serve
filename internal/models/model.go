package models

import (
	"gorm.io/plugin/soft_delete"
)
type Pages struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	Total int `json:"total"`
	TotalPage int `json:"total_page"`
}
// 服务网店
type BaseBranch struct {
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
type Branch struct {
	ID        string `json:"id" gorm:"primary_key;<-:create"`
	Created   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
	BaseBranch
}

// 优惠券
type Coupon struct {
	ID        string `json:"id" gorm:"primary_key;<-:create"`
	Created   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
	Name        string `json:"name"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Description string `json:"description"`
}

// 用户拥有的优惠券(用户关联，一个用户对多个优惠券)
type UserCoupon struct {
	ID        string `json:"id" gorm:"primary_key;<-:create"`
	Created   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
	CouponId string `json:"coupon_id"`
	UserId      string   `json:"user_id"`
	Status      uint8  `json:"status" gorm:"default:0"`
	Name        string `json:"name"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Description string `json:"description"`
}

// 接单服务的战士信息
type Warrior struct {
	ID        string `json:"id" gorm:"primary_key;<-:create"`
	Created   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
	Name           string  `json:"name"`
	IDCard uint `json:"id_card"`
	Phone          uint64  `json:"phone"`
	Score          float32 `json:"score"` //评分
	Age            uint8   `json:"age"`
	Gender         uint8   `json:"gender"`
	JoinTime       int64   `json:"join_time" gorm:"autoCreateTime"`
	BelongBranchId *uint   `json:"belong_branch_id"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Address  string `json:"address"`
	Status uint8 `json:"status" gorm:"default:0"` //账号状态
}

// 附加服务
type BaseExtraService struct {
	Name        string  `json:"name"`
	UnitPrice   int `json:"unit_price"`
	Discount    float32 `json:"discount"` //这个服务的折扣
	Description string  `json:"description"`
}
type ExtraService struct {
	ID        string `json:"id" gorm:"primary_key;<-:create"`
	Created   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
	IsActive int8 `json:"is_active" gorm:"default:0"` // 是否在可用状态,1为可用，0为不可用
	BaseExtraService
}
type BaseOrderCoupon struct {
	UserCouponId string   `json:"user_coupon_id"`
	OrderId      string   `json:"order_id"`
	Name         string `json:"name"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Description  string `json:"description"`
}
// 订单使用的拥有的优惠券(订单表关联，一个订单对应多张优惠券)
type OrderCoupon struct {
	ID        string `json:"id" gorm:"primary_key;<-:create"`
	Created   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
}
// 订单使用的附加服务（订单关联,一订单对应多个附加服务）
type OrderExtraService struct {
	ID        string `json:"id" gorm:"primary_key;<-:create"`
	Created   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
	OrderId        string    `json:"order_id"`
	ExtraServiceId string    `json:"extra_service_id"`
	BaseExtraService
}
// 订单
type Order struct {
	ID        string `json:"id" gorm:"primary_key;<-:create"`
	Created   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
	Status             uint8   `json:"status" gorm:"default:0"` //订单状态
	TotalPrice         float64 `json:"total_price"`           //总费用
	WarriorId          *uint   `json:"warrior_id"`            //接单战士id
	RefundStatus       uint8   `json:"refund_status" gorm:"default:0"`         //退款状态
	RefundSArrivalTime int64   `json:"refund_s_arrival_time"` //退款到账时间
	// 接单战士
	Warrior struct{
		ID string `json:"id"`
		Name           string  `json:"name"`
		Phone          uint64  `json:"phone"`
		BelongBranchId *uint   `json:"belong_branch_id"`
	} `json:"warrior" gorm:"embedded;embeddedPrefix:warrior_"`
	//服务网点信息
	Branch struct{
		ID string `json:"id"`
		BaseBranch
	} `json:"branch" gorm:"embedded;embeddedPrefix:branch_"`
	// 服务地址及联系人信息
	ClientInfo struct{
		UserId   uint   `json:"user_id"`
		Name     string `json:"name"`
		Phone    uint64 `json:"phone"`
		Province string `json:"province"`
		City     string `json:"city"`
		Area     string `json:"area"`
		Address  string `json:"address"`
		Distance uint64  `json:"distance"`   //距离网点的距离
		StartTime int64   `json:"start_time"` //预约的服务时间-开始时间
		EndTime  int64   `json:"end_time"` //预约的服务时间-结束时间
	} `json:"client_info" gorm:"embedded;embeddedPrefix:client_info_"`
	// 电器信息
	Machine struct{
		Brand         string `json:"brand"` // 品牌
		Type          string `json:"type"`
		Mode          string `json:"mode"`
		PhotosJsonStr string `json:"photos_json_str" gorm:"default:'[]'"`
		Remark        string `json:"remark" gorm:"default:''"`
	} `json:"machine" gorm:"embedded;embeddedPrefix:machine_"`

}

package models

import (
	"cleaner-serve/internal/util"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)
type Pages struct {
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	Total int `json:"total"`
	TotalPage int `json:"total_page"`
}
type BaseModel struct {
	ID        string `json:"id" gorm:"primary_key;<-:create"`
	Created   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
}
// 服务网店
type BaseBranch struct {
	Name                string  `json:"name"`
	Latitude            int32   `json:"latitude"`
	Longitude           int32   `json:"longitude"`
	ContactPerson       *string `json:"contact_person"` //联系人
	ContactPhone        *uint64  `json:"contact_phone"`  // 联系人电话
	WarriorManagerId    string   `json:"warrior_manager_id"`     // 管理这个店的战士
	Range               *uint   `json:"range"`                  //服务范围
	BaseCost            *uint64 `json:"base_cost"`              //这个网点的基础费用（单位：分）
	ExtraRangeUnitPrice *uint64 `json:"extra_range_unit_price"` //超出范围的单价（单位：分）
}
type Branch struct {
	BaseModel
	BaseBranch
}

// 优惠券
type Coupon struct {
	BaseModel
	Name        string `json:"name"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Description string `json:"description"`
}

// 用户拥有的优惠券(用户关联，一个用户对多个优惠券)
type UserCoupon struct {
	BaseModel
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
	BaseModel
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


// 订单使用的拥有的优惠券(订单表关联，一个订单对应多张优惠券)
type OrderCoupon struct {
	BaseModel
	UserCouponId string   `json:"user_coupon_id"`
	Name         string `json:"name"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Description  string `json:"description"`
	OrderID      string   `json:"order_id"` //外键
}
// 附加服务
type BaseExtraService struct {
	BaseModel
	Name        string  `json:"name"`
	UnitPrice   int `json:"unit_price"`
	Discount    float32 `json:"discount"` //这个服务的折扣
	Description string  `json:"description"`
}
type ExtraService struct {
	IsActive int8 `json:"is_active" gorm:"default:0"` // 是否在可用状态,1为可用，0为不可用
	BaseExtraService
}
// 订单使用的附加服务（订单关联,一订单对应多个附加服务）
type OrderExtraService struct {
	OrderID        string    `json:"order_id"`  //外键
	ExtraServiceId string    `json:"extra_service_id"`
	BaseExtraService
}
// 订单
type Order struct {
	BaseModel
	UserID string `json:"user_id"`
	Status             uint8   `json:"status" gorm:"default:0"` //订单状态
	TotalPrice         float64 `json:"total_price"`           //总费用
	RefundStatus       uint8   `json:"refund_status" gorm:"default:0"`         //退款状态
	RefundSArrivalTime *int64   `json:"refund_s_arrival_time"` //退款到账时间
	// 接单战士
	Warrior *struct{
		ID *string `json:"id"`
		Name           *string  `json:"name"`
		Phone          *uint64  `json:"phone"`
		BelongBranchId *uint   `json:"belong_branch_id"`
	} `json:"warrior" gorm:"embedded;embeddedPrefix:warrior_"`
	//服务网点信息
	Branch struct{
		ID string `json:"id"`
		BaseBranch
	} `json:"branch" gorm:"embedded;embeddedPrefix:branch_"`
	// 服务地址及联系人信息
	ClientInfo struct{
		UserId   string   `json:"user_id"`
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
	ExtraServices []OrderExtraService `json:"extra_services"`
	OrderCoupons []OrderCoupon `json:"order_coupons"`
}
func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = util.UniqueId()
	return
}

func (pages *Pages) CalcPages(total int64)  {
	if pages.Page == 0 {
		pages.Page = 1
	}
	switch {
	case pages.PageSize > 100:
		pages.PageSize = 100
		break
	case pages.PageSize <= 0:
		pages.PageSize = 10
		break
	}
	pages.Total= util.Int64ToInt(total)
	pages.TotalPage = util.CalcTotalPage(pages.Total,pages.PageSize)
}


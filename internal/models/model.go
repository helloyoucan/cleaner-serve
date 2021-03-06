package models

import (
	"cleaner-serve/internal/util"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Pages struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}
type BaseModel struct {
	ID        string                `json:"id" gorm:"primary_key;<-:create"`
	Created   int64                 `json:"-" gorm:"autoUpdateTime:milli"`
	Updated   int64                 `json:"-" gorm:"autoUpdateTime:milli"`
	DeletedAt soft_delete.DeletedAt `json:"-" `
}

// 服务网点
type BaseBranch struct {
	Name                string  `json:"name"`
	Longitude           string `json:"longitude"` //经度
	Latitude            string `json:"latitude"`  //纬度
	Province            string  `json:"province"`
	City                string  `json:"city"`
	Area                string  `json:"area"`
	Address             string  `json:"address"`
	ContactPerson       *string `json:"contact_person"`           //联系人
	ContactPhone        *uint64 `json:"contact_phone"`            // 联系人电话
	WarriorManagerId    *string `json:"warrior_manager_id"`       // 管理这个店的战士
	Range               *uint   `json:"range"`                    //服务范围（单位米）
	BaseCost            *uint64 `json:"base_cost"`                //这个网点的基础费用（单位：分）
	ExtraRangeUnitPrice *uint64 `json:"extra_range_unit_price"`   //超出范围的单价（单位：分）
	Status              uint8   `json:"status" gorm:"default:0"`  // 网点状态 0关闭，1营业中，2休息中
	Remark              string  `json:"remark" gorm:"default:''"` //备注
}
type Branch struct {
	BaseModel
	BaseBranch
	Created int64 `json:"created" gorm:"autoUpdateTime:milli"`
}
type BaseCoupon struct {
	Name        string `json:"name"`
	Type uint8 `json:"type"` // 优惠类型 0:指定金额，1:折扣
	TypeValue uint `json:"type_value"` // 优惠类型对应的值
	ThresholdType uint `json:"threshold_type"` // 使用门槛 0:无，1:指定金额，2:用户首单
	ThresholdValue *uint `json:"threshold_value"`//有使用门槛时对应的值
	ExpiryType uint `json:"expiry_type"`//有效期类型 0:固定日期,1:领取当日开始N天内有效
	ExpiryTypeValue *uint `json:"expiry_type_value"`//当ExpiryType=1时，绑定的值
	StartTime   *int64  `json:"start_time"`
	EndTime     *int64  `json:"end_time"`
	Description string `json:"description"`
}
// 优惠券
type Coupon struct {
	BaseModel
	BaseCoupon
	TotalAmount uint `json:"total_amount"`//可领取的总数量
	IssuedAmount uint `json:"issued_amount"`//还可领取数量
	Created     int64  `json:"created" gorm:"autoUpdateTime:milli"`
}

// 用户拥有的优惠券(用户关联，一个用户对多个优惠券)
type UserCoupon struct {
	BaseModel
	BaseCoupon
	CouponId    string `json:"coupon_id"`
	UserId      string `json:"user_id"`
	Status      uint8  `json:"status" gorm:"default:0"` // 0未使用，1已使用
}

// 接单服务的战士信息
type Warrior struct {
	BaseModel
	Name              string  `json:"name"`
	Phone             uint64  `json:"phone"`
	Birthday          uint    `json:"birthday"`
	Sex               uint8   `json:"sex" gorm:"default:1"` //0女，1男
	JoinTime          int64   `json:"join_time" gorm:"autoCreateTime:milli"`
	BelongBranchId    *string `json:"belong_branch_id"gorm:"default:''"`
	Status            uint8   `json:"status" gorm:"default:0"` //账号状态
	IDCard            string  `json:"id_card"`
	IDCardImageFront  string  `json:"id_card_image_front" gorm:"default:''"`
	IDCardImageBehind string  `json:"id_card_image_behind" gorm:"default:''"`
	// 户籍所在地
	DomicileProvince string `json:"domicile_province" gorm:"default:''"`
	DomicileCity     string `json:"domicile_city" gorm:"default:''"`
	DomicileArea     string `json:"domicile_area" gorm:"default:''"`
	// 居住地址
	Province string `json:"province" gorm:"default:''"`
	City     string `json:"city" gorm:"default:''"`
	Area     string `json:"area" gorm:"default:''"`
	Address  string `json:"address" gorm:"default:''"`
	Remark   string `json:"remark"`
	Created  int64  `json:"created" gorm:"autoUpdateTime:milli"`
}

// 战士评分表
type WarriorScore struct {
	BaseModel
	WarriorId string `json:"warrior_id"`
	Score     uint8  `json:"score"`
	Comment   string `json:"comment"`
	Anonymous uint8  `json:"anonymous" gorm:"default:0"` //是否匿名，0不匿名，1匿名
	Created   int64  `json:"created" gorm:"autoUpdateTime:milli"`
}

// 订单使用的拥有的优惠券(订单表关联，一个订单对应多张优惠券)
type OrderCoupon struct {
	BaseModel
	UserCouponId string `json:"user_coupon_id"`
	Name         string `json:"name"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Description  string `json:"description"`
	OrderID      string `json:"order_id"` //外键
}

// 附加服务
type BaseExtraService struct {
	BaseModel
	Name        string `json:"name"`
	UnitPrice   int    `json:"unit_price"`
	Discount    uint8  `json:"discount"` //这个服务的折扣(1-100)
	Description string `json:"description"`
}
type ExtraService struct {
	BaseExtraService
	Status  uint8 `json:"status" gorm:"default:0"` // 是否在可用状态,1为可用，0为不可用
	Created int64 `json:"created" gorm:"autoUpdateTime:milli"`
}

// 订单使用的附加服务（订单关联,一订单对应多个附加服务）
type OrderExtraService struct {
	OrderID        string `json:"order_id"` //外键
	ExtraServiceId string `json:"extra_service_id"`
	BaseExtraService
}

// 订单
type Order struct {
	BaseModel
	OrderNum           string `json:"order_num"` // 订单编号
	UserID             string `json:"user_id"`
	Status             uint8  `json:"status" gorm:"default:0"`        //订单状态
	TotalAmount        uint   `json:"total_amount"`                   //总金额
	PaidInAmount       uint   `json:"paid_in_amount"`                 // 实收金额
	DiscountAmount     uint   `json:"discount_amount"`                //优惠金额
	RefundStatus       uint8  `json:"refund_status" gorm:"default:0"` //退款状态
	RefundSArrivalTime *int64 `json:"refund_s_arrival_time"`          //退款到账时间
	Created            int64  `json:"created" gorm:"autoUpdateTime:milli"`
	// 接单战士
	Warrior *struct {
		ID             *string `json:"id"`
		Name           *string `json:"name"`
		Phone          *uint64 `json:"phone"`
		BelongBranchId *uint   `json:"belong_branch_id"`
	} `json:"warrior" gorm:"embedded;embeddedPrefix:warrior_"`
	//服务网点信息
	Branch struct {
		ID string `json:"id"`
		BaseBranch
	} `json:"branch" gorm:"embedded;embeddedPrefix:branch_"`
	// 服务地址及联系人信息
	ClientInfo struct {
		UserId    string `json:"user_id"`
		Name      string `json:"name"`
		Phone     uint64 `json:"phone"`
		Province  string `json:"province"`
		City      string `json:"city"`
		Area      string `json:"area"`
		Address   string `json:"address"`
		Distance  uint64 `json:"distance"`   //距离网点的距离
		StartTime int64  `json:"start_time"` //预约的服务时间-开始时间
		EndTime   int64  `json:"end_time"`   //预约的服务时间-结束时间
	} `json:"client_info" gorm:"embedded;embeddedPrefix:client_info_"`
	// 电器信息
	Machine struct {
		Brand         string `json:"brand"` // 品牌
		Type          string `json:"type"`
		Mode          string `json:"mode"`
		PhotosJsonStr string `json:"photos_json_str" gorm:"default:'[]'"`
		Remark        string `json:"remark" gorm:"default:''"`
	} `json:"machine" gorm:"embedded;embeddedPrefix:machine_"`
	ExtraServices []OrderExtraService `json:"extra_services"`
	OrderCoupons  []OrderCoupon       `json:"order_coupons"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = util.UniqueId()
	return
}
func (m *Order) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = util.UniqueId()
	m.OrderNum = util.GenerateOrderNum()
	return
}

func (pages *Pages) CalcPagesData(page int, pageSize int, total int64) {
	pages.Page = page
	pages.PageSize = pageSize
	pages.Total = util.Int64ToInt(total)
	fmt.Println(pages.Total)
	pages.TotalPage = util.CalcTotalPage(pages.Total, pages.PageSize)
}

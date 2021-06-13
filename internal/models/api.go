package models

import "gorm.io/gorm"

type BaseQuery struct {
	Page int `form:"page"`
	PageSize int `form:"page_size"`
}
// 行政区
type Region struct {
	Province string `form:"province"`
	City     string `form:"city"`
	Area     string `form:"area"`
}
// 网点路由查询
type BranchQuery struct {
	Region
	Name string `form:"name"`
	ContactPerson  string `form:"contact_person"` //联系人
	Status string `form:"status"` // 网点状态 0关闭，1营业中，2休息中
	CreatedStartTime uint `form:"created_start_time"`
	CreatedEndTime uint `form:"created_end_time"`
	BaseQuery
}
type CouponQuery struct {
	BaseQuery
	Name string `form:"name"`
	StartTime   int64  `form:"start_time"`
	EndTime     int64  `form:"end_time"`
	CreatedStartTime uint `form:"created_start_time"`
	CreatedEndTime uint `form:"created_end_time"`
}

type ExtraServiceQuery struct {
	BaseQuery
}

type OrderQuery struct {
	BaseQuery
	UserId string `form:"user_id"`
}

type WarriorQuery struct {
	BaseQuery
}


func (query *BaseQuery) AfterUpdate(tx *gorm.DB) (err error) {
	if query.Page == 0 {
		query.Page = 1
	}
	switch {
	case query.PageSize > 100:
		query.PageSize = 100
		break
	case query.PageSize <= 0:
		query.PageSize = 10
		break
	}
	return
}
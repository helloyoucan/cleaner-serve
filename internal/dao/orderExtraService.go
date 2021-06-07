package dao

import "cleaner-serve/internal/models"

/**
订单使用的附加服务
*/
func CreateOrderExtraServiceList(orderExtraServiceList []*models.OrderExtraService) (err error) {
	return DB.Create(&orderExtraServiceList).Error
}
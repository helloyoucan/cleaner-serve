package logic

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
)

func CreateExtraServiceList(orderId string,extraServiceList []*models.OrderExtraService) (err error) {
	for _, orderExtraService := range extraServiceList {
		orderExtraService.ID=util.UniqueId()
		orderExtraService.OrderId=orderId
	}
	return dao.CreateOrderExtraServiceList(extraServiceList)
}

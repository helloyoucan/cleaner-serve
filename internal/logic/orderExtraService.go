package logic

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
)

func CreateExtraServiceList(orderId string,extraServiceList []*models.OrderExtraService) (err error) {
	for _, orderExtraService := range extraServiceList {
		orderExtraService.OrderID=orderId
	}
	return dao.CreateOrderExtraServiceList(extraServiceList)
}

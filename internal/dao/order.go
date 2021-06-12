package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
)

func CreateAOrder(order *models.Order) (err error) {
	return DB.Create(&order).Error
}
func GetOrderByPages(query *models.OrderQuery) (orderList []*models.Order,total int64, err error) {
	err = DB.Preload("ExtraServices").Preload("OrderCoupons").Scopes(util.Paginate(query.Page,query.PageSize)).Find(&orderList).Error
	if err != nil {
		return nil, 0,err
	}
	DB.Model(&models.Order{}).Count(&total)
	return
}
func GetOrderByUserByPages(query *models.OrderQuery) (orderList []*models.Order, total int64, err error) {
	err = DB.Where("user_id=?",query.UserId).Scopes(util.Paginate(query.Page,query.PageSize)).Find(&orderList).Error
	if err != nil {
		return nil,0, err
	}
	DB.Model(&models.Order{}).Where("user_id=?",query.UserId).Count(&total)
	return
}
func GetAOrderById(id string) (order *models.Order, err error) {
	err = DB.Where("id=?",id).Find(&order).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateAOrder(order *models.Order)(err error)  {
	err=DB.Save(&order).Error
	return
}
func DeleteAOrder(id string)(err error)  {
	err = DB.Where("id=?",id).Delete(&models.Order{}).Error
	return
}
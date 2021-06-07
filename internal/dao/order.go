package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
)

func CreateAOrder(order *models.Order) (err error) {
	return DB.Create(&order).Error
}
func GetOrderByPages(pages *models.Pages) (orderList []*models.Order, err error) {
	var total int64
	DB.Model(&models.Order{}).Count(&total)
	pages.CalcPages(total)
	if err != nil {
		return orderList, err
	}
	err = DB.Preload("ExtraServices").Preload("OrderCoupons").Scopes(util.Paginate(pages.Page,pages.PageSize)).Find(&orderList).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetOrderByUserByPages(pages *models.Pages,userId string) (orderList []*models.Order, err error) {
	var total int64
	DB.Model(&models.Order{}).Where("user_id=?",userId).Count(&total)
	pages.CalcPages(total)
	if err != nil {
		return orderList, err
	}
	err = DB.Where("user_id=?",userId).Scopes(util.Paginate(pages.Page,pages.PageSize)).Find(&orderList).Error
	if err != nil {
		return nil, err
	}
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
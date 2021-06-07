package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
)

func CreateAOrder(order *models.Order) (err error) {
	return DB.Create(&order).Error
}
func GetAllOrder() (orderLIst []*models.Order, err error) {
	err = DB.Find(&orderLIst).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetOrderByPages(pages *models.Pages) (orderLIst []*models.Order, err error) {
	err = DB.Scopes(util.Paginate(pages)).Find(&orderLIst).Error
	if err != nil {
		return nil, err
	}
	var total int64
	DB.Model(&models.Order{}).Count(&total)
	err= util.HandlePages(pages,total)
	if err != nil {
		return orderLIst, err
	}
	return
}
func GetOrderByUserByPages(pages *models.Pages,userId string) (orderLIst []*models.Order, err error) {
	err = DB.Where("user_id=?",userId).Scopes(util.Paginate(pages)).Find(&orderLIst).Error
	if err != nil {
		return nil, err
	}
	var total int64
	DB.Model(&models.Order{}).Where("user_id=?",userId).Count(&total)
	err= util.HandlePages(pages,total)
	if err != nil {
		return orderLIst, err
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
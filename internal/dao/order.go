package dao

import "cleaner-serve/internal/models"

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
func GetAOrderById(id string) (orderLIst []*models.Order, err error) {
	err = DB.Where("id=?",id).Find(&orderLIst).Error
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
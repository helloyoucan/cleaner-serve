package dao

import "cleaner-serve/internal/models"

func CreateAClientInfo(clientInfo *models.ClientInfo) (id string, err error) {
	return clientInfo.ID, DB.Create(&clientInfo).Error
}
func GetACouponByUserId(id string) (clientInfo []*models.ClientInfo, err error) {
	err = DB.Where("id = ?",id).Find(&clientInfo).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateAClientInfo(clientInfo *models.ClientInfo)(err error)  {
	err=DB.Save(&clientInfo).Error
	return
}
func DeleteAClientInfo(id string)(err error)  {
	err = DB.Where("id=?",id).Delete(&models.ClientInfo{}).Error
	return
}
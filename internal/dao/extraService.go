package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
)

func CreateAExtraService(extraService *models.ExtraService) (err error) {
	return DB.Create(&extraService).Error
}
func GetAllExtraService() (extraServiceLIst []*models.ExtraService, err error) {
	err = DB.Find(&extraServiceLIst).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetExtraServiceByPages(query *models.ExtraServiceQuery) (extraServiceList []*models.ExtraService, total int64, err error) {
	err = DB.Order("created desc").Scopes(util.Paginate(query.Page,query.PageSize)).Find(&extraServiceList).Error
	if err != nil {
		return nil,0, err
	}
	DB.Model(&models.ExtraService{}).Count(&total)
	return
}
func GetAllActiveExtraService() (extraServiceList []*models.ExtraService, err error) {
	err = DB.Where("is_active=?",1).Find(&extraServiceList).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetAExtraServiceById(id string) (extraServiceLIst *models.ExtraService, err error) {
	err = DB.Where("id=?",id).Find(&extraServiceLIst).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateAExtraService(extraService *models.ExtraService)(err error)  {
	err=DB.Save(&extraService).Error
	return
}
func DeleteAExtraService(id string)(err error)  {
	err = DB.Where("id=?",id).Delete(&models.ExtraService{}).Error
	return
}
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
func GetExtraServiceByPages(pages *models.Pages) (extraServiceList []*models.ExtraService, err error) {
	var total int64
	DB.Model(&models.ExtraService{}).Count(&total)
	pages.CalcPages(total)
	if err != nil {
		return extraServiceList, err
	}
	err = DB.Scopes(util.Paginate(pages.Page,pages.PageSize)).Find(&extraServiceList).Error
	if err != nil {
		return nil, err
	}
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
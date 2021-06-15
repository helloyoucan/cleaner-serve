package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"errors"
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
	db:=DB.Scopes(util.QueryCreated(query.CreatedStartTime, query.CreatedEndTime)).Order("created desc")
	if query.Name!=""{
		db.Where("name LIKE ?","%"+query.Name+"%")
	}
	if query.Discount!=0{
		db.Where("discount = ?",query.Discount)
	}
	if query.Status!=""{
		db.Where("status = ?",query.Status)
	}
	db.Model(&models.ExtraService{}).Count(&total)
	err = db.Scopes(util.Paginate(query.Page,query.PageSize)).Find(&extraServiceList).Error
	if err != nil {
		return nil,0, err
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
	result:=DB.Where("id=?",extraService.ID).Select("*").Updates(&extraService)
	if result.RowsAffected==0{
		return errors.New("数据不存在")
	}
	return  result.Error
}
func DeleteAExtraService(id string)(err error)  {
	err = DB.Where("id=?",id).Delete(&models.ExtraService{}).Error
	return
}
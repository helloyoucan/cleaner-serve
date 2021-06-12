package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
)

func CreateAWarrior(warrior *models.Warrior) (err error) {
	return DB.Create(&warrior).Error
}
func GetAllWarrior() (warriorList []*models.Warrior, err error) {
	err = DB.Find(&warriorList).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetWarriorByPages(query *models.WarriorQuery) (warrior []*models.Warrior, total int64,err error) {
	err = DB.Scopes(util.Paginate(query.Page,query.PageSize)).Find(&warrior).Error
	if err != nil {
		return nil,0, err
	}
	DB.Model(&models.Warrior{}).Count(&total)
	return
}
func GetAWarriorById(id string) (warrior *models.Warrior, err error) {
	err = DB.Where("id=?",id).Find(&warrior).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateAWarrior(warrior *models.Warrior)(err error)  {
	err=DB.Save(&warrior).Error
	return
}
func DeleteAWarrior(id string)(err error)  {
	err = DB.Where("id=?",id).Delete(&models.Warrior{}).Error
	return
}
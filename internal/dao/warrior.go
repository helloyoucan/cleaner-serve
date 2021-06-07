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
func GetWarriorByPages(pages *models.Pages) (warrior []*models.Warrior, err error) {
	var total int64
	DB.Model(&models.Warrior{}).Count(&total)
	pages.CalcPages(total)
	if err != nil {
		return warrior, err
	}
	err = DB.Scopes(util.Paginate(pages.Page,pages.PageSize)).Find(&warrior).Error
	if err != nil {
		return nil, err
	}
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
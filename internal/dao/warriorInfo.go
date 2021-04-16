package dao

import "cleaner-serve/internal/models"

func CreateAWarriorInfo(warriorInfo *models.WarriorInfo) (err error) {
	return DB.Create(&warriorInfo).Error
}
func GetAllWarriorInfo() (warriorInfoLIst []*models.WarriorInfo, err error) {
	err = DB.Find(&warriorInfoLIst).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetAWarriorInfoById(id string) (warriorInfoLIst []*models.WarriorInfo, err error) {
	err = DB.Where("id=?",id).Find(&warriorInfoLIst).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateAWarriorInfo(warriorInfo *models.WarriorInfo)(err error)  {
	err=DB.Save(&warriorInfo).Error
	return
}
func DeleteAWarriorInfo(id string)(err error)  {
	err = DB.Where("id=?",id).Delete(&models.WarriorInfo{}).Error
	return
}
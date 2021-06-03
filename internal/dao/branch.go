package dao

import "cleaner-serve/internal/models"

func CreateABranch(branch *models.Branch) (err error) {
	return DB.Create(&branch).Error
}
func GetAllBranch() (branchLIst []*models.Branch, err error) {
	err = DB.Find(&branchLIst).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetABranchById(id string) (branch *models.Branch, err error) {
	branch = new(models.Branch)
	err = DB.Where("id=?",id).Find(&branch).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateABranch(branch *models.Branch)(err error)  {
	err=DB.Save(&branch).Error
	return
}
func DeleteABranch(id string)(err error)  {
	err = DB.Where("id=?",id).Delete(&models.Branch{}).Error
	return
}
package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
)

func CreateABranch(branch *models.Branch) (err error) {
	return DB.Create(&branch).Error
}
func GetBranchByPages(pages *models.Pages) (branchLIst []*models.Branch, err error) {
	err = DB.Scopes(util.Paginate(pages)).Find(&branchLIst).Error
	if err != nil {
		return nil, err
	}
	var total int64
	DB.Model(&models.Branch{}).Count(&total)
	err= util.HandlePages(pages,total)
	if err != nil {
		return branchLIst, err
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
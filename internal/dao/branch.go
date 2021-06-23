package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"errors"
)

func CreateABranch(branch *models.Branch) (err error) {
	return DB.Create(&branch).Error
}

func GetBranchByPages(query *models.BranchQuery) (branchList []*models.Branch, total int64, err error) {
	var region models.Region
	region.Province = query.Province
	region.City = query.City
	region.Area = query.Area
	db := DB.Where(&region).Scopes(util.QueryCreated(query.CreatedStartTime, query.CreatedEndTime)).Order("created desc")
	if query.Status != "" {
		db.Where("status = ?", util.StrToUint(query.Status))
	}
	if query.Name != "" {
		db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Address != "" {
		db.Where("address LIKE ?", "%"+query.Address+"%")
	}
	if query.ContactPerson != "" {
		db.Where("contact_person LIKE ?", "%"+query.ContactPerson+"%")
	}
	db.Model(&models.Branch{}).Count(&total)
	err = db.Scopes(util.Paginate(query.Page, query.PageSize)).Find(&branchList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}
func GetAllBranch() (branchList []*models.Branch, err error) {
	err = DB.Order("created desc").Find(&branchList).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetABranchById(id string) (branch *models.Branch, err error) {
	branch = new(models.Branch)
	err = DB.Where("id=?", id).Find(&branch).Error
	if err != nil {
		return nil, err
	}
	return
}
func UpdateABranch(branch *models.Branch) (err error) {
	result := DB.Where("id=?", branch.ID).Select("*").Updates(&branch)
	if result.RowsAffected == 0 {
		return errors.New("数据不存在")
	}
	return result.Error
}
func DeleteBranch(ids []string) (delCount int64, err error) {
	result := DB.Delete(&models.Branch{}, ids)
	return result.RowsAffected, result.Error
}

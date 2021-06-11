package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"errors"
	"fmt"
)

func CreateABranch(branch *models.Branch) (err error) {
	return DB.Create(&branch).Error
}
func GetBranchByPages(pages *models.Pages,query *models.BranchQuery) (branchLIst []*models.Branch, err error) {
	var total int64
	DB.Model(&models.Branch{}).Count(&total)
	pages.CalcPages(total)
	if err != nil {
		return branchLIst, err
	}
	fmt.Println("------")
	fmt.Println(query.Name)
	err = DB.Where(&query).Scopes(util.Paginate(pages.Page,pages.PageSize)).Find(&branchLIst).Error
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
//func UpdateABranch(branch *models.Branch)(err error)  {
//	err=DB.Save(&branch).Error
//	return
//}
func UpdateABranch(branch *models.Branch)(err error)  {
	result:=DB.Where("id=?",branch.ID).Select("*").Updates(&branch)
	if result.RowsAffected==0{
		return errors.New("数据不存在")
	}
	return result.Error
}
func DeleteBranch(ids []string)(delCount int64,err error)  {
	result := DB.Delete(&models.Branch{},ids)
	return result.RowsAffected,result.Error
}
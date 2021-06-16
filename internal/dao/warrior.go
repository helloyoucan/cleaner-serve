package dao

import (
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"errors"
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
	db:=DB
	if query.Status != "" {
		db=db.Where("status = ?", util.StrToUint(query.Status))
	}
	if query.Name != "" {
		db=db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.BelongBranchId != "" {
		db=db.Where("belong_branch_id = ?", query.BelongBranchId)
	}
	db.Model(&models.Warrior{}).Count(&total)
	err = db.Order("created desc").Scopes(util.Paginate(query.Page,query.PageSize)).Find(&warrior).Error
	if err != nil {
		return nil,0, err
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
	result := DB.Where("id=?", warrior.ID).Select("*").Updates(&warrior)
	if result.RowsAffected == 0 {
		return errors.New("数据不存在")
	}
	return result.Error
}
func DeleteAWarrior(id string)(err error)  {
	err = DB.Where("id=?",id).Delete(&models.Warrior{}).Error
	return
}
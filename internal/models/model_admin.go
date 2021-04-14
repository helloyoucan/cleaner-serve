package models

import "github.com/jinzhu/gorm"

/**
后台管理特有
 */
type AdminUser struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
}

package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

func CreateAWarrior(c *gin.Context)  {
	var warrior models.Warrior
	err:=c.BindJSON(&warrior)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err=dao.CreateAWarrior(&warrior)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{})
}
func GetWarriorByPages(c *gin.Context)  {
	var query =new(models.WarriorQuery)
	err:=c.ShouldBind(&query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	warriorList,total,err:=dao.GetWarriorByPages(query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	var pages models.Pages
	pages.CalcPagesData(query.Page,query.PageSize,total)
	util.RespJSON(c,gin.H{
		"data": gin.H{
			"list":warriorList,
			"pages":pages,
		},
	})
}

func UpdateAWarrior(c *gin.Context)  {
	var warrior models.Warrior
	err:=c.BindJSON(&warrior)
	if err!=nil {
		util.RespJSON(c,gin.H{
			"err":err.Error(),
		})
		return
	}
	err = dao.UpdateAWarrior(&warrior)
	if err!=nil {
		util.RespJSON(c,gin.H{
			"err":err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{})
}
func DeleteAWarrior(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	err := dao.DeleteAWarrior(id)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{})
}
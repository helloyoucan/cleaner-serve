package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
	"strconv"
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
	var pages =new (models.Pages)
	pages.Page,_ = strconv.Atoi(c.Query("page"))
	pages.PageSize,_ = strconv.Atoi(c.Query("page_size"))
	warriorList,err:=dao.GetWarriorByPages(pages)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{
		"data": gin.H{
			"list":warriorList,
			"pages":pages,
		},
	})
}

func UpdateAWarrior(c *gin.Context)  {
	id := c.Query("id")
	if id=="" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	warrior,err:=dao.GetAWarriorById(id)
	if err!=nil {
		util.RespJSON(c,gin.H{
			"err":err.Error(),
		})
		return
	}
	err=c.BindJSON(&warrior)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err=dao.UpdateAWarrior(warrior)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{"data":warrior,})
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
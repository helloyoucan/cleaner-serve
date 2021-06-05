package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateAExtraService(c *gin.Context)  {
	var extraService models.ExtraService
	err:=c.BindJSON(&extraService)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	extraService.ID=util.UniqueId()
	err=dao.CreateAExtraService(&extraService)
	util.RespJSON(c,gin.H{
		"err":err,
	})
}

func GetExtraServiceByPages(c *gin.Context)  {
	var pages =new (models.Pages)
	pages.Page,_ = strconv.Atoi(c.Query("page"))
	pages.PageSize,_ = strconv.Atoi(c.Query("page_size"))
	extraServiceList,err:=dao.GetExtraServiceByPages(pages)
	util.RespJSON(c,gin.H{
		"err":err,
		"data": gin.H{
			"list":extraServiceList,
			"pages":pages,
		},
	})
}
func UpdateAExtraService(c *gin.Context)  {
	id := c.Query("id")
	if id=="" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	branch,err:=dao.GetABranchById(id)
	if err!=nil {
		util.RespJSON(c,gin.H{
			"err":err.Error(),
		})
		return
	}
	err=c.BindJSON(&branch)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err=dao.UpdateABranch(branch)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{"data":branch,})
}
func DeleteAExtraService(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	err := dao.DeleteABranch(id)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{})

}
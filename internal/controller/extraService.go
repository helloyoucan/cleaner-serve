package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

func CreateAExtraService(c *gin.Context)  {
	var extraService models.ExtraService
	err:=c.BindJSON(&extraService)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err=dao.CreateAExtraService(&extraService)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{})
}

func GetExtraServiceByPages(c *gin.Context)  {
	var query =new(models.ExtraServiceQuery)
	err:=c.ShouldBind(&query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	extraServiceList,total,err:=dao.GetExtraServiceByPages(query)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	var pages models.Pages
	pages.CalcPagesData(query.Page,query.PageSize,total)
	util.RespJSON(c,gin.H{
		"data": gin.H{
			"list":extraServiceList,
			"pages":pages,
		},
	})
}
func GetAllActiveExtraService(c *gin.Context)  {
	extraServiceList,err:=dao.GetAllActiveExtraService()
	if err!=nil {
		util.RespJSON(c,gin.H{
			"err":err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{"data":extraServiceList,})
}
func UpdateAExtraService(c *gin.Context)  {
	id := c.Query("id")
	if id=="" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	extraService,err:=dao.GetAExtraServiceById(id)
	if err!=nil {
		util.RespJSON(c,gin.H{
			"err":err.Error(),
		})
		return
	}
	err=c.BindJSON(&extraService)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err=dao.UpdateAExtraService(extraService)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{"data":extraService,})
}
func DeleteAExtraService(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	err := dao.DeleteAExtraService(id)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{})

}
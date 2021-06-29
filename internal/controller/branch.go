package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

func CreateABranch(c *gin.Context)  {
	var branch models.Branch
	err:=c.BindJSON(&branch)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	err=dao.CreateABranch(&branch)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{})
}
func GetBranchByPages(c *gin.Context)  {
	var query =new(models.BranchQuery)
	err:=c.ShouldBind(&query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	branchList,total,err:=dao.GetBranchByPages(query)
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
			"list":branchList,
			"pages":pages,
		},
	})
}
func GetAllBranch(c *gin.Context)  {
	var query =new(models.AllBranchQuery)
	err:=c.ShouldBind(&query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	branchList,err:=dao.GetAllBranch(query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{
		"data": branchList,
	})
}
func UpdateABranch(c *gin.Context)  {
	var branch models.Branch
	err:=c.BindJSON(&branch)
	if err!=nil {
		util.RespJSON(c,gin.H{
			"err":err.Error(),
		})
		return
	}
	err = dao.UpdateABranch(&branch)
	if err!=nil {
		util.RespJSON(c,gin.H{
			"err":err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{})
}
func DeleteBranchByIds(c *gin.Context)  {
	ids := c.QueryArray("id")
	if len(ids)==0 {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	delCount,err:=dao.DeleteBranch(ids)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{"data":delCount})
}
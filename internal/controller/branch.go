package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
	var pages =new (models.Pages)
	pages.Page,_ = strconv.Atoi(c.Query("page"))
	pages.PageSize,_ = strconv.Atoi(c.Query("page_size"))
	var query =new(models.BranchQuery)
	err:=c.ShouldBind(&query)
	// todo 处理status 0值问题
	// todo 处理created时间范围问题
	fmt.Println("------")
	fmt.Println(query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	branchList,err:=dao.GetBranchByPages(pages,query)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c,gin.H{
		"data": gin.H{
			"list":branchList,
			"pages":pages,
		},
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
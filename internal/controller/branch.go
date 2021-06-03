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
			"err": err,
		})
		return
	}
	branch.ID=util.UniqueId()
	err=dao.CreateABranch(&branch)
	util.RespJSON(c,gin.H{
		"err":err,
	})
}
func GetAllBranch(c *gin.Context)  {
	branchList,err:=dao.GetAllBranch()
	util.RespJSON(c,gin.H{
		"err":err,
		"data":branchList,
	})
}
func UpdateABranch(c *gin.Context)  {
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
func DeleteABranch(c *gin.Context)  {
	id := c.Query("id")
	if id=="" {
		util.RespJSON(c, gin.H{
			"err": "id无效",
		})
		return
	}
	err:=dao.DeleteABranch(id)
	if err != nil {
		util.RespJSON(c, gin.H{
			"err": err.Error(),
		})
		return
	}
	util.RespJSON(c, gin.H{})
}
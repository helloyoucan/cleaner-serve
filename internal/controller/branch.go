package controller

import (
	"cleaner-serve/internal/dao"
	"cleaner-serve/internal/models"
	"cleaner-serve/internal/util"
	"github.com/gin-gonic/gin"
)

func CreateABranch(c *gin.Context)  {
	var branch models.Branch
	c.BindJSON(&branch)
	err:=dao.CreateABranch(&branch)
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
	id,ok:=c.Params.Get("id")
	if !ok{
		util.RespJSON(c,gin.H{
			"err":"id无效",
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
	c.BindJSON(&branch)
	err=dao.UpdateABranch(branch)
	util.RespJSON(c,gin.H{
		"err":err.Error(),
	})
}
func DeleteABranch(c *gin.Context)  {
	id,ok:=c.Params.Get("id")
	if !ok {
		util.RespJSON(c,gin.H{
			"err":"id无效",
		})
	}
	err:=dao.DeleteABranch(id)
	util.RespJSON(c,gin.H{
		"err":err.Error(),
	})
}
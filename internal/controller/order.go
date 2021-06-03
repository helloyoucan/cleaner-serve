package controller

import (
	"cleaner-serve/internal/logic"
	"cleaner-serve/internal/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type  APIOrder struct {
	//BranchId           uint    `json:"branch_id"`      //服务网点
	//ClientInfoId       uint    `json:"client_info_id"` //客户信息id
	Status             uint8   `json:"status"`
	StartTime          int64   `json:"start_time"`
	EndTime            int64   `json:"end_time"`
	MachineId          uint    `json:"machine_id"`
	Distance           uint64  `json:"distance"`              //距离网点
	TotalPrice         float64 `json:"total_price"`           //总费用
	WarriorId          *uint   `json:"warrior_id"`            //接单战士id
	RefundStatus       uint8   `json:"refund_status"`         //退款状态
	RefundSArrivalTime int64   `json:"refund_s_arrival_time"` //退款到账时间
	ClientInfo struct{
		UserId   uint   `json:"user_id"`
		Name     string `json:"name"`
		Phone    uint64 `json:"phone"`
		Province string `json:"province"`
		City     string `json:"city"`
		Area     string `json:"area"`
		Address  string `json:"address"`
	} `json:"client_info"`
	Machine struct{
		Brand         string `json:"brand"` // 品牌
		Type          string `json:"type"`
		Mode          string `json:"mode"`
		PhotosJsonStr string `json:"photos_json_str" gorm:"default:''"`
		Remark        string `json:"remark"  gorm:"default:''"`
	} `json:"machine"`
}
func CreateAOrder(c *gin.Context)  {
	var apiOrder APIOrder
	err:=c.BindJSON(apiOrder)
	if err!=nil {
		util.RespJSON(c, gin.H{
			"err": err,
		})
		return
	}
	// 创建一条客户信息数据
	clientInfoId,err := logic.CreateAClientInfo(apiOrder.ClientInfo)
	fmt.Print("----"+clientInfoId)
	// 创建一条机器信息数据
}
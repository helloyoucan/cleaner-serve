package logic
//
//import (
//	"cleaner-serve/internal/dao"
//	"cleaner-serve/internal/models"
//	"cleaner-serve/internal/util"
//	"encoding/json"
//)
//
//func CreateAClientInfo(jsonClientInfo interface{})(id string,err error)  {
//	data, err := json.Marshal(jsonClientInfo)
//	var clientInfo models.ClientInfo
//	if err != nil{
//		return
//	}
//	err = json.Unmarshal(data, &clientInfo)
//	if err != nil {
//		return
//	}
//	clientInfo.ID=util.UniqueId()
//	id,err=dao.CreateAClientInfo(&clientInfo)
//	return
//}

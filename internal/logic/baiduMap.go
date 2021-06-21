package logic

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"strconv"
)

const AK = "WahIhIS6713Mxfc34t1PI759xGfOorQM"
const locationIpURL = "http://api.map.baidu.com/location/ip?ak=" + AK + "&coor=bd09ll"
func GetLocationInfoByBaiduMap(ip string)(respMap map[string]interface{},err error) {
	url := locationIpURL
	if ip != "127.0.0.1" {
		url = url + "&ip=" + ip
	}
	status, resp, err := fasthttp.Get(nil, locationIpURL)
	if err != nil||status != fasthttp.StatusOK {
		return nil,errors.New("status is"+strconv.Itoa(status))
	}
	respMap = make(map[string]interface{})
	err = json.Unmarshal(resp, &respMap)
	return respMap,err
}

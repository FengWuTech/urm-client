package urmclient

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

type URMResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

//发送短信
/*
	appID: 分配的渠道号
	appSecretKey: 分配的秘钥
   	mobiles: 手机号列表
   	tplCode: 模板id(内部分配)
   	param: 模板变量map
*/
func SendSMS(appID string, appSecretKey string, mobiles []string, tplCode int, param interface{}) (bool, URMResponse) {

	path := "/urm/sms/send"
	query := genQuery(appID, appSecretKey)
	rawURL := SERVER_ADDRESS + path + "?" + query

	params := map[string]interface{}{
		"mobiles": mobiles,
		"tpl_id":  tplCode,
		"params":  param,
	}
	paramJSON, _ := json.Marshal(params)
	paramStr := string(paramJSON)

	var urlResp URMResponse
	request := gorequest.New()
	_, _, errs := request.Post(rawURL).Send(paramStr).EndStruct(&urlResp)
	if errs != nil || len(errs) > 0 {
		return false, urlResp
	}
	return true, urlResp
}

//获取短信发送量按天统计数据
func GetSMSStatistics(appID string, appSecretKey string) (bool, URMResponse) {
	path := "/urm/sms/statistics"
	query := genQuery(appID, appSecretKey)
	rawURL := SERVER_ADDRESS + path + "?" + query

	var urlResp URMResponse
	request := gorequest.New()
	_, _, errs := request.Get(rawURL).EndStruct(&urlResp)
	if errs != nil || len(errs) > 0 {
		return false, urlResp
	}
	return true, urlResp
}

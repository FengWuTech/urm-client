package urmclient

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

type URMSendSmsResponse struct {
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
func SendSMS(appID string, appSecretKey string, mobiles []string, tplCode int, param interface{}) (bool, URMSendSmsResponse) {

	path := "/urm/sms/send"
	query := genQuery(appID, appSecretKey, nil)
	rawURL := SERVER_ADDRESS + path + "?" + query

	params := map[string]interface{}{
		"mobiles": mobiles,
		"tpl_id":  tplCode,
		"params":  param,
	}
	paramJSON, _ := json.Marshal(params)
	paramStr := string(paramJSON)

	var urlResp URMSendSmsResponse
	request := gorequest.New()
	_, _, errs := request.Post(rawURL).Send(paramStr).EndStruct(&urlResp)
	if errs != nil || len(errs) > 0 {
		return false, urlResp
	}
	return true, urlResp
}

type URMStatisticsResponseItem struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type URMStatisticsResponse struct {
	Code int                         `json:"code"`
	Msg  string                      `json:"msg"`
	Data []URMStatisticsResponseItem `json:"data"`
	Time string                      `json:"time"`
}

//获取短信发送量按天统计数据
func GetSMSStatistics(appID string, appSecretKey string) (bool, URMStatisticsResponse) {
	path := "/urm/sms/statistics"
	query := genQuery(appID, appSecretKey, nil)
	rawURL := SERVER_ADDRESS + path + "?" + query

	var urlResp URMStatisticsResponse
	request := gorequest.New()
	_, _, errs := request.Get(rawURL).EndStruct(&urlResp)
	if errs != nil || len(errs) > 0 {
		return false, urlResp
	}
	return true, urlResp
}

type URMAddTplResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		ID    int    `json:"id"`
		TplID string `json:"tpl_id"`
	} `json:"data"`
	Time string `json:"time"`
}

func AddSmsTpl(appID string, appSecretKey string, typ int, name string, content string) (bool, *URMAddTplResponse) {
	path := "/urm/sms/tpl/add"
	query := genQuery(appID, appSecretKey, nil)
	rawURL := SERVER_ADDRESS + path + "?" + query

	var ret URMAddTplResponse
	params := map[string]interface{}{
		"appid": appID,
	}
	paramJSON, _ := json.Marshal(params)
	paramStr := string(paramJSON)
	_, _, errs := gorequest.New().Post(rawURL).Send(paramStr).EndStruct(&ret)
	if errs != nil {
		return false, nil
	} else {
		return true, &ret
	}
}

type URMGetTplResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		Status int    `json:"status"`
		Reason string `json:"reason"`
	} `json:"data"`
	Time string `json:"time"`
}

func GetSmsTpl(appID string, appSecretKey string, tplID string) (bool, *URMGetTplResponse) {
	path := "/urm/sms/tpl/get"
	query := genQuery(appID, appSecretKey, map[string]string{
		"tpl_id": tplID,
	})
	rawURL := SERVER_ADDRESS + path + "?" + query

	var ret URMGetTplResponse
	_, _, errs := gorequest.New().Get(rawURL).EndStruct(&ret)
	if errs != nil {
		return false, nil
	} else {
		return true, &ret
	}
}

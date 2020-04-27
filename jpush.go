package urmclient

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

type URMJPushReq struct {
	Platform     interface{}            `json:"platform" valid:"Required;"` //推送平台设置
	Audience     interface{}            `json:"audience" valid:"Required;"` //推送设备指定
	Notification map[string]interface{} `json:"notification"`               //通知内容体
	Message      map[string]interface{} `json:"message"`                    //消息内容体
	Options      map[string]interface{} `json:"options"`                    //推送参数
}

type URMJPushResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

//发送短信
/*
	appID: 分配的渠道号
	appSecretKey: 分配的秘钥
   	req: 推送请求体 https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/
*/
func (urm *URM) SendJPush(req URMJPushReq) (bool, URMJPushResponse) {

	path := "/urm/jpush/send"
	query := urm.genQuery(nil)
	rawURL := urm.BaseURL + path + "?" + query

	paramJSON, _ := json.Marshal(req)
	paramStr := string(paramJSON)

	var urlResp URMJPushResponse
	request := gorequest.New()
	_, _, errs := request.Post(rawURL).Send(paramStr).EndStruct(&urlResp)
	if errs != nil || len(errs) > 0 {
		return false, urlResp
	}
	return true, urlResp
}

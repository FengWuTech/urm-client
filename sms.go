package urmclient

import (
	"encoding/json"
	"time"

	"github.com/parkingwang/go-sign"
	"github.com/parnurzeal/gorequest"
	uuid "github.com/satori/go.uuid"
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
   	params: 参数map
*/
func SendSMS(appID string, appSecretKey string, mobiles []string, tplCode int, param interface{}) (bool, URMResponse) {

	path := "/urm/sms/send"
	rawURL := SERVER_ADDRESS + path

	timeUnix := time.Now().Unix()

	signer := sign.NewGoSignerMd5()
	signer.SetKeyNameAppId("appid")
	signer.SetKeyNameTimestamp("timestamp")
	signer.SetKeyNameNonceStr("nonce_str")

	signer.SetAppId(appID)
	signer.SetTimeStamp(timeUnix)

	nonce := uuid.NewV4().String()
	signer.SetNonceStr(nonce)

	signer.SetAppSecretWrapBody(appSecretKey)

	query := signer.GetSignedQuery()
	rawURL = rawURL + "?" + query

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

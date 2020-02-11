package urmclient

import (
	"testing"
)

const APP_ID = "fpms"
const APP_SECRET_KEY = "rcqBYyBHAJ123"

//测试发送短信
func TestSendSms(t *testing.T) {

	mobiles := []string{"18600320375"}
	tplCode := 2
	param := map[string]interface{}{
		"code": "2345",
	}
	isSucceed, result := SendSMS(APP_ID, APP_SECRET_KEY, mobiles, tplCode, param)
	if !isSucceed || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}

}

//获取指定渠道的短信发送量(按天统计，最近10天)
/*
   data:[{"date":"2020-02-02","count":20},....]
*/
func TestGetSmsStatistics(t *testing.T) {

	isSucceed, result := GetSMSStatistics(APP_ID, APP_SECRET_KEY)
	if !isSucceed || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}
}

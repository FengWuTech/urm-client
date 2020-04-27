package urmclient

import (
	"fmt"
	"testing"
)

const APP_ID = "test"
const APP_SECRET_KEY = "6tNgqxXRbo"

func NewURM() *URM {
	return New(APP_ID, APP_SECRET_KEY, BASE_URL_DEV)
}

//测试发送短信
func TestSendSms(t *testing.T) {

	mobiles := []string{"18511775048"}
	tplCode := 42

	param := []SmsParam{
		{
			Key:   "code",
			Value: "2345",
		},
	}
	isSucceed, result := NewURM().SendSMS(mobiles, tplCode, param, map[string]interface{}{"charge_id": 10})
	if !isSucceed || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}

}

//获取指定渠道的短信发送量(按天统计，最近10天)
/*
   data:[{"date":"2020-02-02","count":20},....]
*/
func TestGetSmsStatistics(t *testing.T) {

	isSucceed, result := NewURM().GetSMSStatistics(1)
	if !isSucceed || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}
}

func TestAddSmsTpl(t *testing.T) {
	success, ret := NewURM().AddSmsTpl(1, "client测试", "client测试")
	fmt.Printf("%v %v", success, ret)
}

func TestGetSmsTpl(t *testing.T) {
	success, ret := NewURM().GetSmsTpl(16)
	fmt.Printf("%v %v", success, ret)
}

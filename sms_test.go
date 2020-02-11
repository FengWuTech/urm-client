package urmclient

import (
	"testing"
)

//测试发送短信
func TestSendSms(t *testing.T) {

	mobiles := []string{"18600320375"}
	tplCode := 2
	param := map[string]interface{}{
		"code": "2345",
		"n": 2,
	}
	isSucceed, result := SendSMS(mobiles, tplCode, param)
	if !isSucceed || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}

}

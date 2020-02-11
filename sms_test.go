package urmclient

import (
	"testing"
)

//测试发送短信
func TestSendSms(t *testing.T) {

	appID := "fpms"
	appSecretKey := "rcqBYyBHAJ"

	mobiles := []string{"18600320375"}
	tplCode := 2
	param := map[string]interface{}{
		"code": "2345",
	}
	isSucceed, result := SendSMS(appID, appSecretKey, mobiles, tplCode, param)
	if !isSucceed || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}

}

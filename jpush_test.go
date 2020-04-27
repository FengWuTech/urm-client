package urmclient

import (
	"encoding/json"
	"testing"
)

const JPUSH_APP_ID = "test"
const JPUSH_APP_SECRET_KEY = "6tNgqxXRbo"

//测试发送短信
func TestSendJPush(t *testing.T) {
	raw := `{
		"cid": "8103a4c628a0b98974ec1949-711261d4-5f17-4d2f-a855-5e5a8909b26e",
		"platform": "all",
		"audience": {
			"tag": [
				"深圳",
				"北京"
			]
		},
		"notification": {
			"android": {
				"alert": "Hi, JPush!",
				"title": "Send to Android",
				"builder_id": 1,
				"large_icon": "http://www.jiguang.cn/largeIcon.jpg",
				"intent": {
					"url": "intent:#Intent;component=com.jiguang.push/com.example.jpushdemo.SettingActivity;end"
				},
				"extras": {
					"newsid": 321
				}
			},
			"ios": {
				"alert": "Hi, JPush!",
				"sound": "default",
				"badge": "+1",
				"thread-id": "default",
				"extras": {
					"newsid": 321
				}
			}
		},
		"message": {
			"msg_content": "Hi,JPush",
			"content_type": "text",
			"title": "msg",
			"extras": {
				"key": "value"
			}
		},
		"sms_message":{
		   "temp_id":1250,
		   "temp_para":{
				"code":"123456"
		   },
			"delay_time":3600,
			"active_filter":false
		},
		"options": {
			"time_to_live": 60,
			"apns_production": false,
			"apns_collapse_id":"jiguang_test_201706011100"
		},
		"callback": {
			"url":"http://www.bilibili.com", 
			"params":{
				"name":"joe",
				"age":26
			 },
			 "type":3
		}
	}`
	req := URMJPushReq{}
	err := json.Unmarshal([]byte(raw), &req)
	if err != nil {
		t.Errorf("Parse json err: %v", err.Error())
	}
	isSucceed, result := NewURM().SendJPush(req)
	if !isSucceed || result.Code != 200 {
		t.Errorf("request failed, code[%v]msg[%v]", result.Code, result.Msg)
	}

}

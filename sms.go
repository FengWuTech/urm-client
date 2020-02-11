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
   mobiles: 手机号列表
   tplCode: 模板id(内部分配)
   params: 参数map
*/
func SendSMS(mobiles []string, tplCode int, param interface{}) (bool, URMResponse) {

	path := "/urm/sms/send"
	rawURL := SERVER_ADDRESS + path

	//1. 计算签名
	timeUnix := time.Now().Unix()

	signer := sign.NewGoSignerMd5()
	signer.SetKeyNameAppId("appid")
	signer.SetKeyNameTimestamp("timestamp")
	signer.SetKeyNameNonceStr("nonce_str")

	// 设置签名基本参数
	signer.SetAppId(APP_ID)
	signer.SetTimeStamp(timeUnix)

	nonce := uuid.NewV4().String()
	signer.SetNonceStr(nonce)

	// 设置参与签名的其它参数
	//signer.AddBody("plate_number", "京A66666")

	// AppSecretKey，前后包装签名体字符串
	signer.SetAppSecretWrapBody(APP_SECRETKEY)

	/*
	   fmt.Println("生成签字字符串：" + signer.GetUnsignedString())
	   fmt.Println("输出URL字符串：" + signer.GetSignedQuery())

	   //生成签字字符串：d93047a4d6fe6111appid=9d8a121ce581499d&nonce_str=ibuaiVcKdpRxkhJA&plate_number=京A66666&time_stamp=1532585241d93047a4d6fe6111
	   //输出URL字符串：appid=9d8a121ce581499d&nonce_str=ibuaiVcKdpRxkhJA&plate_number=京A66666&time_stamp=1532585241&sign=072defd1a251dc58e4d1799e17ffe7a4
	*/

	//2. 拼凑query
	/*
			v := url.Values{}
			v.Add("appid", APP_ID)
			v.Add("timestamp", strconv.FormatInt(timeUnix, 10))
			v.Add("sign", sign)
		    query := v.Encode()
	*/
	query := signer.GetSignedQuery()
	rawURL = rawURL + "?" + query

	//3. 实际发送请求
	params := map[string]interface{}{
		"mobiles": mobiles,
		"tpl_id":  tplCode,
		"params":  param,
	}
	paramJson, _ := json.Marshal(params)
	paramStr := string(paramJson)

	var urlResp URMResponse
	request := gorequest.New()
	_, _, errs := request.Post(rawURL).Send(paramStr).EndStruct(&urlResp)
	if errs != nil || len(errs) > 0 {
		return false, urlResp
	}
	return true, urlResp
}

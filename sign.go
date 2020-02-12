package urmclient

import (
	"time"

	"github.com/parkingwang/go-sign"
	uuid "github.com/satori/go.uuid"
)

//内部函数
//生成请求Query
/*
	appID: 分配的渠道号
	appSecretKey: 分配的秘钥
*/
func genQuery(appID string, appSecretKey string, param map[string]string) string {

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

	for k, v := range param {
		signer.AddBody(k, v)
	}

	query := signer.GetSignedQuery()

	return query
}

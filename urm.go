package urmclient

const (
	BASE_URL_PRODUCT = "http://pms-api.gmtech.top"
	BASE_URL_DEV     = "http://pms-api-dev.gmtech.top"
)

type URM struct {
	AppID     string
	AppSecret string
	BaseURL   string
}

func New(appID string, appSecret string, baseURL string) *URM {
	return &URM{
		AppID:     appID,
		AppSecret: appSecret,
		BaseURL:   baseURL,
	}
}

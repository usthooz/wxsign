package wxsign

type WxSign struct {
	// Appid 公众号appid
	Appid string
	// AppSecret 公众号秘钥
	AppSecret string
}

// New 创建对象
func New(appid, secret string) *WxSign {
	return &WxSign{
		Appid:     appid,
		AppSecret: secret,
	}
}

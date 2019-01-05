package wxsign

type WxSign struct {
	// Appid 公众号appid
	Appid string
	// AppSecret 公众号秘钥
	AppSecret string
	// TokenRdsKey access_token缓存key
	TokenRdsKey string
	// TicketRdsKey ticket缓存key
	TicketRdsKey string
}

// New 创建对象
func New(appid, secret, tokenKey, ticketKey string) *WxSign {
	return &WxSign{
		Appid:        appid,
		AppSecret:    secret,
		TokenRdsKey:  tokenKey,
		TicketRdsKey: ticketKey,
	}
}

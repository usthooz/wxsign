package wxsign

const (
	// APIURLPrefix 微信授权请求
	WxAPIURLPrefix = "https://api.weixin.qq.com/cgi-bin"
	// AuthURL get access_token route
	WxAuthURL = "/token?grant_type=client_credential&"
	// GetTicketURL get ticket route
	WxGetTicketURL = "/ticket/getticket?"

	// TokenExpire token缓存的时间
	TokenExpire = 3600
	// TicketExpire tocket缓存时间
	TicketExpire = 3600
)

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

// WxJsSign
type WxJsSign struct {
	Appid     string `json:"appid"`
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
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

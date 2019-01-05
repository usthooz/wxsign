package wxsign

import (
	"time"

	redis "gopkg.in/redis.v3"
)

var (
	rdsCli *redis.Client
)

// Init 初始化 redis client
func WxSignRdsInit(rc *redis.Client) {
	if rdsCli == nil {
		rdsCli = rc
	}
}

// PushToken 将微信token 存到 redis 中
func (wSign *WxSign) PushTokenByCache(token string, duration time.Duration) {
	rdsCli.Set(wSign.TokenRdsKey, token, duration)
}

// PushTicket 将微信jsticket 存到 redis 中
func (wSign *WxSign) PushTicketByCache(token string, duration time.Duration) {
	rdsCli.Set(wSign.TicketRdsKey, token, duration)
}

// GetTokenByCache 从缓存获取access_token
func (wSign *WxSign) GetTokenByCache() string {
	var (
		data string
	)
	if rdsCli.Exists(wSign.TokenRdsKey).Val() {
		data, _ = rdsCli.Get(wSign.TokenRdsKey).Result()
	}
	return data
}

// GetTicketByCache 从缓存获取ticket
func (wSign *WxSign) GetTicketByCache() string {
	var (
		data string
	)
	if rdsCli.Exists(wSign.TicketRdsKey).Val() {
		data, _ = rdsCli.Get(wSign.TicketRdsKey).Result()
	}
	return data
}

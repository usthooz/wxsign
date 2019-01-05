package wxsign

import (
	"fmt"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/usthooz/gutil/http"
)

// GetAccessToken 获取普通api调用需要的access_token 因为有次数限制，需要缓存
func (wSign *WxSign) GetAccessToken() (accessToken string, err error) {
	// 首先从redis缓存中取出token
	accessToken = wSign.GetTokenByCache()
	if len(accessToken) > 0 {
		return accessToken, nil
	}
	/*
		Response:
		{
			"access_token":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8", // token
			"expires_in":7200 // 时效
		}
	*/
	url := fmt.Sprintf("%s%sappid=%s&secret=%s", WxAPIURLPrefix, WxAuthURL, wSign.Appid, wSign.AppSecret)
	_, bs, e := xhttp.Get(url)
	if e != nil {
		err = fmt.Errorf("GetAccessToken: Request Failed, err-> %v", e)
		return
	}
	var (
		resp *simplejson.Json
	)
	resp, e = simplejson.NewJson(bs)
	if e != nil {
		err = fmt.Errorf("GetAccessToken: Unmarshal Failed, err-> %v", err)
		return
	}
	if _, ok := resp.CheckGet("errcode"); ok {
		err = fmt.Errorf("GetAccessToken: get access token err-> %s", string(bs))
		return
	}
	expire := resp.GetPath("expires_in").MustInt()
	if expire < 1 {
		expire = TokenExpire
	} else {
		expire = expire - 100
	}
	wSign.PushTokenByCache(resp.GetPath("access_token").MustString(), time.Duration(expire)*time.Second)
	return
}

// GetTicket 获取JSAPI授权TICKET
func (wSign *WxSign) GetTicket() (ticket string, err error) {
	// 首先从缓存获取
	ticket = wSign.GetTicketByCache()
	if len(ticket) > 0 {
		return ticket, nil
	}
	//重新获取ticket
	accessToken, e := wSign.GetAccessToken()
	if e != nil {
		err = e
		return
	}
	url := fmt.Sprintf("%s%saccess_token=%s&type=jsapi", WxAPIURLPrefix, WxGetTicketURL, accessToken)
	_, bs, e := xhttp.Get(url)
	if e != nil {
		err = fmt.Errorf("GetJsTicket: get ticket err-> %v", e)
		return
	}
	var (
		resp *simplejson.Json
	)
	resp, e = simplejson.NewJson(bs)
	if e != nil {
		err = fmt.Errorf("GetJsTicket: Unmarshal err-> %v", err)
		return
	}
	/*
		Response:
		{
			"errcode":0,
			"errmsg":"ok",
			"ticket":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8",
			"expires_in":7200
		}
	*/
	if _, ok := resp.CheckGet("ticket"); !ok {
		err = fmt.Errorf("GetJsTicket: get ticket err-> %s", string(bs))
		return
	}
	expire := resp.GetPath("expires_in").MustInt()
	if expire < 1 {
		expire = TicketExpire
	} else {
		expire = expire - 100
	}
	wSign.PushTicketByCache(resp.GetPath("ticket").MustString(), time.Duration(expire)*time.Second)
	return
}

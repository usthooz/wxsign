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

	accessToken = resp.GetPath("access_token").MustString()
	expire := resp.GetPath("expires_in").MustInt()
	if expire < 1 {
		expire = TokenExpire
	} else {
		expire = expire - 100
	}
	wSign.PushTokenByCache(accessToken, time.Duration(expire)*time.Second)
	return
}

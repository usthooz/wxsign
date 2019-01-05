package wxsign

import (
	"crypto/sha1"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/usthooz/gutil"
)

// GetJsSign GetJsSign
func (wSign *WxSign) GetJsSign(url string) (*WxJsSign, error) {
	jsTicket, err := wSign.GetTicket()
	if err != nil {
		return nil, err
	}
	// splite url
	urlSlice := strings.Split(url, "#")
	jsSign := &WxJsSign{
		Appid:     wSign.Appid,
		Noncestr:  gutil.RandString(16),
		Timestamp: strconv.FormatInt(time.Now().UTC().Unix(), 10),
		Url:       urlSlice[0],
	}
	jsSign.Signature = Signature(jsTicket, jsSign.Noncestr, jsSign.Timestamp, jsSign.Url)
	return jsSign, nil
}

// Signature
func Signature(jsTicket, noncestr, timestamp, url string) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsTicket, noncestr, timestamp, url)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

package wechatoffice

import (
	"context"
	"crypto/sha1"
	"fmt"
	"go.dtapp.net/gorandom"
	"io"
	"time"
)

type ShareResponse struct {
	AppId     string `json:"app_id"`
	NonceStr  string `json:"nonce_str"`
	Timestamp int64  `json:"timestamp"`
	Url       string `json:"url"`
	RawString string `json:"raw_string"`
	Signature string `json:"signature"`
}

type ShareResult struct {
	Result ShareResponse // 结果
	Err    error         // 错误
}

func newShareResult(result ShareResponse, err error) *ShareResult {
	return &ShareResult{Result: result, Err: err}
}

func (c *Client) Share(ctx context.Context, url string) *ShareResult {
	c.getAccessToken(ctx)
	c.config.jsapiTicket = c.GetJsapiTicket(ctx)
	var response ShareResponse
	response.AppId = c.GetAppId()
	response.NonceStr = gorandom.Alphanumeric(32)
	response.Timestamp = time.Now().Unix()
	response.Url = url
	response.RawString = fmt.Sprintf("jsapi_ticket=%v&noncestr=%v&timestamp=%v&url=%v", c.config.jsapiTicket, response.NonceStr, response.Timestamp, response.Url)
	t := sha1.New()
	_, err := io.WriteString(t, response.RawString)
	response.Signature = fmt.Sprintf("%x", t.Sum(nil))
	return newShareResult(response, err)
}

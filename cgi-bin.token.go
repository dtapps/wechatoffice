package wechatoffice

import (
	"context"
	"fmt"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type CgiBinTokenResponse struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode     int    `json:"errcode"`      // 错误码
	Errmsg      string `json:"errmsg"`       // 错误信息
}

type CgiBinTokenResult struct {
	Result CgiBinTokenResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newCgiBinTokenResult(result CgiBinTokenResponse, body []byte, http gorequest.Response) *CgiBinTokenResult {
	return &CgiBinTokenResult{Result: result, Body: body, Http: http}
}

// CgiBinToken
// 接口调用凭证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/access-token/auth.getAccessToken.html
func (c *Client) CgiBinToken(ctx context.Context, notMustParams ...gorequest.Params) (*CgiBinTokenResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", c.GetAppId(), c.GetAppSecret()), params, http.MethodGet)
	if err != nil {
		return newCgiBinTokenResult(CgiBinTokenResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinTokenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinTokenResult(response, request.ResponseBody, request), err
}

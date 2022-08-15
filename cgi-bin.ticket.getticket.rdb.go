package wechatoffice

import (
	"context"
	"fmt"
	"time"
)

// GetJsapiTicket 获取api_ticket
func (c *Client) GetJsapiTicket(ctx context.Context) string {
	if c.redisClient.Db == nil {
		return c.config.JsapiTicket
	}
	newCache := c.redisClient.NewSimpleStringCache(c.redisClient.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := c.CgiBinTicketGetTicket(ctx, "jsapi")
		return token.Result.Ticket
	}
	return newCache.GetCache(ctx, c.getJsapiTicketCacheKeyName())
}

func (c *Client) getJsapiTicketCacheKeyName() string {
	return fmt.Sprintf("wechat_jsapi_ticket:%v", c.GetAppId())
}

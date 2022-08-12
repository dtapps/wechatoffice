package wechatoffice

import (
	"context"
	"fmt"
	"time"
)

// GetJsapiTicket 获取api_ticket
func (c *Client) GetJsapiTicket(ctx context.Context) string {
	if c.config.RedisClient.Db == nil {
		return c.config.JsapiTicket
	}
	newCache := c.config.RedisClient.NewSimpleStringCache(c.config.RedisClient.NewStringOperation(), time.Second*7000)
	newCache.DBGetter = func() string {
		token := c.CgiBinTicketGetTicket(ctx, "jsapi")
		return token.Result.Ticket
	}
	return newCache.GetCache(c.getJsapiTicketCacheKeyName())
}

func (c *Client) getJsapiTicketCacheKeyName() string {
	return fmt.Sprintf("wechat_jsapi_ticket:%v", c.GetAppId())
}

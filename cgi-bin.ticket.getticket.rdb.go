package wechatoffice

import (
	"context"
	"time"
)

// GetJsapiTicket 获取api_ticket
func (c *Client) GetJsapiTicket(ctx context.Context) string {
	if c.cache.redisClient == nil {
		return c.config.jsapiTicket
	}
	result, _ := c.cache.redisClient.Get(ctx, c.getJsapiTicketCacheKeyName()).Result()
	if result != "" {
		return result
	}
	token, _ := c.CgiBinTicketGetTicket(ctx, "jsapi")
	c.cache.redisClient.Set(ctx, c.getJsapiTicketCacheKeyName(), token.Result.Ticket, time.Second*7000)
	return token.Result.Ticket
}

func (c *Client) getJsapiTicketCacheKeyName() string {
	return c.cache.wechatJsapiTicketPrefix + c.GetAppId()
}

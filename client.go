package wechatoffice

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

type ConfigClient struct {
	AppId       string            // 小程序唯一凭证，即 appId
	AppSecret   string            // 小程序唯一凭证密钥，即 appSecret
	AccessToken string            // 接口调用凭证
	JsapiTicket string            // 签名凭证
	RedisClient *dorm.RedisClient // 缓存数据库
	GormClient  *dorm.GormClient  // 日志数据库
	LogClient   *golog.ZapLog     // 日志驱动
	LogDebug    bool              // 日志开关
}

// Client 微信公众号服务
type Client struct {
	requestClient *gorequest.App    // 请求服务
	logClient     *golog.ApiClient  // 日志服务
	redisClient   *dorm.RedisClient // 缓存服务
	config        *ConfigClient     // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.requestClient = gorequest.NewHttp()

	if c.config.GormClient.Db != nil {
		c.logClient, err = golog.NewApiClient(&golog.ApiClientConfig{
			GormClient: c.config.GormClient,
			TableName:  logTable,
			LogClient:  c.config.LogClient,
			LogDebug:   c.config.LogDebug,
		})
		if err != nil {
			return nil, err
		}
	}

	c.redisClient = c.config.RedisClient

	return c, nil
}

// ConfigApp 配置
func (c *Client) ConfigApp(appId, appSecret string) *Client {
	c.config.AppId = appId
	c.config.AppSecret = appSecret
	return c
}

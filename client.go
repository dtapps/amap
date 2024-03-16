package amap

import (
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	Key string
}

// Client 实例
type Client struct {
	config struct {
		key string
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(key string) (*Client, error) {
	c := &Client{}
	c.config.key = key
	return c, nil
}

package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// Config 配置结构体
type Config struct {
	MySQL struct {
		DSN          string `json:"dsn"`
		MaxIdleConns int    `json:"max_idle_conns"`
		MaxOpenConns int    `json:"max_open_conns"`
	} `json:"mysql"`

	Redis struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Password string `json:"password"`
		DB       int    `json:"db"`
		PoolSize int    `json:"pool_size"`
	} `json:"redis"`

	RocketMQ struct {
		NameServer string `json:"name_server"`
		GroupName  string `json:"group_name"`
		Topic      struct {
			Seckill string `json:"seckill"`
		} `json:"topic"`
	} `json:"rocketmq"`

	Server struct {
		Port int `json:"port"`
	} `json:"server"`

	Limiter struct {
		RatePerSecond int `json:"rate_per_second"` // 每秒限流数
		Burst         int `json:"burst"`           // 突发流量处理能力
	} `json:"limiter"`

	Seckill struct {
		OrderTTL        int `json:"order_ttl"`        // 订单过期时间(秒)
		CheckInterval   int `json:"check_interval"`   // 订单检查间隔(秒)
		MaxRetryCount   int `json:"max_retry_count"`  // 最大重试次数
		ConcurrentLimit int `json:"concurrent_limit"` // 并发限制
	} `json:"seckill"`

	Monitor struct {
		LogLevel string `json:"log_level"`
	} `json:"monitor"`
}

var (
	_config *Config
	once    sync.Once
)

// Init 初始化配置
func Init(configPath string) {
	once.Do(func() {
		var err error

		// 如果没有提供配置路径，使用默认路径
		if configPath == "" {
			configPath = "config.json"
			// 检查当前目录
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				// 检查上层目录的config目录
				parentPath := filepath.Join("..", "config", "config.json")
				if _, err = os.Stat(parentPath); err == nil {
					configPath = parentPath
				}
			}
		}

		// 读取配置文件
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic("failed to read config file: " + err.Error())
		}

		// 解析配置
		_config = &Config{}
		if err = json.Unmarshal(data, _config); err != nil {
			panic("failed to parse config file: " + err.Error())
		}
	})
}

// GetConf 获取配置
func GetConf() *Config {
	return _config
}

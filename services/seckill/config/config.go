package config

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	conf     *Config
	confOnce sync.Once
)

// Config 配置结构体
type Config struct {
	Service  ServiceConfig  `json:"service"`
	MySQL    MySQLConfig    `json:"mysql"`
	Redis    RedisConfig    `json:"redis"`
	RocketMQ RocketMQConfig `json:"rocketmq"`
	Etcd     EtcdConfig     `json:"etcd"`
	Business BusinessConfig `json:"business"`
	Jaeger   JaegerConfig   `json:"jaeger"`
	Seckill  SeckillConfig  `json:"seckill"`
}

// ServiceConfig 服务配置
type ServiceConfig struct {
	Name           string `json:"name"`
	Port           int    `json:"port"`
	MaxConnections int64  `json:"max_connections"`
	MaxQPS         int64  `json:"max_qps"`
}

// MySQLConfig MySQL配置
type MySQLConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	MaxOpen  int    `json:"max_open"`
	MaxIdle  int    `json:"max_idle"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// RocketMQConfig RocketMQ配置
type RocketMQConfig struct {
	Endpoints     []string `json:"endpoints"`
	AccessKey     string   `json:"access_key"`
	SecretKey     string   `json:"secret_key"`
	ConsumerGroup string   `json:"consumer_group"`
	Topic         struct {
		Seckill string `json:"seckill"`
	} `json:"topic"`
}

// EtcdConfig Etcd配置
type EtcdConfig struct {
	Address string `json:"address"`
}

// BusinessConfig 业务配置
type BusinessConfig struct {
	OrderTTL       int64 `json:"order_ttl"`        // 订单超时时间(秒)
	InventoryCheck int64 `json:"inventory_check"`  // 库存检查间隔(秒)
	BloomFilterTTL int64 `json:"bloom_filter_ttl"` // 布隆过滤器超时时间(秒)
	LimitOneBuy    bool  `json:"limit_one_buy"`    // 是否限制每人只能购买一次
}

// JaegerConfig Jaeger配置
type JaegerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// SeckillConfig 秒杀业务配置
type SeckillConfig struct {
	ExpireTime            int64 `json:"expire_time"`             // 秒杀流水过期时间（秒）
	CheckExpiredInterval  int64 `json:"check_expired_interval"`  // 检查过期流水的间隔（秒）
	ConsistencyInterval   int64 `json:"consistency_interval"`    // 一致性检查的间隔（秒）
	PreloadProductsBefore int64 `json:"preload_products_before"` // 预加载商品的提前时间（秒）
}

// Init 初始化配置
func Init(configPath string) error {
	var err error
	confOnce.Do(func() {
		// 读取配置文件
		data, readErr := os.ReadFile(configPath)
		if readErr != nil {
			klog.Errorf("Failed to read config file: %v", readErr)
			err = readErr
			return
		}

		// 解析配置
		conf = &Config{}
		jsonErr := json.Unmarshal(data, conf)
		if jsonErr != nil {
			klog.Errorf("Failed to parse config: %v", jsonErr)
			err = jsonErr
			return
		}

		// 设置默认值
		if conf.Seckill.ExpireTime <= 0 {
			conf.Seckill.ExpireTime = 900 // 默认15分钟
		}
		if conf.Seckill.CheckExpiredInterval <= 0 {
			conf.Seckill.CheckExpiredInterval = 60 // 默认1分钟
		}
		if conf.Seckill.ConsistencyInterval <= 0 {
			conf.Seckill.ConsistencyInterval = 300 // 默认5分钟
		}

		klog.Infof("Config loaded: %+v", conf)
	})
	return err
}

// GetConf 获取配置
func GetConf() *Config {
	if conf == nil {
		klog.Warn("Config not initialized, initializing with default path")
		if err := Init("config/config.json"); err != nil {
			klog.Fatalf("Failed to initialize config: %v", err)
		}
	}
	return conf
}

// LoadConfig 从文件加载配置
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// 设置默认值
	if cfg.Seckill.ExpireTime <= 0 {
		cfg.Seckill.ExpireTime = 900 // 默认15分钟
	}
	if cfg.Seckill.CheckExpiredInterval <= 0 {
		cfg.Seckill.CheckExpiredInterval = 60 // 默认1分钟
	}
	if cfg.Seckill.ConsistencyInterval <= 0 {
		cfg.Seckill.ConsistencyInterval = 300 // 默认5分钟
	}

	conf = &cfg
	return &cfg, nil
}

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"gomall/kitex_gen/seckill/seckillservice"
	"gomall/services/seckill/config"
	"gomall/services/seckill/dal/model"
	"gomall/services/seckill/handler"
	"gomall/services/seckill/initialize"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var configPath = flag.String("config", "", "config file path")

// 初始化日志
func initLogger() {
	// 日志级别
	logLevel := zap.InfoLevel
	if config.GetConf().Monitor.LogLevel == "debug" {
		logLevel = zap.DebugLevel
	}

	// 编码器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 输出配置
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)

	// 创建Logger
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

// 初始化数据库
func initDB() {
	// 初始化MySQL
	initialize.InitMySQL()

	// 自动迁移表结构
	db := initialize.GetMysql()
	err := db.AutoMigrate(&model.SeckillProduct{}, &model.InventoryFlow{})
	if err != nil {
		klog.Fatal("failed to migrate database: ", err)
	}
}

// 初始化Redis
func initRedis() {
	initialize.InitRedis()
}

func main() {
	flag.Parse()

	// 初始化配置
	config.Init(*configPath)

	// 初始化日志
	initLogger()

	// 初始化数据库
	initDB()

	// 初始化Redis
	initRedis()

	// 创建服务实例
	seckillService := handler.NewSeckillServiceImpl()
	kitexHandler := handler.NewKitexSeckillServiceImpl()

	// 创建并启动定时任务
	cronTaskManager := handler.NewCronTaskManager(seckillService)
	cronTaskManager.Start()
	defer cronTaskManager.Stop()

	// 获取服务地址
	addr := fmt.Sprintf(":%d", config.GetConf().Server.Port)

	// 创建RPC服务
	svr := seckillservice.NewServer(
		kitexHandler,
		server.WithServiceAddr(&net.TCPAddr{Port: config.GetConf().Server.Port}),
		server.WithLimit(&limit.Option{
			MaxConnections: 1000,
			MaxQPS:         500,
		}),
		server.WithRegistry(nil), // 可以根据需要设置注册中心
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "seckill",
		}),
	)

	// 处理优雅退出
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		klog.Info("Shutting down server...")

		// 停止定时任务
		cronTaskManager.Stop()

		// 关闭服务
		svr.Stop()

		klog.Info("Server exited")
	}()

	// 启动服务
	klog.Infof("Starting seckill service on %s", addr)
	if err := svr.Run(); err != nil {
		klog.Fatal("Error in running server: ", err)
	}
}

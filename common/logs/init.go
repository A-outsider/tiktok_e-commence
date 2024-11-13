package logs

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

// LogInit 初始化日志系统

const fileMaxAge = 30
const fileMaxBackups = 60
const fileMaxSize = 200
const logfileBasePath = "./data/logs/"

func LogInit(serviceName string) {
	// 设置日志文件路径
	path := filepath.Join(logfileBasePath, serviceName+".log")

	// 确保日志目录存在
	logDir := filepath.Dir(path)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}

	// 配置日志文件切割器
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    fileMaxSize,
		MaxBackups: fileMaxBackups,
		MaxAge:     fileMaxAge,
		Compress:   true,
	})

	// 配置日志编码
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 创建 consoleEncoder 和 JSON Encoder
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// 将日志同时输出到文件和控制台
	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, writeSyncer, zapcore.InfoLevel),                // 文件输出
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.InfoLevel), // 控制台输出
	)

	// 创建 zap logger
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(zapLogger)

	// 使用 zapLogger 创建 hertzZapLogger
	hertzZapLogger := hertzzap.NewLogger(hertzzap.WithZapOptions(zap.AddCaller()))

	//hertzZapLogger.SetOutput(writeSyncer)                // 设置输出到日志中

	hlog.SetLogger(hertzZapLogger) // 替换 hlog 的默认 logger 为 Hertz zap logger

	_ = zapLogger.Sync()
	return
}

// AccessLog 是一个记录访问日志的中间件，类似于 gin.Logger
func AccessLog() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		c.Next(ctx) // 执行请求

		// 记录访问日志
		latency := time.Since(start).Microseconds()
		hlog.CtxInfof(ctx, "status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s",
			c.Response.StatusCode(), latency,
			c.Request.Header.Method(), c.Request.URI().PathOriginal(), c.ClientIP(), c.Request.Host())
	}
}

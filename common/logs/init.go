package logs

import (
	"github.com/natefinch/lumberjack"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gomall/common/config"
	"os"
	"path/filepath"
)

// LogInit 初始化日志系统
func LogInit() {
	// 设置日志文件路径
	cfg := config.Common.Log

	path := filepath.Join(cfg.BasePath, cfg.Name+".log")

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

	// 创建 zap core
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writeSyncer, zapcore.InfoLevel),
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	)

	// 创建 zap logger
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer zapLogger.Sync()

	// 使用 zapLogger 创建 otelzap logger
	otelZapLogger := otelzap.New(zapLogger, otelzap.WithStackTrace(true))
	otelzap.ReplaceGlobals(otelZapLogger) // 将其设置为全局的 otelzap logger

	return
}

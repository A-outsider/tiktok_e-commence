package logs

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
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
	hertzZapLogger := hertzzap.NewLogger(
		hertzzap.WithCoreEnc(jsonEncoder),
		hertzzap.WithZapOptions(zap.AddCaller(), zap.AddCallerSkip(1)),
	)

	//hertzZapLogger.SetOutput(writeSyncer) // 设置输出到日志中

	hlog.SetLogger(hertzZapLogger) // 替换 hlog 的默认 logger 为 Hertz zap logger

	_ = zapLogger.Sync()
	return
}

package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

var MediaLogger *zap.Logger

var ForwardLogger *zap.Logger

// InitLogger 初始化日志
func InitLogger() {
	initRootLogger()
	initMediaLogger()
	initForwardLogger()
}

func initRootLogger() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "app.log", // 日志文件路径
		MaxSize:    1024,      // 单个日志文件的最大大小（单位：MB）
		MaxBackups: 30,        // 保留的备份文件最大数量
		MaxAge:     7,         // 日志文件的最大保存天数（单位：天）
		Compress:   false,     // 是否压缩归档的旧日志文件（gzip）
		LocalTime:  true,      // 归档文件名是否使用本地时间（默认是UTC时间）
	}
	defer lumberjackLogger.Close()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.CallerKey = "caller"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	level := zapcore.InfoLevel

	core := zapcore.NewCore(encoder, zapcore.AddSync(lumberjackLogger), level)

	zl := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(zl)

	Logger = zl
	defer zl.Sync()
}

func initMediaLogger() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "media.log", // 日志文件路径
		MaxSize:    1024,        // 单个日志文件的最大大小（单位：MB）
		MaxBackups: 30,          // 保留的备份文件最大数量
		MaxAge:     7,           // 日志文件的最大保存天数（单位：天）
		Compress:   false,       // 是否压缩归档的旧日志文件（gzip）
		LocalTime:  true,        // 归档文件名是否使用本地时间（默认是UTC时间）
	}
	defer lumberjackLogger.Close()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	level := zapcore.InfoLevel

	core := zapcore.NewCore(encoder, zapcore.AddSync(lumberjackLogger), level)

	zl := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(zl)

	MediaLogger = zl
	defer zl.Sync()
}

func initForwardLogger() {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "media.log", // 日志文件路径
		MaxSize:    1024,        // 单个日志文件的最大大小（单位：MB）
		MaxBackups: 30,          // 保留的备份文件最大数量
		MaxAge:     7,           // 日志文件的最大保存天数（单位：天）
		Compress:   false,       // 是否压缩归档的旧日志文件（gzip）
		LocalTime:  true,        // 归档文件名是否使用本地时间（默认是UTC时间）
	}
	defer lumberjackLogger.Close()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	level := zapcore.InfoLevel

	core := zapcore.NewCore(encoder, zapcore.AddSync(lumberjackLogger), level)

	zl := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zap.ReplaceGlobals(zl)

	ForwardLogger = zl
	defer zl.Sync()
}

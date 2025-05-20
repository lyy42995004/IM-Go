package zap

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once    sync.Once
	Logger  *zap.Logger
	String  = zap.String
	Any     = zap.Any
	Err     = zap.Error
	Int     = zap.Int
	Float32 = zap.Float32
)

const (
	defaultLogPath    = "./logs"
	defaultLogLevel   = "info"
	defaultMaxSize    = 100  // MB
	defaultMaxBackups = 30   // 保留的旧日志文件数量
	defaultMaxAge     = 7    // 保留天数
	defaultCompress   = true // 是否压缩
)

// Init 初始化日志记录器(单例模式)
func Init() *zap.Logger {
	once.Do(func() {
		// 确保日志目录存在
		if err := os.MkdirAll(defaultLogPath, 0755); err != nil {
			panic(err)
		}

		// 设置日志级别
		level := getLogLevel(defaultLogLevel)

		// 日志文件路径
		logFile := getDatedLogFilename(defaultLogPath)

		// 设置日志轮转
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    defaultMaxSize,    // MB
			MaxBackups: defaultMaxBackups, // 保留的旧日志文件数量
			MaxAge:     defaultMaxAge,     // 保留天数
			Compress:   defaultCompress,   // 是否压缩
		})

		// 编码器配置
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 不带颜色

		// 核心配置
		core := zapcore.NewCore(
			// zapcore.NewJSONEncoder(encoderConfig), // json格式
			zapcore.NewConsoleEncoder(encoderConfig), // 使用 Console 编码器
			writer,
			level,
		)

		// 创建Logger
		Logger = zap.New(core)
	})

	return Logger
}

// GetLogger 获取日志记录器实例
func GetLogger() *zap.Logger {
	if Logger == nil {
		Init() // 自动初始化
	}
	return Logger
}

// getDatedLogFilename 生成带日期的日志文件名
func getDatedLogFilename(basePath string) string {
	now := time.Now()
	return filepath.Join(basePath, fmt.Sprintf("chat_%s.log", now.Format("2006-01-02")))
}

// getLogLevel 将字符串日志级别转换为zapcore.Level
func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// 直接可用的日志方法
func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
}

func Sync() error {
	return GetLogger().Sync()
}

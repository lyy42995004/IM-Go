package zap

import (
	"os"
	"path"
	"runtime"

	"github.com/lyy42995004/IM-Go/internal/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var logPath string

// 自动调用
func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 创建JSON编码器
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 获取配置
	conf := config.GetConfig()
	logPath = conf.LogPath

	// 创建文件写入同步器
	fileWriteSyncer := getFileLogWriter()

	// 创建核心日志处理器
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(encoder, fileWriteSyncer, zapcore.DebugLevel),
	)
	logger = zap.New(core)
}

// 创建了一个支持日志轮转的写入器
func getFileLogWriter() (writeSyncer zapcore.WriteSyncer) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    100,   // 单个文件最大100MB
		MaxBackups: 60,    // 最多保留60个备份文件
		MaxAge:     7,     // 日志文件最多保留7天
		Compress:   false, // 不压缩旧日志
	}

	return zapcore.AddSync(lumberJackLogger)
}

// 获得调用方的日志信息，包括函数名，文件名，行号
func getCallerInfoForLog() []zap.Field {
	// 获取调用者的程序计数器、文件名和行号(跳过2层调用栈)
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return nil
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) // 只保留函数名

	// 创建包含调用者信息的字段
	return []zap.Field{
		zap.String("func", funcName),
		zap.String("file", file),
		zap.Int("line", line),
	}
}

func Info(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()    // 获取调用者信息
	fields = append(fields, callerFields...) // 合并字段
	logger.Info(message, fields...)          // 记录日志
}

func Warn(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Fatal(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Debug(message, fields...)
}

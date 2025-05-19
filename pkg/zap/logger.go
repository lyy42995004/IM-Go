package zap

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lyy42995004/IM-Go/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once     sync.Once
	std      *Logger
	logPath  string
	initDone bool
)

type Level = zapcore.Level

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

type Logger struct {
	l  *zap.Logger
	al *zap.AtomicLevel
}

type RotateConfig struct {
	Filename     string        // 完整文件名
	MaxAge       int           // 保留旧日志文件的最大天数
	RotationTime time.Duration // 日志文件轮转时间
	MaxSize      int           // 日志文件最大大小（MB）
	MaxBackups   int           // 保留日志文件的最大数量
	Compress     bool          // 是否对日志文件进行压缩归档
	LocalTime    bool          // 是否使用本地时间，默认 UTC 时间
}

func initLogger() {
	conf := config.GetConfig()
	logPath = filepath.Join(conf.LogConfig.LogPath, "chat.log")

	// 确保日志目录存在
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		panic(err)
	}

	std = NewWithRotate(InfoLevel, NewProductionRotateConfig(logPath))
	initDone = true
}

// Default 获取默认Logger，确保只初始化一次
func Default() *Logger {
	once.Do(initLogger)
	return std
}

// NewWithRotate 创建一个带日志轮转功能的Logger
func NewWithRotate(level Level, cfg *RotateConfig, opts ...zap.Option) *Logger {
	var writer io.Writer
	if cfg.MaxSize > 0 {
		writer = NewRotateBySize(cfg)
	} else {
		writer = NewRotateByTime(cfg)
	}

	al := zap.NewAtomicLevelAt(level)
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(writer),
		al,
	)

	return &Logger{
		l:  zap.New(core, opts...),
		al: &al,
	}
}

func NewProductionRotateConfig(filename string) *RotateConfig {
	return &RotateConfig{
		Filename:     filename,
		MaxAge:       30,             // 日志保留30天
		RotationTime: time.Hour * 24, // 24小时轮转一次
		MaxSize:      100,            // 100M
		MaxBackups:   100,
		Compress:     true,
		LocalTime:    false,
	}
}

func NewRotateByTime(cfg *RotateConfig) io.Writer {
	opts := []rotatelogs.Option{
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge) * time.Hour * 24),
		rotatelogs.WithRotationTime(cfg.RotationTime),
		rotatelogs.WithLinkName(cfg.Filename),
	}
	if !cfg.LocalTime {
		opts = append(opts, rotatelogs.WithClock(rotatelogs.UTC))
	}
	filename := strings.SplitN(cfg.Filename, ".", 2)
	l, _ := rotatelogs.New(
		filename[0]+".%Y-%m-%d-%H-%M-%S."+filename[1],
		opts...,
	)
	return l
}

func NewRotateBySize(cfg *RotateConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  cfg.LocalTime,
		Compress:   cfg.Compress,
	}
}

func (l *Logger) SetLevel(level Level) {
	if l.al != nil {
		l.al.SetLevel(level)
	}
}

type Field = zap.Field

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

func ReplaceDefault(l *Logger) { std = l }

func SetLevel(level Level) { Default().SetLevel(level) }

func Debug(msg string, fields ...Field) { Default().Debug(msg, fields...) }
func Info(msg string, fields ...Field)  { Default().Info(msg, fields...) }
func Warn(msg string, fields ...Field)  { Default().Warn(msg, fields...) }
func Error(msg string, fields ...Field) { Default().Error(msg, fields...) }
func Panic(msg string, fields ...Field) { Default().Panic(msg, fields...) }
func Fatal(msg string, fields ...Field) { Default().Fatal(msg, fields...) }

func Sync() error { return Default().Sync() }

// SetLogPath 设置日志路径并重新初始化默认Logger
func SetLogPath(path string) {
	logPath = filepath.Join(path, "app.log")
	// 确保日志目录存在
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		panic(err)
	}

	if initDone {
		ReplaceDefault(NewWithRotate(InfoLevel, NewProductionRotateConfig(logPath)))
	}
}

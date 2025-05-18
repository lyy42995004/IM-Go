# Zap Logger 文档

## 概述

这是一个基于 `go.uber.org/zap` 的高性能日志库封装，提供了日志轮转、动态级别调整等增强功能。

**功能特性**

• 高性能结构化日志记录

• 自动日志文件轮转（按时间/大小）

• 动态日志级别调整

• 线程安全的单例模式

• 多级别日志支持（Debug 到 Fatal）

• 自动创建日志目录


## 快速开始

**基本使用**

```go
import log "github.com/yourproject/zap"

func main() {
    // 记录日志（自动初始化）
    log.Info("服务启动")
    log.Error("操作失败", log.String("error", "连接超时"))
}
```

**配置示例**

在项目配置文件中（如 `config.toml`）:

```toml
[logConfig]
logPath = "logs/"  # 日志目录
```

## API 文档

**类型定义**

`Level`

日志级别类型，等同于 `zapcore.Level`:

```go
type Level = zapcore.Level
```

可用级别常量:

```go
const (
    DebugLevel = zapcore.DebugLevel
    InfoLevel  = zapcore.InfoLevel
    WarnLevel  = zapcore.WarnLevel
    ErrorLevel = zapcore.ErrorLevel
    PanicLevel = zapcore.PanicLevel
    FatalLevel = zapcore.FatalLevel
)
```

`Logger`

核心日志记录器:

```go
type Logger struct {
    l  *log.Logger      // 底层 zap logger
    al *log.AtomicLevel // 原子级别控制
}
```

`RotateConfig`

日志轮转配置:

```go
type RotateConfig struct {
    Filename     string        // 完整文件名
    MaxAge       int           // 保留天数
    RotationTime time.Duration // 轮转时间间隔
    MaxSize      int           // 文件最大大小(MB)
    MaxBackups   int           // 最大备份数
    Compress     bool          // 是否压缩
    LocalTime    bool          // 是否使用本地时间
}
```

**核心函数**

`Default() *Logger`

获取默认 logger 实例（线程安全单例）:

```go
logger := log.Default()
```

`NewWithRotate(level Level, cfg *RotateConfig, opts ...log.Option) *Logger`

创建带轮转功能的 logger:

```go
logger := log.NewWithRotate(
    log.InfoLevel,
    log.NewProductionRotateConfig("app.log"),
)
```

`NewProductionRotateConfig(filename string) *RotateConfig`

获取生产环境默认轮转配置:

```go
cfg := log.NewProductionRotateConfig("app.log")
// 等效于:
// &RotateConfig{
//     Filename:     filename,
//     MaxAge:       30,
//     RotationTime: 24 * time.Hour,
//     MaxSize:      100,
//     MaxBackups:   100,
//     Compress:     true,
//     LocalTime:    false,
// }
```

**日志方法**

级别方法

```go
func (l *Logger) Debug(msg string, fields ...Field)
func (l *Logger) Info(msg string, fields ...Field)
func (l *Logger) Warn(msg string, fields ...Field)
func (l *Logger) Error(msg string, fields ...Field)
func (l *Logger) Panic(msg string, fields ...Field)
func (l *Logger) Fatal(msg string, fields ...Field)
```

对应的全局快捷方法:

```go
log.Debug("debug message")
log.Info("info message")
// ...其他级别
```

其他方法

```go
func (l *Logger) SetLevel(level Level)  // 动态设置日志级别
func (l *Logger) Sync() error           // 同步日志到磁盘
```

**工具函数**

`SetLogPath(path string)`

动态修改日志路径:

```go
log.SetLogPath("/new/log/path")
```

`SetLevel(level Level)`

全局修改日志级别:

```go
log.SetLevel(log.DebugLevel)
```

示例

基本日志记录

```go
log.Info("用户登录",
    log.String("username", "张三"),
    log.Int("attempt", 3),
)
```

错误处理

```go
if err := doSomething(); err != nil {
    log.Error("操作失败",
        log.String("module", "payment"),
        log.Err(err), // 使用 log.Err 记录错误
    )
}
```

动态调整

```go
// 开发环境使用 Debug 级别
if os.Getenv("ENV") == "dev" {
    log.SetLevel(log.DebugLevel)
}

// 修改日志路径
log.SetLogPath("/tmp/logs")
```

**日志轮转策略**

按时间轮转

默认配置为每天轮转一次，保留30天:

```go
cfg := &log.RotateConfig{
    Filename:     "app.log",
    RotationTime: 24 * time.Hour,
    MaxAge:       30,
}
```

按大小轮转

当文件达到100MB时轮转:

```go
cfg := &log.RotateConfig{
    Filename:   "app.log",
    MaxSize:    100,
    MaxBackups: 10,
    Compress:   true,
}
```

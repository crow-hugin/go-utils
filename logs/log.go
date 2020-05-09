package logs

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

const (
	loggerModelDev string = "dev"
	loggerModelPro string = "pro"
	defaultPath    string = "logs/log"
	defaultMaxAge  int    = 7
	defaultDelay   int    = 24
)

var (
	Notes  *zap.Logger
	writer io.Writer
)

type Config struct {
	Path   string // 日志文件路径
	MaxAge int    // 保留多少天内的日志文件
	Delay  int    //
}

func New(config Config, model string) {
	initializeOption(&config)
	// 默认输出到控制台
	writer = os.Stdout
	//  如果是生产模式并且日志文件路径不为空则将日志打印至文件
	if model == loggerModelPro && config.Path != "" {
		writer = hook(
			config.Path,
			time.Hour*24*time.Duration(config.MaxAge),
			time.Hour*time.Duration(config.Delay))
	}
	// 设置默认日志级别
	atomicLevel := zap.NewAtomicLevel()
	if model == loggerModelDev {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	// 初始化log
	Notes = zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder()),
			zapcore.AddSync(writer), atomicLevel),
		zap.AddCaller())
}

func initializeOption(config *Config) {
	if config.Path == "" {
		config.Path = defaultPath
	}
	if config.MaxAge == 0 {
		config.MaxAge = defaultMaxAge
	}
	if config.Delay == 0 {
		config.Delay = defaultDelay
	}
}

func hook(path string, maxAge time.Duration, delay time.Duration) *rotatelogs.RotateLogs {
	if r, err := rotatelogs.New(
		path+".%Y%m%d%H"+".log",
		rotatelogs.WithLinkName(path+".log"),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(delay),
	); err == nil {
		return r
	} else {
		panic(err)
	}
}

func encoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

// 时间解码
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

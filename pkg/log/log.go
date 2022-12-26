package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	FilePath       string
	FileMaxSize    int
	FileMaxBackups int
	FileMaxAge     int
}

const TimeFormat = "2006-01-02 15:04:05.000"

var logger *zap.Logger

func Init(conf *Config) {
	// 对日志进行分隔
	lumberJackLogger := &lumberjack.Logger{
		Filename:   conf.FilePath,
		MaxSize:    conf.FileMaxSize,
		MaxBackups: conf.FileMaxBackups,
		MaxAge:     conf.FileMaxAge,
		Compress:   false,
		LocalTime:  true,
	}
	// 设置 zap core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller_line", // 打印文件名和行数
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg", // 忽略第一个msg字段
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     customTimeEncoder, // 自定义时间格式
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   customCallerEncoder, // 全路径编码器
		}),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(lumberJackLogger),
		),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)
	logger = zap.New(core)
}

// 自定义时间输出格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(TimeFormat))
}

// 自定义文件：行号输出项
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(caller.TrimmedPath())
}

func Error(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func Debug(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

func Warn(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

func Info(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Close() {
	logger.Sync()
}

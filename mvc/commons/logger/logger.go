package logger

import (
	"os"
	"strings"
	"sync"

	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once sync.Once
	log  *zap.Logger
)

type Fields map[string]interface{}

func lazyInit() {
	once.Do(func() {
		var syncWriter zapcore.WriteSyncer
		config := zap.NewProductionConfig()

		// setup logger level
		config.Level = zap.NewAtomicLevelAt(func(level string) zapcore.Level {
			switch level {
			case "debug":
				return zapcore.DebugLevel
			case "info":
				return zapcore.InfoLevel
			case "warn":
				return zapcore.WarnLevel
			case "error":
				return zapcore.ErrorLevel
			default:
				return zapcore.InfoLevel
			}
		}(strings.ToLower(Config.Level)))

		// setup logger encodeer config
		config.EncoderConfig = zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		if len(Config.FilePath) > 0 {
			syncWriter = zapcore.AddSync(&lumberjack.Logger{
				Filename:   Config.FilePath,
				MaxSize:    Config.MaxSize, // megabytes
				MaxBackups: Config.MaxBackups,
				MaxAge:     Config.MaxAge,   //days
				Compress:   Config.Compress, // disabled by default
			})

			// setup log rotate timer
			timer := cron.New()
			timer.AddFunc("*/1 * * * * *", func() {
				syncWriter.Sync()
			})
			timer.Start()

			// override close function
			copyCloseFunc := Close
			Close = func() {
				copyCloseFunc()
				timer.Stop()
			}
		} else {
			syncWriter = os.Stdout
		}
		// setup logger
		log = zap.New(
			zapcore.NewCore(
				zapcore.NewJSONEncoder(config.EncoderConfig),
				syncWriter,
				config.Level,
			),
		)
	})
}

func Debug(msg string, keyvals map[string]interface{}) {
	lazyInit()
	log.Debug(msg, convertToFields(&keyvals, nil)...)
}

func Info(msg string, keyvals map[string]interface{}) {
	lazyInit()
	log.Info(msg, convertToFields(&keyvals, nil)...)
}

func Warn(msg string, keyvals map[string]interface{}, cause error) {
	lazyInit()
	log.Warn(msg, convertToFields(&keyvals, cause)...)
}

func Error(msg string, keyvals map[string]interface{}, cause error) {
	lazyInit()
	log.Error(msg, convertToFields(&keyvals, cause)...)
}

func Fatal(msg string, keyvals map[string]interface{}, cause error) {
	lazyInit()
	log.Fatal(msg, convertToFields(&keyvals, cause)...)
}

func convertToFields(keyvals *map[string]interface{}, err error) (fields []zapcore.Field) {
	for k, v := range *keyvals {
		fields = append(fields, zap.Any(k, v))
	}
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	return
}

var Close = func() {
	if log != nil {
		log.Sync()
	}
}

package applog

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() *zap.SugaredLogger {
	loggerLevelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}

	logLevel := viper.GetString("log_level")
	level, exist := loggerLevelMap[logLevel]
	if !exist {
		level = zapcore.DebugLevel
	}
	logWriter := zapcore.AddSync(os.Stdout)

	development := viper.GetString("development")
	var encoderCfg zapcore.EncoderConfig
	if development == "true" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "L"
	encoderCfg.CallerKey = "C"
	encoderCfg.TimeKey = "T"
	encoderCfg.NameKey = "N"
	encoderCfg.MessageKey = "M"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder = zapcore.NewJSONEncoder(encoderCfg)

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(level))
	z := zap.New(core, zap.AddCaller())

	logger := z.Sugar()
	if err := logger.Sync(); err != nil {
		// ignore
	}
	return logger
}

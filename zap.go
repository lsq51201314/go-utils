package utils

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志配置
type ZapOptions struct {
	Level      string //debug
	FileName   string //./log/log.txt
	MaxSize    int    //100MB
	MaxAge     int    //90天
	MaxBackups int    //100个
}

// 初始化配置
func ZapInit(options ...ZapOptions) (err error) {
	//默认配置
	cfg := ZapOptions{
		Level:      "debug",
		FileName:   "./log/log.txt",
		MaxSize:    100,
		MaxAge:     90,
		MaxBackups: 100,
	}
	//自定义配置
	if len(options) > 0 {
		if options[0].Level != "" {
			cfg.Level = options[0].Level
		}
		if options[0].FileName != "" {
			cfg.FileName = options[0].FileName
		}
		if options[0].MaxSize > 0 {
			cfg.MaxSize = options[0].MaxSize
		}
		if options[0].MaxAge > 0 {
			cfg.MaxAge = options[0].MaxAge
		}
		if options[0].MaxBackups > 0 {
			cfg.MaxBackups = options[0].MaxBackups
		}
	}
	//初始化
	writeSyncer := getLogWriter(cfg.FileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err = l.UnmarshalText([]byte(cfg.Level)); err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	lg := zap.New(core)
	zap.ReplaceGlobals(lg)
	return
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = getDateTime
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getDateTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.In(location).Format("2006-01-02 15:04:05.999999"))
}

package zap

import (
	"os"
	"time"

	"github.com/lsq51201314/go-utils/location"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//初始化
func Init() error {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte("debug")); err != nil {
		return err
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	lg := zap.New(core)
	zap.ReplaceGlobals(lg)
	return nil
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/log.txt",
		MaxSize:    10,  //文件大小（MB）
		MaxBackups: 100, //文件个数
		MaxAge:     90,  //保留时间（天）
	}
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout)) //输出到控制台
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = getDateTime
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
	//return zapcore.NewJSONEncoder(encoderConfig)
}

func getDateTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.In(location.Location).Format("2006-01-02 15:04:05.000000"))
}

package config

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

var Logger *zap.SugaredLogger

type Log struct {
	Level      string `mapstructure:"level" json:"level"`
	Format     string `mapstructure:"format" json:"format"`
	MaxSize    int    `mapstructure:"max-size" json:"maxSize"`
	MaxAge     int    `mapstructure:"max-age" json:"maxAge"`
	MaxBackups int    `mapstructure:"max-backups" json:"maxBackups"`
	Compress   bool   `mapstructure:"compress" json:"compress"`
}

func init() {
	l := Conf.Log
	logger := zap.New(zapcore.NewTee(l.getZapcores()...))
	Logger = logger.Sugar()
}
func (l *Log) getZapcores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)

	var level zapcore.Level
	err := level.UnmarshalText([]byte(l.Level))
	if err != nil {
		log.Fatal("无效的日志级别:", err)
		return nil
	}
	for ; level < zapcore.FatalLevel; level++ {
		cores = append(cores, l.getZapcore(level))
	}
	return cores
}

func (l *Log) getZapcore(level zapcore.Level) zapcore.Core {
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(l.getEncoderConfig()),
		zapcore.NewMultiWriteSyncer(l.getSyncWriters(level.String()), zapcore.AddSync(os.Stdout)),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl <= level
		}),
	)
	return core
}

func (l *Log) getSyncWriters(level string) zapcore.WriteSyncer {
	now := time.Now()
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("log/%s/%04d-%02d-%02d.log", level, now.Year(), now.Month(), now.Day()),
		MaxSize:    l.MaxSize,    //文件大小限制,单位MB
		MaxAge:     l.MaxAge,     //日志文件保留天数
		MaxBackups: l.MaxBackups, //最大保留日志文件数量
		LocalTime:  false,
		Compress:   l.Compress, //是否压缩处理
	})
	return writeSyncer
}

func (l Log) getEncoderConfig() zapcore.EncoderConfig {
	ecf := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "file",
		FunctionKey:   "func",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return ecf
}

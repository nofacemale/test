package Loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var Logger *zap.Logger

// 初始化或更新 Logger
func InitOrRefreshLogger(logDir string, logName string, errLogName string) {
	if Logger != nil {
		// 关闭之前的 Logger
		err := Logger.Sync()
		if err != nil {
			panic(err.Error())
		}
	}

	// 创建 encoder
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 创建 core
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	infoWriter := getLogWriters(logDir, logName)
	errorWriter := getLogWriters(logDir, errLogName)
	infoCore := zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel)
	errorCore := zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel)

	// 创建 Logger
	Logger = zap.New(zapcore.NewTee(infoCore, errorCore), zap.AddCaller())
}

// 获取日志写入器
func getLogWriters(logDir string, logName string) zapcore.WriteSyncer {
	// 创建日志目录
	os.MkdirAll(filepath.Join("logs", logDir), os.ModePerm)

	// 拼接日志文件路径
	logPath := filepath.Join("logs", logDir, logName)

	// 创建日志文件
	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// 返回写入器
	return zapcore.AddSync(file)
}

// 自定义时间编码器
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-15 15:04:05")) // 将时间格式化为中国日期格式字符串
}

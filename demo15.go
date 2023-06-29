package main

import (
	"demo/loggers"
	"fmt"
	"time"
)

//var logger *zap.Logger

func main() {
	// 初始化 logger
	loggers.InitOrRefreshLogger("gangbao", "app.log", "error.log")
	loggers.Logger.Info("Hello, world!")
	loggers.Logger.Error("test")

	// 动态修改日志文件路径
	newLogDir := "gangbao"
	newLogName := fmt.Sprintf("app_%s.log", time.Now().Format("20060102"))
	newErrLogName := fmt.Sprintf("error_%s.log", time.Now().Format("20060102"))
	loggers.InitOrRefreshLogger(newLogDir, newLogName, newErrLogName)
	loggers.Logger.Info("日志路径已修改")
	loggers.Logger.Error("test")
}

//func main() {
//	// 初始化 logger
//	initOrRefreshLogger("gangbao", "pay", "pay")
//	logger.Info("pay")
//	logger.Error("pay error")
//
//	// 动态修改日志文件路径
//	//newLogName := fmt.Sprintf("app_%s.log", time.Now().Format("20060102"))
//	//newErrLogName := fmt.Sprintf("error_%s.log", time.Now().Format("20060102"))
//	initOrRefreshLogger("gangbao", "refund", "refund")
//	logger.Info("refund info")
//	logger.Error("refund error")
//}

//// 初始化或更新 logger
//func initOrRefreshLogger(logDir string, logName string, errLogName string) {
//	if logger != nil {
//		// 关闭之前的 logger
//		err := logger.Sync()
//		if err != nil {
//			panic(err.Error())
//		}
//	}
//
//	newLogName := fmt.Sprintf("app_%s_%s.log", logName, time.Now().Format("20060102"))
//	newErrLogName := fmt.Sprintf("error_%s_%s.log", errLogName, time.Now().Format("20060102"))
//
//	// 创建 encoder
//	encoderConfig := zapcore.EncoderConfig{
//		TimeKey:        "time",
//		LevelKey:       "level",
//		NameKey:        "logger",
//		CallerKey:      "caller",
//		MessageKey:     "msg",
//		StacktraceKey:  "stacktrace",
//		LineEnding:     zapcore.DefaultLineEnding,
//		EncodeLevel:    zapcore.LowercaseLevelEncoder,
//		EncodeTime:     customTimeEncoder,
//		EncodeDuration: zapcore.SecondsDurationEncoder,
//		EncodeCaller:   zapcore.ShortCallerEncoder,
//	}
//	encoder := zapcore.NewConsoleEncoder(encoderConfig)
//
//	// 创建 core
//	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
//		return lvl < zapcore.ErrorLevel
//	})
//	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
//		return lvl >= zapcore.ErrorLevel
//	})
//	infoWriter := getLogWriters(logDir, newLogName)
//	errorWriter := getLogWriters(logDir, newErrLogName)
//	infoCore := zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel)
//	errorCore := zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel)
//
//	// 创建 logger
//	logger = zap.New(zapcore.NewTee(infoCore, errorCore), zap.AddCaller())
//}
//
//// 获取日志写入器
//func getLogWriters(logDir string, logName string) zapcore.WriteSyncer {
//	// 创建日志目录
//	os.MkdirAll(filepath.Join("logs", logDir), os.ModePerm)
//
//	// 拼接日志文件路径
//	logPath := filepath.Join("logs", logDir, logName)
//
//	// 创建日志文件
//	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//
//	// 返回写入器
//	return zapcore.AddSync(file)
//}
//
//// 自定义时间编码器
//func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//	enc.AppendString(t.Format("2006-01-15 15:04:05")) // 将时间格式化为中国日期格式字符串
//}

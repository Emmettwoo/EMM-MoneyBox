package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var consoleLogger *zap.Logger
var sugaredConsoleLogger *zap.SugaredLogger
var Logger = getLogger()

func initLogger() {
	if consoleLogger == nil {
		consoleLogger = initConsoleLogger()
	}
	if sugaredConsoleLogger == nil {
		sugaredConsoleLogger = consoleLogger.Sugar()
	}

	consoleLogger.Debug("loggers initialize succeed")
}

func getLogger() *zap.SugaredLogger {
	if sugaredConsoleLogger == nil {
		initLogger()
	}
	return sugaredConsoleLogger
}

func initConsoleLogger() *zap.Logger {

	// 設定 console 輸出
	consoleOutput := zapcore.Lock(os.Stdout)
	// 設定 file 輸出
	fileOutput := zapcore.Lock(zapcore.AddSync(createLogFile()))

	// 設定日誌等級
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// 合併輸出和編碼器
	consoleCore := zapcore.NewCore(consoleEncoder, consoleOutput, zapcore.DebugLevel)
	fileCore := zapcore.NewCore(fileEncoder, fileOutput, zapcore.InfoLevel)

	// 同時使用 console 和 file 輸出
	logger := zap.New(zapcore.NewTee(consoleCore, fileCore), zap.AddCaller())

	// 記得在程序結束時關閉 logger
	// defer consoleLogger.Sync()

	return logger
}

func createLogFile() *os.File {
	// 創建 log 文件，您可以自行更改文件路徑和名稱
	file, err := os.Create("./emm-moneybox.log")
	if err != nil {
		log.Fatal("failed to create log file: ", err)
	}
	return file
}

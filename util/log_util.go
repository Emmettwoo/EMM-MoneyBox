package util

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger
var sugaredLogger *zap.SugaredLogger
var Logger *zap.SugaredLogger

func init() {
	logger = initConsoleLogger()
	sugaredLogger = logger.Sugar()
	Logger = sugaredLogger
}

func initFileLogger() *zap.Logger {
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	file, _ := os.Create("./emm-moneybox.log")
	writerSyncer := zapcore.NewMultiWriteSyncer(file)
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	newLogger := zap.New(core)
	return newLogger
}

func initConsoleLogger() *zap.Logger {
	// newLogger, _ := zap.NewProduction()
	rawJSON := []byte(`{
	         "level": "debug",
	         "encoding": "json",
	         "outputPaths": ["stdout"],
	         "errorOutputPaths": ["stderr"],
	         "encoderConfig": {
	           "messageKey": "message",
	           "levelKey": "level",
	           "levelEncoder": "lowercase"
	         }
	       }`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	newLogger, _ := cfg.Build()
	return newLogger
}

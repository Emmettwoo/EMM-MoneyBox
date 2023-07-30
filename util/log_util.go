package util

import (
    "bytes"
    "encoding/json"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "os"
)

var consoleLogger *zap.Logger
var fileLogger *zap.Logger
var sugaredConsoleLogger *zap.SugaredLogger
var Logger = getLogger()
var FileLogger *zap.Logger

func initLogger() {
    if consoleLogger == nil {
        consoleLogger = initConsoleLogger()
    }
    if sugaredConsoleLogger == nil {
        sugaredConsoleLogger = consoleLogger.Sugar()
    }
    if fileLogger == nil {
        fileLogger = initFileLogger()
        FileLogger = fileLogger
    }

    consoleLogger.Debug("loggers initialize succeed")
}

func getLogger() *zap.SugaredLogger {
    if sugaredConsoleLogger == nil {
        initLogger()
    }
    return sugaredConsoleLogger
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
    var loggerLevel = GetConfigByKey("logger.level")
    if loggerLevel == "" {
        loggerLevel = "info"
    }

    var jsonConfiguration bytes.Buffer
    jsonConfiguration.WriteString(`{`)
    jsonConfiguration.WriteString(`  "level": "` + loggerLevel + `",`)
    jsonConfiguration.WriteString(`  "encoding": "json",`)
    jsonConfiguration.WriteString(`  "outputPaths": ["stdout"],`)
    jsonConfiguration.WriteString(`  "errorOutputPaths": ["stderr"],`)
    jsonConfiguration.WriteString(`  "encoderConfig": {`)
    jsonConfiguration.WriteString(`    "messageKey": "message",`)
    jsonConfiguration.WriteString(`    "levelKey": "level",`)
    jsonConfiguration.WriteString(`    "levelEncoder": "lowercase"`)
    jsonConfiguration.WriteString(`  }`)
    jsonConfiguration.WriteString(`}`)

    var cfg zap.Config
    if err := json.Unmarshal(jsonConfiguration.Bytes(), &cfg); err != nil {
        panic(err)
    }
    newLogger, _ := cfg.Build()
    return newLogger
}

package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func InitializeProductionLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	logfile, _ := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logfile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	Sugar = Logger.Sugar()
}

func InitializeDevLogger() {
	Logger, _ = zap.NewDevelopment()
	Sugar = Logger.Sugar()
	defer Logger.Sync()
}

func InitializeCustomLogger() {

	rawJson, err := os.ReadFile("./loggerconfig.json")
	if errors.Is(err, os.ErrNotExist) {
		// b, _ := json.MarshalIndent(defaultLoggerSettings, "", "    ")
		f, _ := os.Create("loggerconfig.json")
		f.WriteString(defaultLoggerSettings)
		rawJson = []byte(defaultLoggerSettings)

	} else if err != nil {
		log.Fatal("Error opening loggerconfig.json", err)
	}

	var cfg zap.Config
	if err = json.Unmarshal(rawJson, &cfg); err != nil {
		log.Fatal("Error unmarshalling data: ", err)
	}

	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	Sugar = Logger.Sugar()

	defer Logger.Sync()
}

const defaultLoggerSettings string = `{
    "level": "info",
    "encoding": "console",
    "development": false,
    "outputPaths": [
        "./gen.log"
    ],
    "errorOutputPaths": [
        "stderr",
        "./gen.log"
    ],
    "encoderConfig": {
        "timeKey": "time",
        "timeEncoder": "iso8601",
        "messageKey": "message",
        "levelKey": "level",
        "levelEncoder": "capital",
        "stacktraceKey": "stacktrace"
    }
}`

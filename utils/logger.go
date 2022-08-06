package utils

import (
	"encoding/json"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

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
}

func InitializeDevLogger() {
	Logger, _ = zap.NewDevelopment()
	defer Logger.Sync()
}

func InitializeCustomLogger() {

	rawJson, err := os.ReadFile("./loggerconfig.json")
	if err != nil {
		log.Fatal("Error opening loggerconfig.json", err)
	}

	var cfg zap.Config
	if err = json.Unmarshal(rawJson, &cfg); err != nil {
		log.Fatal("Error unmarshalling data", err)
	}

	Logger, err = cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer Logger.Sync()
}

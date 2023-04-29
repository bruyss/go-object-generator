package logger

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger

var defaultLoggerSettings = []byte(`{
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
        "stacktraceKey": ""
    }
}`)

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
	Sugar = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
}

func InitializeDevLogger() {
	logger, _ := zap.NewDevelopment()
	Sugar = logger.Sugar()
	defer Sugar.Sync()
}

func InitializeCustomLogger() {
	jsonString, err := ioutil.ReadFile("./loggerconfig.json")
	if errors.Is(err, os.ErrNotExist) {
		err = ioutil.WriteFile("loggerconfig.json", defaultLoggerSettings, 0777)
		if err != nil {
			log.Fatal(err)
		}
		jsonString = defaultLoggerSettings
	} else if err != nil {
		log.Fatal("Error opening logging config file", err)
	}

	var cfg zap.Config
	if err = json.Unmarshal(jsonString, &cfg); err != nil {
		log.Fatal("Error unmarshalling data: ", err)
	}

	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	Sugar = logger.Sugar()

	defer Sugar.Sync()
}

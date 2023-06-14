package logger

import (
	"os"

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
}

func InitializeTestLogger() {
	logger, _ := zap.NewDevelopment()
	Sugar = logger.Sugar()
}

func InitializeCustomLogger() {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	var level zapcore.Level
	if os.Getenv("LOG_LEVEL") == "debug" {
		level = zap.DebugLevel
	} else {
		level = zap.InfoLevel
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"./gen.log"},
		ErrorOutputPaths:  []string{"stderr", "./gen.log"},
	}

	logger := zap.Must(config.Build())

	Sugar = logger.Sugar()
}

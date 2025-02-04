package logger

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LogLevel = zap.NewAtomicLevel()
var Logger *zap.Logger

func InitLogger(filePath string) error {
    if filePath == "" {
        return errors.New("wrong file path")
    }

    LogLevel.SetLevel(zap.ErrorLevel)
    config := CustomConfig(filePath)

    var err error
    Logger, err = config.Build()
    if err != nil {
        return err
    }
    defer Logger.Sync()
    return nil
}

func CloseLogger() {
    Logger.Sync()
}

func CustomConfig(filePath string) zap.Config {
    config := zap.Config {
        Level: zap.NewAtomicLevelAt(zap.InfoLevel),
        Encoding: "json",
        OutputPaths: []string{"stdout", "log/" + filePath},
        ErrorOutputPaths: []string{"stderr"},
        EncoderConfig: zapcore.EncoderConfig{
            TimeKey:        "time",
            LevelKey:       "level",
            NameKey:        "logger",
            CallerKey:      "caller",
            MessageKey:     "msg",
            LineEnding:     zapcore.DefaultLineEnding,
            EncodeLevel:    zapcore.LowercaseLevelEncoder,
            EncodeTime:     zapcore.ISO8601TimeEncoder,
            EncodeCaller:   zapcore.ShortCallerEncoder,
        },
    }
    return config
}

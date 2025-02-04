package main

import (
	"errors"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LogLevel = zap.NewAtomicLevel()
var Logger *zap.Logger

func main() {
    LogLevel.SetLevel(zap.ErrorLevel)
    config := CustomConfig(LogLevel)

    Logger, _ = config.Build()
    defer Logger.Sync()

    Logger.Error(
        "Failed to say hi",
        zap.String("url", "http://localhost:8080"),
    )

    LogLevel.SetLevel(zap.PanicLevel)
    Logger.Panic(
        "Serious Bug in the Database",
        zap.String("url", "http://localhost:8080/database"),
    )
}

func Level(rawLevel string) error {
    level := strings.ToLower(strings.TrimSpace(rawLevel))
    switch level {
    case "error":
        LogLevel.SetLevel(zap.ErrorLevel)
    case "warm":
        LogLevel.SetLevel(zap.WarnLevel)
    case "panic":
        LogLevel.SetLevel(zap.PanicLevel)
    case "info":
        LogLevel.SetLevel(zap.InfoLevel)
    case "debug":
        LogLevel.SetLevel(zap.DebugLevel)
    case "fatal":
        LogLevel.SetLevel(zap.FatalLevel)
    default:
        return errors.New("no level as given")
    }
    return nil
}

func CustomConfig(LogLevel zap.AtomicLevel) zap.Config {
    config := zap.Config {
        Level: LogLevel,
        Encoding: "json",
        OutputPaths: []string{"stdout", "app.log"},
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

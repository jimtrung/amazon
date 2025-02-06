package logger

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout", "log/" + filePath},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "time",
			LevelKey:     "level",
			NameKey:      "logger",
			CallerKey:    "caller",
			MessageKey:   "msg",
			LineEnding:   zapcore.DefaultLineEnding,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	return config
}

func LogAndRespond(
	c *gin.Context, logFile string, message string, erro error, statusCode int,
	data ...interface{},
) {
	if err := InitLogger(logFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer CloseLogger()

	if erro != nil {
		Logger.Error(
			message,
			zap.String("error", erro.Error()),
			zap.String("url", c.Request.URL.String()),
		)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	Logger.Info(
		message,
		zap.Any("data", data),
		zap.String("url", c.Request.URL.String()),
	)

	if data != nil {
		c.JSON(statusCode, data)
	} else {
		c.JSON(statusCode, message)
	}
}

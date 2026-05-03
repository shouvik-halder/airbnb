package logger

import (
	"AuthenticationService/config"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zerolog.Logger

func InitLogger(cfg *config.Config) *zerolog.Logger {
	var writers []io.Writer

	// writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr })
	writers = append(writers, initRollingFileLogger(cfg))
	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().
		Timestamp().
		Logger()
	Log = &logger
	return Log
}

func initRollingFileLogger(cfg *config.Config) *lumberjack.Logger {
	date := time.Now().Format("02-01-2006")
	fileName := fmt.Sprintf("%s-app.log", date)

	loggerPath := filepath.Join("logs", fileName)

	fmt.Printf("logging to file: %s\n", loggerPath)

	// ensure logs directory exists
	_ = os.MkdirAll("logs", os.ModePerm)

	return &lumberjack.Logger{
		Filename:   loggerPath,
		MaxBackups: cfg.Logger.MAXBACKUPCOUNT,
		MaxSize:    cfg.Logger.MAXSIZEMB,
		MaxAge:     cfg.Logger.MAXAGEDAYS,
		Compress:   true,
	}
}
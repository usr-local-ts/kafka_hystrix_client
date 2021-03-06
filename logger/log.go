package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/usr-local-ts/kafka_hystrix_client/config"
)

type Logger struct {
	*logrus.Logger
}

var Log *Logger

type LoggerError struct {
	Error error
}

func panicIfError(err error) {
	if err != nil {
		panic(LoggerError{err})
	}
}

func SetupLogger() {
	level, err := logrus.ParseLevel(config.LogLevel())
	panicIfError(err)

	logrusVar := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	Log = &Logger{logrusVar}
}

func BuildContext(context string) logrus.Fields {
	return logrus.Fields{
		"context": context,
	}
}
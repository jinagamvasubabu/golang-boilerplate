package loggers

import (
	"os"
	"strings"

	"github.com/jinagamvasubabu/golang-boilerplate/config"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func InitLogger() {
	cfg := config.GetConfig()
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Logger = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}

	Logger.Formatter = &logrus.JSONFormatter{}

}

func Debug(msg string, tags ...string) {
	if Logger.Level < logrus.DebugLevel {
		return
	}
	Logger.WithFields(parseFields(tags...)).Debug(msg)
}

func Debugf(format string, args ...interface{}) {
	if Logger.Level < logrus.DebugLevel {
		return
	}
	Logger.Debugf(format, args...)
}

func Info(msg string, tags ...string) {
	if Logger.Level < logrus.InfoLevel {
		return
	}
	Logger.WithFields(parseFields(tags...)).Info(msg)
}

func Infof(format string, args ...interface{}) {
	if Logger.Level < logrus.InfoLevel {
		return
	}
	Logger.Infof(format, args...)
}

func Error(msg string, tags ...string) {
	if Logger.Level < logrus.InfoLevel {
		return
	}
	Logger.WithFields(parseFields(tags...)).Error(msg)
}

func Errorf(format string, args ...interface{}) {
	if Logger.Level < logrus.InfoLevel {
		return
	}
	Logger.Errorf(format, args...)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	return result
}

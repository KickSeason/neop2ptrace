package log

import (
	log "github.com/sirupsen/logrus"
)

func NewLogger() *log.Logger {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
	})
	return logger
}

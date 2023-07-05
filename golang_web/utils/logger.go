package utils

import (
	"blog_web/utils/logger"
	"github.com/spf13/viper"
)

var logInstance *logger.Logger

func Logger() *logger.Logger {
	return logInstance
}

func CreateLogger() {
	logInstance = logger.NewFileLogger(viper.GetString("log.filepath"),
		viper.GetString("log.filename"),
		viper.GetString("log.level"),
	)
}

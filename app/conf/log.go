package conf

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

type LOG struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
	LogLines bool   `env:"LOG_FILE_INLINE" envDefault:"true"`
	Json     bool   `env:"LOG_JSON" envDefault:"false"`
}

func (c *config) LOG() *logrus.Entry {
	if c.log != nil {
		return c.log
	}

	logConf := &LOG{}

	if err := env.Parse(logConf); err != nil {
		panic(err)
	}

	logger := logrus.New()
	if logConf.Json {
		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	level, _ := logrus.ParseLevel(logConf.LogLevel)

	logger.SetLevel(level)
	logger.SetReportCaller(logConf.LogLines)

	c.log = logrus.NewEntry(logger)

	return c.log
}

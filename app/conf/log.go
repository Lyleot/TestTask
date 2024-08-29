package conf

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

// LOG содержит настройки для логирования.
type LOG struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`      // Уровень логирования
	LogLines bool   `env:"LOG_FILE_INLINE" envDefault:"true"` // Включить отображение номеров строк
	Json     bool   `env:"LOG_JSON" envDefault:"false"`       // Использовать формат JSON для логов
}

// LOG возвращает экземпляр логгера с настройками, инициализируя его при необходимости.
func (c *config) LOG() *logrus.Entry {
	if c.log != nil {
		return c.log // Возвращает уже инициализированный логгер
	}

	logConf := &LOG{}

	if err := env.Parse(logConf); err != nil {
		panic(err) // Паника при ошибке загрузки переменных окружения
	}

	logger := logrus.New()
	if logConf.Json {
		logger.SetFormatter(&logrus.JSONFormatter{}) // Форматирование логов в JSON
	}

	level, _ := logrus.ParseLevel(logConf.LogLevel) // Установка уровня логирования
	logger.SetLevel(level)
	logger.SetReportCaller(logConf.LogLines) // Включение отображения номеров строк в логах

	c.log = logrus.NewEntry(logger)

	return c.log
}

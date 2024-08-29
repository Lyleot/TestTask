package conf

import "github.com/caarlos0/env"

// HTTP содержит настройки для HTTP сервера.
type HTTP struct {
	HTTP_PORT string `env:"HTTP_PORT" envDefault:"80"` // Порт HTTP сервера
}

// HTTP возвращает экземпляр конфигурации HTTP, инициализируя его при необходимости.
func (c *config) HTTP() *HTTP {
	if c.http != nil {
		return c.http // Возвращает уже инициализированную конфигурацию
	}

	c.http = &HTTP{}

	if err := env.Parse(c.http); err != nil {
		panic(err) // Паника при ошибке загрузки переменных окружения
	}

	return c.http
}

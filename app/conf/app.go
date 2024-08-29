package conf

import "github.com/caarlos0/env"

// APP содержит конфигурацию для приложения.
type APP struct {
	ApiBaseURL   string `env:"API_BASE_URL" envDefault:"localhost"`   // Базовый URL API
	ApiWorkerURL string `env:"API_WORKER_URL" envDefault:"localhost"` // URL API рабочего процесса
}

// APP возвращает конфигурацию приложения, загруженную из переменных окружения.
func (c *config) APP() *APP {
	// Если конфигурация уже загружена, возвращаем её.
	if c.app != nil {
		return c.app
	}

	// Инициализируем конфигурацию.
	c.app = &APP{}

	// Загружаем значения из переменных окружения в конфигурацию.
	if err := env.Parse(c.app); err != nil {
		panic(err) // Если произошла ошибка, приложение завершает работу.
	}

	return c.app
}

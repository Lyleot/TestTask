package conf

import (
	"TestTask/app/store"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Config определяет интерфейс конфигурации приложения.
type Config interface {
	DB() store.Store    // Метод для получения подключения к базе данных
	LOG() *logrus.Entry // Метод для получения объекта логгера
	APP() *APP          // Метод для получения конфигурации приложения
	HTTP() *HTTP        // Метод для получения конфигурации HTTP
}

// config содержит внутренние поля конфигурации и реализует интерфейс Config.
type config struct {
	sync.Mutex

	db   store.Store
	log  *logrus.Entry
	app  *APP
	http *HTTP
}

// LoadDotEnv загружает переменные окружения из файла .env.
func LoadDotEnv(stepsUp int) error {
	path := strings.Repeat("../", stepsUp) + ".env" // Формирует путь к .env файлу
	if err := godotenv.Load(path); err != nil {
		return err // Возвращает ошибку, если не удалось загрузить .env файл
	}
	return nil
}

// New создает новый экземпляр конфигурации и загружает переменные окружения.
func New() Config {
	LoadDotEnv(0) // Загружает переменные окружения

	c := &config{}

	// Инициализация минимально необходимых полей
	c.DB()

	return c
}

package conf

import (
	"TestTask/app/store"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config interface {
	DB() store.Store
	LOG() *logrus.Entry
	APP() *APP
	HTTP() *HTTP
}

type config struct {
	sync.Mutex

	db   store.Store
	log  *logrus.Entry
	app  *APP
	http *HTTP
}

func LoadDotEnv(stepsUp int) error {
	path := strings.Repeat("../", stepsUp) + ".env"
	if err := godotenv.Load(path); err != nil {
		return err
	}
	return nil
}

func New() Config {

	LoadDotEnv(0)

	c := &config{}

	// minimal init
	c.DB()

	return c
}

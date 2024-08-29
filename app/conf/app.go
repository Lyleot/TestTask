package conf

import "github.com/caarlos0/env"

type APP struct {
	ApiBaseURL   string `env:"API_BASE_URL" envDefault:"localhost"`
	ApiWorkerURL string `env:"API_WORKER_URL" envDefault:"localhost"`
}

func (c *config) APP() *APP {
	if c.app != nil {
		return c.app
	}

	c.app = &APP{}

	if err := env.Parse(c.app); err != nil {
		panic(err)
	}

	return c.app
}

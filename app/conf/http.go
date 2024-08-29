package conf

import "github.com/caarlos0/env"

type HTTP struct {
	HTTP_PORT string `env:"HTTP_PORT" envDefault:"80"`
}

func (c *config) HTTP() *HTTP {
	if c.http != nil {
		return c.http
	}

	c.http = &HTTP{}

	if err := env.Parse(c.http); err != nil {
		panic(err)
	}

	return c.http
}

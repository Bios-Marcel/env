package config

import "module_b/conf"

type Nested struct {
	Field string `env:"FIELD"`
}

type Config struct {
	Nested Nested `envPrefix:"NESTED_"`
}

func New() (*Config, error) {
	var c Config
	if err := conf.Parse(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

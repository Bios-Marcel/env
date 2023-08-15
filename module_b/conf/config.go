package conf

import "github.com/caarlos0/env/v9"

func Parse[T any](target *T) error {
	return env.ParseWithOptions(target, env.Options{})
}

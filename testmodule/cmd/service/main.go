package main

import (
	"fmt"

	"github.com/caarlos0/env/v9"

	"testmodule/internal/config"
)

func main() {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	fmt.Println(cfg.Nested.Field)
}

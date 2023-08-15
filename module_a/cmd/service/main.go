package main

import (
	"fmt"

	"module_a/internal/config"
)

func main() {
	if cfg, err := config.New(); err != nil {
		panic(err)
	} else {
		fmt.Println(cfg.Nested.Field)
	}
}

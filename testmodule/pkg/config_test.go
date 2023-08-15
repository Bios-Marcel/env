package pkg

import (
	"testing"

	"github.com/caarlos0/env/v9"

	"testmodule/internal/config"
)

func TestParseInternalConfig(t *testing.T) {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

package pkg

import (
	"testing"

	"module_a/internal/config"
)

func TestParseInternalConfig(t *testing.T) {
	t.Setenv("NESTED_FIELD", "value")

	if cfg, err := config.New(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	} else {
		if cfg.Nested.Field != "value" {
			t.Fatalf("unexpected value: %v", cfg.Nested.Field)
		}
	}
}

package config

type Nested struct {
	Field string `env:"FIELD"`
}

type Config struct {
	Nested Nested `envPrefix:"NESTED_"`
}

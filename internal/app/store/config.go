package store

type Config struct {
	DatabaseURL string `toml:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{}
}

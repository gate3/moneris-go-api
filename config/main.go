package config

type Config struct {
	IsLiveEnvironment bool
}

func NewConfig() *Config {
	return &Config{}
}

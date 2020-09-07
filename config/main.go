package config

type Config struct {
	IsLiveEnvironment bool
	ApiToken          string
	StoreId           string
}

func NewConfig(isLiveEnvironment bool, apiToken, storeId string) *Config {
	return &Config{
		IsLiveEnvironment: isLiveEnvironment,
		ApiToken:          apiToken,
		StoreId:           storeId,
	}
}

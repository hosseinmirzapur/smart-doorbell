package config

import "os"

type kavenegarConfig struct {
	ApiKey string
}

func GetKavenegarConfig() *kavenegarConfig {
	return &kavenegarConfig{
		ApiKey: os.Getenv("KAVENEGAR_API_KEY"),
	}
}

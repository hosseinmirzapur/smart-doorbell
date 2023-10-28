package config

import "os"

type botConfig struct {
	Token string
}

func GetBotConfig() *botConfig {
	return &botConfig{
		Token: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}
}

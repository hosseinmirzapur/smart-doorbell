package config

import (
	"os"

	"github.com/hosseinmirzapur/sdb/utils"
)

type emailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func GetEmailConfig() *emailConfig {

	return &emailConfig{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     utils.StringToInt(os.Getenv("MAIL_PORT")),
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
}

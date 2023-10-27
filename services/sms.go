package services

import (
	"github.com/hosseinmirzapur/sdb/config"
	"github.com/kavenegar/kavenegar-go"
)

type smsService struct{}

var api *kavenegar.Kavenegar

func NewSmsService() *smsService {
	kavenegarAPIKey := config.GetKavenegarConfig().ApiKey
	api = kavenegar.New(kavenegarAPIKey)
	return &smsService{}
}

func (sms *smsService) Send(receptor []string, sender, message string) error {
	_, err := api.Message.Send(sender, receptor, message, nil)

	return err
}

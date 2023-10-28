package services

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hosseinmirzapur/sdb/config"
)

type botService struct{}

var telegramBot *tgbotapi.BotAPI

func NewBotService() *botService {

	conf := config.GetBotConfig()

	bot, err := tgbotapi.NewBotAPI(conf.Token)
	if err != nil {
		log.Fatalln(err)
	}

	bot.Debug = true

	telegramBot = bot

	return &botService{}
}

func (b *botService) SendMessage(msg string) error {
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := telegramBot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		messageToSend := tgbotapi.NewMessage(update.Message.Chat.ID, msg)

		_, err := telegramBot.Send(messageToSend)
		if err != nil {
			return err
		}
	}
	return nil
}

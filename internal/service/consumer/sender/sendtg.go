package sender

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *service) SendToTelegram(text string) error {
	bot, err := tgbotapi.NewBotAPI(s.senderConf.Token())
	if err != nil {
		return fmt.Errorf("не удалось создать бота: %v", err)
	}

	msg := tgbotapi.NewMessage(s.senderConf.ID(), text)
	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("не удалось отправить сообщение в Telegram: %v", err)
	}

	return nil
}

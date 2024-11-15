package sender

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Dnlbb/notifier/internal/model"
	"github.com/IBM/sarama"
)

func (s *service) NoteSaveHandler(_ context.Context, msg *sarama.ConsumerMessage) error {
	messageText := string(msg.Value)
	var user model.User
	err := json.Unmarshal([]byte(messageText), &user)
	if err != nil {
		return fmt.Errorf("ошибка декодирования: %w", err)
	}
	messageSend := fmt.Sprintf("Зарегался пользователь %s c ролью: %s", user.Name, user.Role)

	err = s.SendToTelegram(messageSend)
	if err != nil {
		log.Printf("Ошибка отправки в Telegram: %v", err)
		return err
	}

	log.Println("Сообщение успешно отправлено в Telegram и на почту.")
	return nil
}

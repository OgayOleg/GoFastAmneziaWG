package service

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (s *service) Update() error {
	for {
		updates, err := s.tg.Bot.GetUpdatesChan(tgbotapi.NewUpdate(0))
		if err != nil {
			return fmt.Errorf("not created updates: %w", err)
		}
		for update := range updates {
			if update.Message == nil {
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!")
			s.tg.Bot.Send(msg)

		}
	}
}

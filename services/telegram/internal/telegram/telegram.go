package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Tg struct {
	Bot *tgbotapi.BotAPI
}

func New(TELEGRAM_KEY string) (*Tg, error) {
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_KEY)
	if err != nil {
		return nil, err
	}
	return &Tg{
		Bot: bot,
	}, nil
}

func (tg *Tg) Configure() ([]tgbotapi.Update, error) {

	// включаем дебагер для лучшего вывода в консоль
	tg.Bot.Debug = true

	log.Println("Authorized on account", tg.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := tg.Bot.GetUpdates(u)
	if err != nil {
		return nil, err
	}

	return updates, nil

}

package service

import "telegram-service/internal/telegram"

type service struct {
	tg *telegram.Tg
}

func New(tg *telegram.Tg) service {

	return service{
		tg: tg,
	}
}

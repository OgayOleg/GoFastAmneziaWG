package main

import (
	"os"
	"telegram-service/internal/service"
	"telegram-service/internal/telegram"
)

func main() {
	tgKey := os.Getenv("TELEGRAM_KEY")
	if len(tgKey) == 0 {
		panic("TELEGRAM_KEY environment variable is not set")
	}

	tg, err := telegram.New(tgKey)
	if err != nil {
		panic(err)
	}

	service := service.New(tg)

	service.Update()
}

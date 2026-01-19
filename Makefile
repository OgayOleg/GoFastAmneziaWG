AWG_ENV := services/awg/.env
TELEGRAM_ENV := services/telegram/.env

.PHONY: awg-run telegram-run

awg-run:
	@export $$(cat $(AWG_ENV) | xargs) && \
	cd services/awg && sudo -E go run cmd/main.go

telegram-run:
	@export $$(cat $(TELEGRAM_ENV) | xargs) && \
	cd services/telegram && go run cmd/main.go

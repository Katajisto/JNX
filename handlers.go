package main

import (
	tba "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	handle("ping", func(x []string, bot *tba.BotAPI, id int64) {
		send("pong", bot, id)
	})
	handle("subscribe", func(x []string, bot *tba.BotAPI, id int64) {
		go routine(bot, id)
	})
}

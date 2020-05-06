package main

import (
	"log"

	tba "github.com/go-telegram-bot-api/telegram-bot-api"
)

func send(msg string, bot *tba.BotAPI, id int64) {
	toSend := tba.NewMessage(id, msg)
	bot.Send(toSend)
	log.Println("BOT: ", msg)
}

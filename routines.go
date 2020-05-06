package main

import (
	"log"

	tba "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	addRoutine("2223", func(bot *tba.BotAPI, id int64) {
		log.Println("SENT")
		send("ITS TIME TO STOP", bot, id)
	})
}

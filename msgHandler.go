package main

import (
	"log"

	tba "github.com/go-telegram-bot-api/telegram-bot-api"
)

type handler func([]string, *tba.BotAPI, int64)

var currentAuthChat int64
var handlerMap = make(map[string]handler)

func handle(toHandle string, h handler) {
	handlerMap[toHandle] = h

}

func callHandler(bot *tba.BotAPI, msg *tba.Message) {

	//Split message into payload and such.
	sepMsg := stringToSlice(msg.Text, ' ')
	log.Println("MSG: ", sepMsg)
	log.Println(msg.Chat.ID)
	if currentAuthChat != msg.Chat.ID {
		if sepMsg[0] == "auth" {
			if sepMsg[1] == AUTH_PASS {
				currentAuthChat = msg.Chat.ID
				send("Auth successful!", bot, msg.Chat.ID)
			} else {
				send("Auth unsuccessful!", bot, msg.Chat.ID)
			}
		} else {
			send("Please auth first!", bot, msg.Chat.ID)
		}
		return
	}
	_, ok := handlerMap[sepMsg[0]]
	if !ok {
		log.Println("No handler for: ", sepMsg[0])
		return
	}
	handlerMap[sepMsg[0]](sepMsg[1:], bot, msg.Chat.ID)
}

package main

import (
	"log"
	"time"

	tba "github.com/go-telegram-bot-api/telegram-bot-api"
)

type routineHandler func(bot *tba.BotAPI, id int64)

var routineMap = make(map[string][]routineHandler)

func addRoutine(time string, h routineHandler) {
	routineMap[time] = append(routineMap[time], h)
}

func routine(bot *tba.BotAPI, id int64) {
	log.Println("CALL")
	lastDone := ""
	for {
		time.Sleep(20 * time.Second)
		curTime := time.Now()
		timeString := curTime.Format("1504")
		if timeString == lastDone {
			continue
		}
		_, ok := routineMap[timeString]
		if ok {
			for _, h := range routineMap[timeString] {
				h(bot, id)
			}
		}
		lastDone = timeString
	}
}

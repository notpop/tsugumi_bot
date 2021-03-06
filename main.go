package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"tsugumi_bot/weather"
)

func main() {
	lineBot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	result, err := weather.GetWeather()
	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(result)
	if _, err := lineBot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}

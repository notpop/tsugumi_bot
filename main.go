package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	// "github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/robfig/cron/v3"
	"tsugumi_bot/weather"
)

func main() {
	cron := cron.New()
	// cron.AddFunc("CRON_TZ=Asia/Tokyo 0 7 * * *", func() {
	cron.AddFunc("@every 30s", func() {
		fmt.Println("start linebot")
		broadcastWeather()
		fmt.Println("end linebot")
	})
	cron.Start()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf(err.Error())
	}
}

func broadcastWeather() {
	// godotenv.Load(".env")
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

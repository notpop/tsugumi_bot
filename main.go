package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"strings"
	"tsugumi_bot/config"
	"tsugumi_bot/line"
	"tsugumi_bot/openai"
	"tsugumi_bot/utils"
	"tsugumi_bot/weather"
)

func init() {
	utils.LoggingSettings(config.Config.SystemLog)
}

func broadcastWeather() {
	result, err := weather.GetWeather()
	if err != nil {
		log.Fatal(err)
	}

	line, err := line.New(config.Config.ChannelSecret, config.Config.ChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	err = line.BroadcastMessage(result)
	if err != nil {
		log.Fatal(err)
	}
}

func mainHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("start main handler")
	line, err := line.New(config.Config.ChannelSecret, config.Config.ChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	events, err := line.Client.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				response, err := openai.SendQuestion(message.Text)
				if err != nil {
					log.Print(err)
				}

				replymessage := strings.ReplaceAll(response.Choices[0].Text, "\n", "")
				if _, err = line.Client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replymessage)).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.StickerMessage:
				replyMessage := fmt.Sprintf("スタンプIDが%sで種類が%sだよ！", message.StickerID, message.StickerResourceType)
				if _, err = line.Client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
	log.Println("end main handler")
}

func main() {
	instance := cron.New()
	instance.AddFunc("CRON_TZ=Asia/Tokyo 0 8 * * *", func() {
		log.Println("start term execute linebot.")
		broadcastWeather()
		log.Println("end term execute linebot.")
	})
	instance.Start()

	http.HandleFunc("/webhook", mainHandler)

	if err := http.ListenAndServe(config.Config.Port, nil); err != nil {
		log.Fatal(err)
	}
}

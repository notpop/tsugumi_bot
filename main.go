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

const (
	DEFAULT_OUTPUT_LOG_BUFFER_TIME = 10
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

func getAnswer(message string) string {
	log.Println("question: " + message)
	response, err := openai.SendQuestion(message)
	if err != nil {
		log.Print(err)
	}

	answer := response.Choices[0].Text
	log.Println("answer: " + answer)
	return answer
}

func replaceIndention(message string) string {
	return strings.ReplaceAll(message, "\n", "")
}

func getStampMessage(id string, resouceType linebot.StickerResourceType) string {
	return fmt.Sprintf("スタンプIDが%sで種類が%sだよ！", id, resouceType)
}

func webhooker(w http.ResponseWriter, req *http.Request) {
	log.Println("start webhooker")

	if line.IsLimit() {
		log.Println("message api limit")

		return
	}

	line, err := line.New(config.Config.ChannelSecret, config.Config.ChannelToken)
	if err != nil {
		log.Fatal(err)
	}

	events, err := line.Client.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)

			return
		}

		w.WriteHeader(500)

		return
	}

	for _, event := range events {
		if event.Type != linebot.EventTypeMessage {
			continue
		}

		switch message := event.Message.(type) {
		case *linebot.TextMessage:
			question := message.Text
			answer := getAnswer(question)
			replacedAnswer := replaceIndention(answer)
			err = line.ReplyMessageWithLog(replacedAnswer, event.ReplyToken, DEFAULT_OUTPUT_LOG_BUFFER_TIME)
			if err != nil {
				log.Fatal(err)
			}
		case *linebot.StickerMessage:
			replyMessage := getStampMessage(message.StickerID, message.StickerResourceType)
			err = line.ReplyMessage(replyMessage, event.ReplyToken)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("end webhooker")
}

func main() {
	instance := cron.New()
	instance.AddFunc("CRON_TZ=Asia/Tokyo 0 8 * * *", func() {
		log.Println("start term execute linebot.")
		broadcastWeather()
		log.Println("end term execute linebot.")
	})
	instance.Start()

	http.HandleFunc("/webhook", webhooker)

	if err := http.ListenAndServe(config.Config.Port, nil); err != nil {
		log.Fatal(err)
	}
}

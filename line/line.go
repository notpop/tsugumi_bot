package line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
)

type Line struct {
	Client *linebot.Client
}

type LineResponse struct {
	RequestBody string `json:"RequestBody"`
}

type LineRequest struct {
	Destination string `json:"destination"`
	Events      []struct {
		Type    string `json:"type"`
		Message struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"message,omitempty"`
		Timestamp int64 `json:"timestamp"`
		Source    struct {
			Type   string `json:"type"`
			UserID string `json:"userId"`
		} `json:"source"`
		ReplyToken      string `json:"replyToken,omitempty"`
		Mode            string `json:"mode"`
		WebhookEventID  string `json:"webhookEventId"`
		DeliveryContext struct {
			IsRedelivery bool `json:"isRedelivery"`
		} `json:"deliveryContext"`
	} `json:"events"`
}

func New(secret string, token string) (*Line, error) {
	lineBot, err := linebot.New(secret, token)
	if err != nil {
		log.Println("linebot new error.")

		return nil, err
	}

	client := &Line{lineBot}
	return client, nil
}

func (line *Line) BroadcastMessage(_message string) error {
	message := linebot.NewTextMessage(_message)
	if _, err := line.Client.BroadcastMessage(message).Do(); err != nil {
		log.Println("linebot broadcast error.")

		return err
	}

	return nil
}

package line

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"time"
)

type Client struct {
	Client *linebot.Client
}

type Webhook struct {
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

func New(secret string, token string) (*Client, error) {
	lineBot, err := linebot.New(secret, token)
	if err != nil {
		log.Println("linebot new error.")

		return nil, err
	}

	client := &Client{lineBot}
	return client, nil
}

func (client *Client) BroadcastMessage(message string) error {
	if _, err := client.Client.BroadcastMessage(linebot.NewTextMessage(message)).Do(); err != nil {
		log.Println("linebot broadcast error.")

		return err
	}

	return nil
}

func (client *Client) ReplyMessageWithLog(message string, replyToken string, coefficient int64) error {
	if _, err := client.Client.ReplyMessage(replyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Println("linebot replaymessagewithlog error.")

		return err
	}

	time.Sleep(time.Second * time.Duration(coefficient))
	log.Println("replymessage: " + message)

	return nil
}

func (client *Client) ReplyMessage(message string, replyToken string) error {
	if _, err := client.Client.ReplyMessage(replyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Println("linebot replaymessage error.")

		return err
	}

	return nil
}

func (client *Client) pushMessage(message string, id string) error {
	if _, err := client.Client.PushMessage(id, linebot.NewTextMessage(message)).Do(); err != nil {
		log.Println("linebot pushmessage error.")

		return err
	}

	return nil
}

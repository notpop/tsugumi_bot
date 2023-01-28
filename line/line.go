package line

import (
	"encoding/json"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"time"
	"tsugumi_bot/config"
	"tsugumi_bot/utils"
)

const (
	ERROR_INT64           = 1000000000
	ERROR_SETTING_INT64   = 999999999
	MESSAGE_API_V2_URL    = "https://api.line.me/v2/bot/message"
	CONSUMPTION_END_POINT = "quota/consumption"
	QUOTA_END_POINT       = "quota"
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

type MessageLimitResponse struct {
	TotalUsage int64 `json:"totalUsage"`
}

func GetMessageApiLimit() (int64, error) {
	client := utils.New(MESSAGE_API_V2_URL, map[string]string{
		"Authorization": "Bearer " + config.Config.ChannelToken,
	})
	resp, err := client.DoRequest("GET", CONSUMPTION_END_POINT, map[string]string{}, []byte{})
	if err != nil {
		log.Println("Get Http Error: ", err)
		return ERROR_INT64, err
	}

	var response MessageLimitResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Println("JSON Unmarshal error: ", err)
		return ERROR_INT64, err
	}

	return response.TotalUsage, nil
}

type MessageLimitSettingResponse struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

func GetMessageApiLimitSetting() (int64, error) {
	client := utils.New(MESSAGE_API_V2_URL, map[string]string{
		"Authorization": "Bearer " + config.Config.ChannelToken,
	})
	resp, err := client.DoRequest("GET", QUOTA_END_POINT, map[string]string{}, []byte{})
	if err != nil {
		log.Println("Get Http Error: ", err)
		return ERROR_SETTING_INT64, err
	}

	var response MessageLimitSettingResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Println("JSON Unmarshal error: ", err)
		return ERROR_SETTING_INT64, err
	}

	return response.Value, nil
}

func IsLimit() bool {
	limit, err := GetMessageApiLimit()
	if err != nil {
		return true
	}

	limitSetting, err := GetMessageApiLimitSetting()
	if err != nil {
		return true
	}

	return limit > limitSetting
}

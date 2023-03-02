package openai

import (
	"encoding/json"
	"log"
	"tsugumi_bot/config"
	"tsugumi_bot/utils"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGpt struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	MaxToken int64     `json:"max_tokens"`
}

type ChatGptResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

const (
	URL                         = "https://api.openai.com/v1/"
	CREATE_COMPLETION_END_POINT = "chat/completions"
	DEFAULT_MODEL               = "gpt-3.5-turbo"
	DEFAULT_MAX_TOKEN           = 4000
)

func header() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + config.Config.ApiKey,
	}
}

func SendQuestion(prompt string) (*ChatGptResponse, error) {
	chatGpt := &ChatGpt{
		Model:    DEFAULT_MODEL,
		MaxToken: DEFAULT_MAX_TOKEN,
		Messages: []Message{
			{Role: "system", Content: config.Config.TsugumiSettings},
			{Role: "user", Content: prompt},
		},
	}

	data, err := json.Marshal(chatGpt)
	if err != nil {
		log.Println("JSON Marshal error: ", err)
		return nil, err
	}

	client := utils.New(URL, header())
	resp, err := client.DoRequest("POST", CREATE_COMPLETION_END_POINT, map[string]string{}, data)
	if err != nil {
		log.Println("Post Http Error: ", err)
		return nil, err
	}

	var response ChatGptResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Println("JSON Unmarshal error: ", err)
		return nil, err
	}

	return &response, nil
}

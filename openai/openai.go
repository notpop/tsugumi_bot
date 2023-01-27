package openai

import (
	"encoding/json"
	"log"
	"tsugumi_bot/config"
	"tsugumi_bot/utils"
)

type ChatGpt struct {
	Model    string `json:"model"`
	Prompt   string `json:"prompt"`
	MaxToken int64  `json:"max_tokens"`
}

type ChatGptResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

const (
	URL                         = "https://api.openai.com/v1/"
	CREATE_COMPLETION_END_POINT = "completions"
)

func header() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + config.Config.ApiKey,
	}
}

func SendQuestion(prompt string) (*ChatGptResponse, error) {
	chatGpt := &ChatGpt{
		Model:    "text-davinci-003",
		Prompt:   prompt,
		MaxToken: 4000,
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

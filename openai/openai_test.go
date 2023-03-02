package openai

import (
	"fmt"
	"testing"
)

func ExampleSendQuestion() {
	result, err := SendQuestion("Say this is a test")
	if err != nil {
		fmt.Println("結果が取れてません!!!")
	}

	text := result.Choices[0].Message.Content
	fmt.Println(text)
}

func TestSendQuestion(t *testing.T) {
	result, err := SendQuestion("本日の天気はなんですか？")
	t.Log(result)
	if err != nil {
		t.Error("\n結果が取れてません!!!")
	}

	if len(result.Choices) == 0 {
		t.Error("\nAnswerが取れてません!!!")
	}

	text := result.Choices[0].Message.Content
	if text == "" {
		t.Error("\nAnswerが取れてません!!!")
	}

	t.Log(text)
	t.Log("SendQuestion終了")
}

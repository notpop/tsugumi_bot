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

	text := result.Choices[0].Text
	fmt.Println(text)
}

func TestHSendQuestion(t *testing.T) {
	result, err := SendQuestion("Say this is a test")
	if err != nil {
		t.Error("\n結果が取れてません!!!")
	}

	text := result.Choices[0].Text
	if text == "" {
		t.Error("\nAnswerが取れてません!!!")
	}

	t.Log(text)
	t.Log("SendQuestion終了")
}

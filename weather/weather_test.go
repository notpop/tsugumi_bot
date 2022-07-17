package weather

import (
	"fmt"
)

func ExampleToS() {
	w := new(Weather)
	w.Area = "テスト地方"
	w.HeadLine = "一日良い天気です。"
	w.Body = "傘を持ち歩く必要はないでしょう。"

	fmt.Println(w.ToS())
	// Output: テスト地方の天気です。
	// 一日良い天気です。
	// 傘を持ち歩く必要はないでしょう。
}

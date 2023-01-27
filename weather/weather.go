package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"tsugumi_bot/config"
	"tsugumi_bot/utils"
)

type Weather struct {
	Area     string `json:"targetArea"`
	HeadLine string `json:"headlineText"`
	Body     string `json:"text"`
}

const (
	URL             = "https://www.jma.go.jp/bosai/forecast/data/overview_forecast/"
	PREFECTURE_CODE = "270000"
	END_POINT       = ".json"
)

func formatWeather(body []byte) (*Weather, error) {
	weather := new(Weather)
	if err := json.Unmarshal(body, weather); err != nil {
		log.Println("JSON unmarshal error: ", err)
		return nil, err
	}
	return weather, nil
}

func (w *Weather) ToS() string {
	area := fmt.Sprintf("%sの天気です。\n", w.Area)
	head := fmt.Sprintf("%s\n", w.HeadLine)
	body := fmt.Sprintf("%s\n", w.Body)
	result := area + head + body

	return result
}

func GetWeather() (str string, err error) {
	prefecture_code := config.Config.PrefectureCode
	if prefecture_code == "" {
		prefecture_code = PREFECTURE_CODE
	}

	client := utils.New(URL, map[string]string{})
	body, err := client.DoRequest("GET", prefecture_code+END_POINT, map[string]string{}, nil)
	if err != nil {
		log.Println("client doRequest error: ", err)
		return "", err
	}

	weather, err := formatWeather(body)
	if err != nil {
		log.Println("weather formatting error: ", err)
		return "", err
	}

	result := weather.ToS()
	return result, nil
}

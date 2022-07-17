package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Weather struct {
	Area     string `json:"targetArea"`
	HeadLine string `json:"headlineText"`
	Body     string `json:"text"`
}

func GetWeather() (str string, err error) {
	body, err := httpGetBody("https://www.jma.go.jp/bosai/forecast/data/overview_forecast/" + os.Getenv("PREFECTURE_CODE") + ".json")
	if err != nil {
		return str, err
	}
	weather, err := formatWeather(body)
	if err != nil {
		return str, err
	}

	result := weather.ToS()

	return result, nil
}

func httpGetBody(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("Get Http Error: %s", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("IO Read Error:: %s", err)
		return nil, err
	}

	return body, nil
}

func formatWeather(body []byte) (*Weather, error) {
	weather := new(Weather)
	if err := json.Unmarshal(body, weather); err != nil {
		err = fmt.Errorf("JSON Unmarshal error: %s", err)
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

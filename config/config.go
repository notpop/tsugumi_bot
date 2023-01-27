package config

import (
	// "github.com/go-ini/ini"
	// "log"
	"os"
)

type ConfigList struct {
	ApiKey               string
	ChannelSecret        string
	ChannelToken         string
	SystemLog            string
	SystemLogFromPackage string
	Port                 string
	PrefectureCode       string
}

var Config ConfigList

func init() {
	// cfg, err := ini.Load("config.ini")
	// if err != nil {
	// 	log.Printf("faild to read file: %v", err)

	// 	log.Println("target changed. retry to read file.")
	// 	cfg, err = ini.Load("../config.ini")
	// 	if err != nil {
	// 		log.Printf("faild to read file again: %v", err)
	// 		os.Exit(1)
	// 	}
	// }

	Config = ConfigList{
		// ApiKey:        cfg.Section("openai").Key("api_key").String(),
		// ChannelSecret: cfg.Section("line").Key("channel_secret").String(),
		// ChannelToken:  cfg.Section("line").Key("channel_token").String(),
		ApiKey:        os.Getenv("OPEN_AI_API_KEY"),
		ChannelSecret: os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		ChannelToken:  os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
		// SystemLog:        cfg.Section("system").Key("log_file").String(),
		// Port:             ":" + cfg.Section("web").Key("port").String(),
		SystemLog:      os.Getenv("SYSTEM_LOG"),
		Port:           ":" + os.Getenv("PORT"),
		PrefectureCode: os.Getenv("PREFECTURE_CODE"),
	}

	Config.SystemLogFromPackage = "../" + Config.SystemLog
}

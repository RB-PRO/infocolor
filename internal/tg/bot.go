package tg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	Token  string `json:"token"`
	Folder string `json:"folder"`
	Bot    *tgbotapi.BotAPI
}

// Получение значение из файла
func NewTelegram(filename string) (bot *Telegram, Err error) {

	// Прочитать файл
	fileBytes, Err := os.ReadFile(filename)
	if Err != nil {
		return nil, fmt.Errorf("os.ReadFile: %v", Err)
	}

	// Распарсить
	Err = json.Unmarshal(fileBytes, &bot)
	if Err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", Err)
	}

	bot.Bot, Err = tgbotapi.NewBotAPI(bot.Token)
	if Err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", Err)
	}

	bot.Bot.Debug = false

	log.Printf("Авторизация в аккаунте %s", bot.Bot.Self.UserName)

	return bot, nil
}

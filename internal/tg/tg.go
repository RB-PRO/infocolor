package tg

import (
	"fmt"
	"log"
	"slices"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	tg, ErrTg := NewTelegram("tg.json")
	if ErrTg != nil {
		log.Fatalf("json.Unmarshal: %v\n", ErrTg)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tg.Bot.GetUpdatesChan(u)

	type SearchStruct struct {
		Brand     string
		ColorCode string
	}

	UserSearch := make(map[int64]SearchStruct)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			// Construct a new message from the given chat ID and containing
			// the text that we received.
			// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "update.Message.Text")

			switch update.Message.Command() {
			case "start":
				msgBrand := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите бренд в меню")
				brands, _ := tg.FileList()
				KeyBrands := Menu(brands)
				msgBrand.ReplyMarkup = KeyBrands
				tg.Bot.Send(msgBrand)
				// if _, ok := UserSearch[update.Message.Chat.ID]; ok {
				delete(UserSearch, update.Message.Chat.ID)
				// }
			case "ping":
				msgPing := tgbotapi.NewMessage(update.Message.Chat.ID,
					fmt.Sprintf("pong\nMessage.Chat.ID - %d", update.Message.Chat.ID))
				tg.Bot.Send(msgPing)
			default:
				msgSearch := tgbotapi.NewMessage(update.Message.Chat.ID,
					"Вы выбрали: "+update.Message.Text)
				// msg.Text = "I don't know that command"
				if _, ok := UserSearch[update.Message.Chat.ID]; !ok {
					brands, _ := tg.FileList()
					// Если такой файл существует
					if slices.Contains(brands, update.Message.Text) {
						UserSearch[update.Message.Chat.ID] = SearchStruct{
							Brand: update.Message.Text,
						}
						msgSearch.Text = fmt.Sprintf(`Вы выбрали бренд "%s". Теперь нужно выбрать код краски:`, update.Message.Text)
						// ExtracttCode
						// ExtractCode
						dataFormuls, ErrLoadFile := tg.LoadFile(update.Message.Text)
						if ErrLoadFile != nil {
							msgSearch.Text = fmt.Sprintf(`Не смог открыть данные по бренду %s. /start`, update.Message.Text)
							break
						}
						PaintCodes := ExtractCode(dataFormuls)
						KeyPaintCodes := Menu(append([]string{"Всё"}, PaintCodes...))
						msgSearch.ReplyMarkup = KeyPaintCodes
					} else {
						msgSearch.Text = fmt.Sprintf(`Бренда "%s" нет в нашей базе. Выберите другой бренд: /start`, update.Message.Text)
					}
				} else { // Выбор краски
					ID := update.Message.Chat.ID
					search := UserSearch[ID]
					search.ColorCode = update.Message.Text
					fmt.Println("Пошли искать", update.Message.Chat.UserName, UserSearch[ID])
					if search.ColorCode == "Всё" {
						search.ColorCode = ""
					}
					UserSearch[ID] = search

					dataFormuls, ErrLoadFile := tg.LoadFile(UserSearch[ID].Brand)
					if ErrLoadFile != nil {
						msgSearch.Text = fmt.Sprintf(`Не смог открыть данные по бренду %s. /start`, update.Message.Text)
						break
					}

					DataCodeFormuls := ExtractFormulaFromCode(dataFormuls, UserSearch[ID].ColorCode)

					fmt.Println("Выводим", len(dataFormuls), len(DataCodeFormuls))
					for _, DataCodeFormul := range DataCodeFormuls {
						msgFormul := tgbotapi.NewMessage(update.Message.Chat.ID,
							PrintFormul(DataCodeFormul))

						time.Sleep(time.Millisecond * 100)
						tg.Bot.Send(msgFormul)
					}
					break
				}
				tg.Bot.Send(msgSearch)
			}

			// // Send the message.
			// if _, err := tg.Bot.Send(msg); err != nil {
			// 	panic(err)
			// }
		}
	}
}

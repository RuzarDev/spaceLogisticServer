package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Println("Bot token not found")
	}
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/webhook", webhookHandler)
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}

}
func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var update tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверяем, пришло ли сообщение из WebApp
	if update.Message != nil && update.Message.WebAppData != nil && update.Message.WebAppData.Data != "" {
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(update.Message.WebAppData.Data), &data); err != nil {
			log.Println("Ошибка при чтении данных из WebApp:", err)
			return
		}

		log.Printf("📦 Получены данные из WebApp: %+v\n", data)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "✅ Данные успешно получены!")
		if _, err := bot.Send(msg); err != nil {
			log.Println("Ошибка при отправке ответа:", err)
		}
	}

	w.WriteHeader(http.StatusOK)
}

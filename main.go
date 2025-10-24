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
	log.Printf("Received update: %+v\n", update)
	w.WriteHeader(http.StatusOK)
}

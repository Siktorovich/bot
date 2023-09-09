package main

import (
	"log"
	"os"

	"github.com/Siktorovich/bot/internal/app/commands"
	"github.com/Siktorovich/bot/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		commander.HandleUpdate(update)
	}
}

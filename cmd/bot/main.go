package main

import (
	"github.com/ryodanqqe/telegramHello/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	token := os.Getenv("TOKEN")

	productService := product.NewService()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
				continue
			case "list":
				listCommand(bot, update.Message, productService)
				continue
			default:
				defultBehavior(bot, update.Message)
				continue
			}
		}

	}
}

func helpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	bot.Send(tgbotapi.NewMessage(message.Chat.ID,
		"/help - help \n"+"/list - list products"))
}

func listCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, src *product.Service) {
	output := ""

	products := src.List()
	for _, product := range products {
		output += product.Title + "\n"
	}

	bot.Send(tgbotapi.NewMessage(message.Chat.ID, output))
}

func defultBehavior(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Unknown command"))
}

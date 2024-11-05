package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ryodanqqe/telegramHello/internal/service/product"
)

type Commander struct {
	bot        *tgbotapi.BotAPI
	productSrc *product.Service
	commands   map[string]func(message *tgbotapi.Message)
}

func NewCommander(bot *tgbotapi.BotAPI, productSrc *product.Service) *Commander {
	c := &Commander{
		bot:        bot,
		productSrc: productSrc,
		commands:   make(map[string]func(message *tgbotapi.Message)),
	}
	c.registerCommand()
	return c
}

func (c *Commander) registerCommand() {
	c.commands["/help"] = c.Help
	c.commands["/list"] = c.List
	c.commands["/default"] = c.Default
}

func (c *Commander) HandleMessage(m *tgbotapi.Message) {
	command := m.Command()
	handler, exists := c.commands[command]
	if !exists {
		c.Default(m)
	}

	handler(m)
}

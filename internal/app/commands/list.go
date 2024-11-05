package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(message *tgbotapi.Message) {
	output := ""

	products := c.productSrc.List()
	for _, product := range products {
		output += product.Title + "\n"
	}

	c.bot.Send(tgbotapi.NewMessage(message.Chat.ID, output))
}

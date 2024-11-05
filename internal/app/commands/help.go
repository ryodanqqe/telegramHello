package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(message *tgbotapi.Message) {
	c.bot.Send(tgbotapi.NewMessage(message.Chat.ID,
		"/help - help \n"+"/list - list products"))
}

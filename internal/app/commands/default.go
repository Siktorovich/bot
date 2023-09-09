package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	c.bot.Send(msg)
}

package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := "Your products: \n\n"
	products := c.productService.List()

	for _, p := range products {
		outputMsg += p.Title
		outputMsg += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	c.bot.Send(msg)
}

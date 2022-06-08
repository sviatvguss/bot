package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	var msgText string
	product, err := c.productService.Get(idx)
	if err != nil {
		msgText = fmt.Sprintf("Sorry, failed to get product with index %d: %v", idx+1, err)
	} else {
		msgText = product.Title
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgText,
	)

	c.bot.Send(msg)
}

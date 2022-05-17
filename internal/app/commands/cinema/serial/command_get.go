package serial

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *CinemaSerialCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("CinemaSerialCommander.Get: wrong args (%s) - %v", args, err)
		return
	}

	serial, err := c.serialService.Describe(id)
	if err != nil {
		log.Printf("CinemaSerialCommander.Get: fail to get serial with id %d - %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		serial.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CinemaSerialCommander.Get: error sending reply message to chat - %v", err)
	}
}

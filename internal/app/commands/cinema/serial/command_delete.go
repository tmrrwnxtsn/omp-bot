package serial

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *CinemaSerialCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("CinemaSerialCommander.Delete: wrong args (%s) - %v", args, err)
		return
	}

	_, err = c.serialService.Remove(id)
	if err != nil {
		log.Printf("CinemaSerialCommander.Delete: fail to delete serial with id %d - %v", id, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("serial with id %d successfully deleted", id),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CinemaSerialCommander.Delete: error sending reply message to chat - %v", err)
	}
}

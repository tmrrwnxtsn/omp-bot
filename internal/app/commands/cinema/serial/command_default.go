package serial

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *CinemaSerialCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("CinemaSerialCommander.Default: [%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CinemaSerialCommander.Default: error sending reply message to chat - %v", err)
	}
}

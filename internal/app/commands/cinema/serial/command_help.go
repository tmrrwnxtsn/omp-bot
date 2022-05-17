package serial

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *CinemaSerialCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__cinema__serial — print list of commands\n"+
			"/get__cinema__serial serialID — get a serial\n"+
			"/list__cinema__serial cursor limit — get a list of serials\n"+
			"/delete__cinema__serial serialID — delete an existing serial\n"+
			"/new__cinema__serial title genre seasons_number — create a new serial\n"+
			"/edit__cinema__serial serialID field-new_value — edit a serial",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CinemaSerialCommander.Help: error sending reply message to chat - %v", err)
	}
}

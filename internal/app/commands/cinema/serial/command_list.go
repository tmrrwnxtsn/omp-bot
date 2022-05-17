package serial

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

var serialsOnPage = 3

func (c *CinemaSerialCommander) List(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	cursorLimit := strings.Split(args, " ")

	cursor, err := strconv.ParseUint(cursorLimit[0], 10, 64)
	if err != nil {
		log.Println("wrong args", cursorLimit[0])
		return
	}
	if cursor == 0 {
		cursor = 1
	}

	limit, err := strconv.ParseUint(cursorLimit[1], 10, 64)
	if err != nil {
		log.Println("wrong args", cursorLimit[1])
		return
	}

	serials, err := c.serialService.List(cursor, limit)
	if err != nil {
		log.Println("error occurred getting serial list: ", err.Error())
		return
	}

	outputMsgText := "Here the serials: \n\n"

	for i := 0; i < serialsOnPage && i < len(serials); i++ {
		outputMsgText += serials[i].String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if len(serials) > serialsOnPage {
		callbackListData := &CallbackListData{
			Start: int(cursor),
			Curr:  serialsOnPage,
			End:   int(limit),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackListData.serialize()),
			),
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CinemaSerialCommander.List: error sending reply message to chat - %v", err)
	}
}

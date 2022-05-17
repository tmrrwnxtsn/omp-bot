package serial

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

var serialsOnPage = 3

func (c *CinemaSerialCommander) List(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	cursor, limit, err := parseListArguments(args)
	if err != nil {
		log.Printf("CinemaSerialCommander.List: fail to parse cursor and limit - %v", err)
		return
	}

	serials, err := c.serialService.List(cursor, limit)
	if err != nil {
		log.Printf("CinemaSerialCommander.List: error occurred getting serial list - %v", err)
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

func parseListArguments(args string) (uint64, uint64, error) {
	cursorLimit := strings.Split(args, " ")
	if len(cursorLimit) != 2 {
		return 0, 0, fmt.Errorf("wrong count of the arguments")
	}

	cursor, err := strconv.ParseUint(cursorLimit[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("wrong cursor %v", err)
	}
	if cursor == 0 {
		cursor = 1
	}

	limit, err := strconv.ParseUint(cursorLimit[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("wrong limit %v", err)
	}

	return cursor, limit, nil
}

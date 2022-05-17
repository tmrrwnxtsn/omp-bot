package serial

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type CallbackListData struct {
	Start int `json:"start"`
	Curr  int `json:"offset"`
	End   int `json:"limit"`
}

func (d *CallbackListData) serialize() string {
	serializedData, _ := json.Marshal(d)

	callbackPath := path.CallbackPath{
		Domain:       "cinema",
		Subdomain:    "serial",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return callbackPath.String()
}

func (c *CinemaSerialCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("CinemaSerialCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	serials, err := c.serialService.List(uint64(parsedData.Start), uint64(parsedData.End))
	if err != nil {
		log.Printf("CinemaSerialCommander.CallbackList: " +
			"error error occurred getting serial list - %v" + err.Error())
		return
	}

	outputMsgText := "Here the serials: \n\n"

	for i := parsedData.Curr; i < parsedData.Curr+serialsOnPage && i < len(serials); i++ {
		outputMsgText += serials[i].String()
		outputMsgText += "\n"
	}

	inlineKeyboardRow := make([]tgbotapi.InlineKeyboardButton, 0, 2)

	if parsedData.Curr > 0 {
		callbackListData := &CallbackListData{
			Start: parsedData.Start,
			Curr:  parsedData.Curr - serialsOnPage,
			End:   parsedData.End,
		}

		inlineKeyboardRow = append(inlineKeyboardRow, tgbotapi.NewInlineKeyboardButtonData(
			"Prev page", callbackListData.serialize(),
		))
	}

	if parsedData.Curr+serialsOnPage < len(serials) {
		callbackListData := &CallbackListData{
			Start: parsedData.Start,
			Curr:  parsedData.Curr + serialsOnPage,
			End:   parsedData.End,
		}

		inlineKeyboardRow = append(inlineKeyboardRow, tgbotapi.NewInlineKeyboardButtonData(
			"Next page", callbackListData.serialize(),
		))
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(inlineKeyboardRow...),
	)

	editConf := tgbotapi.EditMessageTextConfig{
		Text: outputMsgText,
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:      callback.Message.Chat.ID,
			MessageID:   callback.Message.MessageID,
			ReplyMarkup: &inlineKeyboard,
		},
	}

	_, err = c.bot.Send(editConf)
	if err != nil {
		log.Printf("CinemaSerialCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}

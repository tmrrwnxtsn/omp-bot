package serial

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"log"
	"strconv"
	"strings"
)

func (c *CinemaSerialCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idStr := strings.Split(args, " ")[0]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("CinemaSerialCommander.Edit: wrong args (%s) - %v", args, err)
		return
	}

	serial, err := c.serialService.Describe(id)
	if err != nil {
		log.Printf("CinemaSerialCommander.Edit: fail to get serial with id %d - %v", id, err)
		return
	}

	if err = parseEditArguments(args, serial); err != nil {
		log.Printf("CinemaSerialCommander.Edit: wrong args - %v", err)
		return
	}

	if err = c.serialService.Update(serial.ID, *serial); err != nil {
		log.Printf("CinemaSerialCommander.Edit: fail to edit serial with id %d - %v", serial.ID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("serial with id %d successfully edited", serial.ID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CinemaSerialCommander.Edit: error sending reply message to chat - %v", err)
	}
}

func parseEditArguments(args string, serial *cinema.Serial) error {
	serialInfo := strings.Split(args, " ")
	if len(serialInfo) < 1 {
		return fmt.Errorf("not enough arguments")
	}

	for _, arg := range serialInfo[1:] {
		if err := parseEditArgument(arg, serial); err != nil {
			return fmt.Errorf("wrong arg %v", err)
		}
	}

	return nil
}

func parseEditArgument(arg string, serial *cinema.Serial) error {
	serialInfo := strings.Split(arg, "-")
	if len(serialInfo) < 2 {
		return fmt.Errorf("wrong format")
	}

	switch serialInfo[0] {
	case "title":
		if serialInfo[1] == "" {
			return fmt.Errorf("title is empty")
		}
		serial.Title = serialInfo[1]
	case "genre":
		if serialInfo[1] == "" {
			return fmt.Errorf("genre is empty")
		}
		serial.Genre = serialInfo[1]
	case "seasons_number":
		seasonsNum, err := strconv.Atoi(serialInfo[1])
		if err != nil {
			return fmt.Errorf("seasons number: %v", err)
		}
		serial.SeasonsNum = seasonsNum
	default:
		return fmt.Errorf("unknown: %s", serialInfo[0])
	}
	return nil
}

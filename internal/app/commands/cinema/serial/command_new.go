package serial

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
	"log"
	"strconv"
	"strings"
)

func (c *CinemaSerialCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	serial, err := parseNewArguments(args)
	if err != nil {
		log.Printf("CinemaSerialCommander.New: fail to parse serial arguments - %v", err)
		return
	}

	serialID, err := c.serialService.Create(serial)
	if err != nil {
		log.Printf("CinemaSerialCommander.New: fail to create serial - %v", err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("serial successfully created with id %d", serialID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CinemaSerialCommander.New: error sending reply message to chat - %v", err)
	}
}

func parseNewArguments(args string) (cinema.Serial, error) {
	serialInfo := strings.Split(args, " ")
	if len(serialInfo) != 3 {
		return cinema.Serial{}, fmt.Errorf("not enougth arguments")
	}

	title := serialInfo[0]
	if title == "" {
		return cinema.Serial{}, fmt.Errorf("wrong args: title is empty")
	}

	genre := serialInfo[1]
	if genre == "" {
		return cinema.Serial{}, fmt.Errorf("wrong args: genre is empty")
	}

	seasonsNum, err := strconv.Atoi(serialInfo[2])
	if err != nil {
		return cinema.Serial{}, fmt.Errorf("wrong seasons number: %v", err)
	}

	return cinema.Serial{
		Title:      title,
		Genre:      genre,
		SeasonsNum: seasonsNum,
	}, nil
}

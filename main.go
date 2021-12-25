package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/NicoNex/echotron/v3"
)

type bot struct {
	chatID int64
	echotron.API
}

//go:embed token
var token string

func newBot(chatID int64) echotron.Bot {
	return &bot{
		chatID,
		echotron.NewAPI(token),
	}
}

func (b *bot) handleMessage(update *echotron.Update) {
	var (
		christmas = time.Date(getYear(), time.December, 25, 0, 0, 0, 0, time.UTC)
		msg       = strings.ToLower(update.Message.Text)
	)

	if update.Message.From.LanguageCode == "it" {
		if strings.Contains(msg, "quanto manca a natale") {
			loc, _ := time.LoadLocation("Europe/Rome")
			until := GetUntil(christmas.In(loc))

			b.SendMessage(
				fmt.Sprintf(
					"%d giorn%s, %d or%s, %d minut%s e %d second%s.",
					until.Days,
					IfThenElse(until.Days != 1, "i", "o"),
					until.Hours,
					IfThenElse(until.Hours != 1, "e", "a"),
					until.Minutes,
					IfThenElse(until.Minutes != 1, "i", "o"),
					until.Seconds,
					IfThenElse(until.Seconds != 1, "i", "o"),
				),
				b.chatID,
				&echotron.MessageOptions{
					ReplyToMessageID: update.Message.ID,
				},
			)
		}
	} else {
		if strings.Contains(msg, "how long until christmas") {
			until := GetUntil(christmas)

			b.SendMessage(
				fmt.Sprintf(
					"%d day%s, %d hour%s, %d minute%s and %d second%s.",
					until.Days,
					IfThenElse(until.Days != 1, "s", ""),
					until.Hours,
					IfThenElse(until.Hours != 1, "s", ""),
					until.Minutes,
					IfThenElse(until.Minutes != 1, "s", ""),
					until.Seconds,
					IfThenElse(until.Seconds != 1, "s", ""),
				),
				b.chatID,
				&echotron.MessageOptions{
					ReplyToMessageID: update.Message.ID,
				},
			)
		}
	}
}

func (b *bot) Update(update *echotron.Update) {
	if update.Message != nil {
		b.handleMessage(update)
	}
}

func main() {
	dsp := echotron.NewDispatcher(token, newBot)
	log.Println(dsp.Poll())
}

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

var (
	christmas = time.Date(getYear(), time.December, 25, 0, 0, 0, 0, time.UTC)

	//go:embed token
	token string
)

func newBot(chatID int64) echotron.Bot {
	return &bot{
		chatID,
		echotron.NewAPI(token),
	}
}

func (b *bot) handleMessage(update *echotron.Update) {
	var msg = strings.ToLower(update.Message.Text)
	var until = time.Until(christmas)

	if update.Message.From.LanguageCode == "it" {
		if strings.Contains(msg, "quanto manca a natale") {
			b.SendMessage(
				fmt.Sprintf(
					"%.0f giorni, %.0f ore, %.0f minuti e %.0f secondi.",
					until.Hours()/24,
					getHours(until.Hours()),
					getSixties(until.Minutes()),
					getSixties(until.Seconds()),
				),
				b.chatID,
				&echotron.MessageOptions{
					ReplyToMessageID: update.Message.ID,
				},
			)
		}
	} else {
		if strings.Contains(msg, "how long until christmas") {
			b.SendMessage(
				fmt.Sprintf(
					"%.0f days, %.0f hours, %.0f minutes and %.0f seconds.",
					until.Hours()/24,
					getHours(until.Hours()),
					getSixties(until.Minutes()),
					getSixties(until.Seconds()),
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

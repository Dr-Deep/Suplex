package event

import (
	"fmt"

	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"
)

type EventReadyHandler struct {
	*internal.SuplexBot
}

func NewReady(bot *internal.SuplexBot) *EventReadyHandler {
	return &EventReadyHandler{SuplexBot: bot}
}

func (bot *EventReadyHandler) Exec(s *discordgo.Session, ev *discordgo.Ready) {
	bot.Logger.Info(
		"READY",
		fmt.Sprintf("Username: '%s' ID: '%s'", ev.User.Username, ev.User.ID),
	)

	// Register Slash Commands
	//? aus ireiner liste
}

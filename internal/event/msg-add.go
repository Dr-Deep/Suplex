package event

import (
	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"
)

type EventMessageCreate struct {
	*internal.SuplexBot
}

func NewMessageCreate(bot *internal.SuplexBot) *EventMessageCreate {
	return &EventMessageCreate{SuplexBot: bot}
}

func (bot *EventMessageCreate) Exec(s *discordgo.Session, ev *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if ev.Author.ID == s.State.User.ID {
		return
	}

	if ev.Author.Bot {
		return
	}

	// store in DB für llm

	// uwuify?
}

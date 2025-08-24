package event

import (
	"fmt"

	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"
)

type EventDebugLogHandler struct {
	*internal.SuplexBot
}

func NewDebugLog(bot *internal.SuplexBot) *EventDebugLogHandler {
	return &EventDebugLogHandler{SuplexBot: bot}
}

func (bot *EventDebugLogHandler) Exec(s *discordgo.Session, ev *discordgo.Event) {
	bot.Logger.Debug(
		ev.Type,
		fmt.Sprintf("%#v", ev.Struct),
	)
}

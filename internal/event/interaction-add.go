/*
 * Slash Command Handler
 */

package event

import (
	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"
)

type EventInteractionCreateHandler struct {
	*internal.SuplexBot
}

func NewInteractionCreate(bot *internal.SuplexBot) *EventInteractionCreateHandler {
	return &EventInteractionCreateHandler{SuplexBot: bot}
}

func (bot *EventInteractionCreateHandler) Exec(s *discordgo.Session, ev *discordgo.InteractionCreate) {
	//? slash cmd handler
}

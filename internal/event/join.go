package event

import (
	"suplex/internal"

	"github.com/bwmarrin/discordgo"
)

type JoinHandler struct {
	self *internal.Suplex
}

func NewJoinHandler(self *internal.Suplex) *JoinHandler {
	return &JoinHandler{self: self}
}

func (h *JoinHandler) Exec(s *discordgo.Session, e *discordgo.GuildMemberAdd) {}

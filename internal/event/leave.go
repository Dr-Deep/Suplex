package event

import (
	"suplex/internal"

	"github.com/bwmarrin/discordgo"
)

type LeaveHandler struct {
	self *internal.Suplex
}

func NewLeaveHandler(self *internal.Suplex) *LeaveHandler {
	return &LeaveHandler{self: self}
}

func (h *LeaveHandler) Exec(s *discordgo.Session, e *discordgo.GuildMemberRemove) {}

package event

import (
	"suplex/internal"

	"github.com/bwmarrin/discordgo"
)

type ReadyHandler struct {
	self *internal.Suplex
}

func NewReadyHandler(self *internal.Suplex) *ReadyHandler {
	return &ReadyHandler{self: self}
}

func (h *ReadyHandler) Exec(s *discordgo.Session, e *discordgo.Ready) {
	h.self.Logger.Info("Ready", "Bot up as", e.User.Username)
}

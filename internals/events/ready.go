package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tripledoomer/Suplex/internals/logging"
)

type ReadyHandler struct{}

func NewReadyHandler() *ReadyHandler {
	return &ReadyHandler{}
}

func (h *ReadyHandler) Handler(s *discordgo.Session, e *discordgo.Ready) {
	logging.Logf(logging.Info, "[events/ready] %s is running", e.User.String())
}

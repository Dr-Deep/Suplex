package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type ReadyHandler struct{}

func NewReadyHandler() *ReadyHandler {
	return &ReadyHandler{}
}

func (h *ReadyHandler) Handler(s *discordgo.Session, e *discordgo.Ready) {
	log.Printf("[events/ready] %s is running\n", e.User.String())
}

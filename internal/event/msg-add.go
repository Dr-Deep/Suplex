package event

import (
	"strings"
	"suplex/internal"

	"github.com/bwmarrin/discordgo"
)

type MessageAddHandler struct {
	*internal.Suplex
}

func NewMessageAddHandler(self *internal.Suplex) *MessageAddHandler {
	return &MessageAddHandler{self}
}

func (h *MessageAddHandler) Exec(s *discordgo.Session, e *discordgo.MessageCreate) {

	//
	if e.Author.Bot {
		return
	}

	// Command
	if strings.HasPrefix(e.Message.Content, h.Config.Prefix) {
		h.Handler.Handle(e.Message)
	}
}

package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tripledoomer/Suplex/internals/logging"
)

type MessageHandler struct{}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if channel, err := s.Channel(e.ChannelID); err != nil {
		logging.Logf(logging.Warn, "[events/message] ERR: %s", err.Error())
		return
	} else {
		//  									 chan author msg
		logging.Logf(logging.Info, "[events/message] %s %s %s", channel.Name, e.Author.String(), e.Content)
	}

}

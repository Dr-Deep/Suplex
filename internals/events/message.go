package events

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type MessageHandler struct{}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) Handler(s *discordgo.Session, e *discordgo.MessageCreate) {
	if channel, err := s.Channel(e.ChannelID); err != nil {
		log.Printf("[events/message] ERR: %s", err.Error())
		return
	} else {
		//  						 chan author msg
		log.Printf("[events/message] '%s' '%s' '%s'", channel.Name, e.Author.String(), e.Content)
	}

}

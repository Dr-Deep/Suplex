package events

import (
	"fmt"
	"log"

	"github.com/tripledoomer/Suplex/internals/config"

	"github.com/bwmarrin/discordgo"
)

type JoinLeaveHandler struct{}

func NewJoinLeaveHandler() *JoinLeaveHandler {
	return &JoinLeaveHandler{}
}

func (h *JoinLeaveHandler) HandlerJoin(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	log.Printf("[events/joinleave] '%s' joined\n", e.Member.User.String())
	if _, err := s.ChannelMessageSend("934176121161920702", fmt.Sprintf("Willkommen <@%s>", e.Member.User.ID)); err != nil {
		log.Printf("[events/joinleave] ERR: %s", err.Error())
	} else {
		if err := s.GuildMemberRoleAdd(config.Cfg.Guild.ID, e.Member.User.ID, config.Cfg.Guild.Roles.DefaultRole); err != nil {
			log.Printf("[events/joinleave] ERR: %s", err.Error())
		}
	}

}

func (h *JoinLeaveHandler) HandlerLeave(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	log.Printf("[events/joinleave] '%s' left\n", e.Member.User.String())
}

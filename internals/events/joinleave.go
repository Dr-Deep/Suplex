package events

import (
	"fmt"

	"github.com/tripledoomer/Suplex/internals/config"
	"github.com/tripledoomer/Suplex/internals/logging"

	"github.com/bwmarrin/discordgo"
)

type JoinLeaveHandler struct{}

func NewJoinLeaveHandler() *JoinLeaveHandler {
	return &JoinLeaveHandler{}
}

func (h *JoinLeaveHandler) HandlerJoin(s *discordgo.Session, e *discordgo.GuildMemberAdd) {
	logging.Logf(logging.Info, "[events/joinleave] '%s' joined", e.Member.User.String())
	if _, err := s.ChannelMessageSend("934176121161920702", fmt.Sprintf("Willkommen <@%s>", e.Member.User.ID)); err != nil {
		logging.Logf(logging.Warn, "[events/joinleave] ERR: %s", err.Error())
	} else {
		if err := s.GuildMemberRoleAdd(config.Cfg.Guild.ID, e.Member.User.ID, config.Cfg.Guild.Roles.DefaultRole); err != nil {
			logging.Logf(logging.Warn, "[events/joinleave] ERR: %s", err.Error())
		}
	}

}

func (h *JoinLeaveHandler) HandlerLeave(s *discordgo.Session, e *discordgo.GuildMemberRemove) {
	logging.Logf(logging.Info, "[events/joinleave] '%s' left", e.Member.User.String())
}

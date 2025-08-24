package event

import (
	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"

	_ "github.com/zekroTJA/shinpuru/pkg/discordoauth/v2"
)

type AutojoinHandler struct {
	*internal.SuplexBot
}

func NewAutojoin_Handler(bot *internal.SuplexBot) *AutojoinHandler {
	return &AutojoinHandler{SuplexBot: bot}
}

func (bot *AutojoinHandler) Exec_GuildMemberAdd(s *discordgo.Session, ev *discordgo.GuildMemberAdd) {
	/*
	 * per DM, discord oauth2 link für API perms
	 * dann member rolle
	 */

	//
}

func (bot *AutojoinHandler) Exec_GuildMemberRemove(s *discordgo.Session, ev *discordgo.GuildMemberRemove) {
	/*
	 * gucken in DB nach OAuth2 daten für userID + refresh token?
	 */
}

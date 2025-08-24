package event

import (
	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"
)

// welcome member per banner

type WelcomeUserBanner struct {
	*internal.SuplexBot
}

func NewWelcomeUserBannerHandler(bot *internal.SuplexBot) *WelcomeUserBanner {
	return &WelcomeUserBanner{SuplexBot: bot}
}

func (bot *WelcomeUserBanner) Exec(s *discordgo.Session, ev *discordgo.GuildMemberAdd) {
	// banner??
}

package command

import (
	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"
)

type SetupGuildConfig struct {
	*internal.SuplexBot
}

func NewSetupGuildConfigCommand(bot *internal.SuplexBot) *internal.Command {
	var cmd = &SetupGuildConfig{SuplexBot: bot}

	return &internal.Command{
		Exec: cmd.Exec,
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "setup-guild",
			Description: "Setup the guild configuration",
			//? Type              ApplicationCommandType
			//? Options           []*ApplicationCommandOption
			//? IntegrationTypes *[]ApplicationIntegrationType
		},
	}
}
func (cmd *SetupGuildConfig) Exec(s *discordgo.Session, ev *discordgo.Interaction) {
	// TODO cfg abfragen und embed mit config details senden und in DB
}

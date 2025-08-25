package internal

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

/*
createdCommands, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, *GuildID, commands)

if *RemoveCommands {
		for _, cmd := range createdCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, cmd.ID)
			if err != nil {
				log.Fatalf("Cannot delete %q command: %v", cmd.Name, err)
			}
		}
	}
*/

type CommandHandler struct {
	*SuplexBot
	cmdMap sync.Map
}

func NewCommandHandler(bot *SuplexBot) *CommandHandler {
	return &CommandHandler{
		SuplexBot: bot,
	}
}

// Register a *Command
func (bot *CommandHandler) Register(cmd *Command) {
	// check if it exiists already
	_, exists := bot.cmdMap.Load(cmd.ApplicationCommand.Name)
	if exists {
		bot.Logger.Info("command exists", cmd.ApplicationCommand.Name)
		return
	}

	// Store Command
	bot.cmdMap.Store(
		cmd.ApplicationCommand.Name,
		cmd,
	)
}

func (bot *CommandHandler) List() []*Command {
	var cmds []*Command
	bot.cmdMap.Range(
		func(_key, _value any) bool {
			_cmd, oke := _value.(*Command)
			if !oke {
				return true
			}
			cmds = append(cmds, _cmd)

			return true
		},
	)

	return cmds
}

/*
// Iterieren
m.Range(func(key, value any) bool {
    fmt.Printf("Command: %v -> %v\n", key, value)
    return true // weiter iterieren
})
*/

// Command Handler
func (bot *CommandHandler) Exec(s *discordgo.Session, ev *discordgo.InteractionCreate) {
	_cmd, exists := bot.cmdMap.Load(ev.ApplicationCommandData().Name)
	if !exists {
		bot.Logger.Error("Command does not exists", ev.ApplicationCommandData().Name)
		return
	}

	// Run Command
	cmd, oke := _cmd.(*Command)
	if !oke {
		bot.Logger.Error("Command Type Error", ev.ApplicationCommandData().Name)
		return
	}

	go cmd.Exec(bot.Session, ev.Interaction)
}

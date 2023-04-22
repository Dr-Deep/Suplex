package internal

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	cmdMap map[string]*Command
	sync.Mutex
}

func (h *CommandHandler) RegisterCommand(cmd *Command) error {
	if _, ok := h.cmdMap[cmd.Invoke]; ok {
		return fmt.Errorf("command exists")
	}

	h.Lock()
	h.cmdMap[cmd.Invoke] = cmd
	h.Unlock()

	return nil
}

// Todo: middlewares
func (h *CommandHandler) Handle(message *discordgo.Message) {

	var split = strings.Split(
		strings.Replace(message.Content, "%", "", 1),
		" ",
	)

	cmd, ok := h.cmdMap[split[0]]
	if !ok {
		return
	}

	go cmd.Exec(
		&CommandContext{
			Arguments: split,
			Message:   message,
		},
	)
}

/*
func (bot *Suplex) RegisterCommands() {
	time.Sleep(time.Second * 3)

	for _, cmd := range bot.Handler.cmdMap {
		fmt.Printf("%#v\n", cmd.AppCmd)

		command := &discordgo.ApplicationCommand{
			Name:        "basic-command",
			Description: "Basic command",
		}

		if _, err := bot.Session.ApplicationCommandCreate(
			bot.Session.State.User.ID,
			"",
			command,
		); err != nil {
			bot.Logger.Error("(*Suplex).RegisterCommands()", cmd.AppCmd.Name, "couldnt register Command", err.Error())
			continue
		}
		fmt.Printf("OK: %s\n", cmd.AppCmd.Name)
	}
}
*/

package internal

import "github.com/bwmarrin/discordgo"

type Command struct {
	Invoke      string
	Description string
	//perms
	Exec func(*CommandContext)
}

type CommandContext struct {
	Arguments []string
	*discordgo.Message
}

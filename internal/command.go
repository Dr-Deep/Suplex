package internal

import "github.com/bwmarrin/discordgo"

type Command struct {
	//? middleware
	*discordgo.ApplicationCommand
	Exec func(*discordgo.Session, *discordgo.Interaction)
}

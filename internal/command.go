package internal

import "github.com/bwmarrin/discordgo"

type Command struct {
	//? middleware
	Exec func(*discordgo.Session, *discordgo.Interaction)
	*discordgo.ApplicationCommand
}

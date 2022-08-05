package cogs

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tripledoomer/Suplex/internals/config"
)

type MwPermissions struct{}

func (mw *MwPermissions) Exec(ctx *Context, cmd Command) (next bool, err error) {
	if !cmd.ModRequired() {
		// wenn der cmd kein admin braucht:
		next = true
		return
	}

	defer func() {
		if !next && err == nil {
			ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Du bist nicht dazu berechtigt dieses Kommando aufzurufen")
		}
	}()

	guild, err := ctx.Session.Guild(ctx.Message.GuildID)
	if err != nil {
		return
	}

	if guild.OwnerID == ctx.Message.Author.ID {
		next = true
		return
	}

	roleMap := make(map[string]*discordgo.Role)
	for _, role := range guild.Roles {
		roleMap[role.ID] = role
	}

	for _, rID := range ctx.Message.Member.Roles {
		if rID == config.Cfg.Guild.Roles.ModeratorRole {
			next = true
			return
		}
		if role, ok := roleMap[rID]; ok && role.Permissions&discordgo.PermissionAdministrator > 0 {
			next = true
			break
		}
	}

	return
}

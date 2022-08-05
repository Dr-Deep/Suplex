package commands

import (
	"github.com/tripledoomer/Suplex/internals/cogs"
)

type CmdPing struct{}

func (c *CmdPing) Invokes() []string {
	return []string{"ping", "p"}
}

func (c *CmdPing) Description() string {
	return "Ping command"
}

func (c *CmdPing) ModRequired() bool {
	return false
}

func (c *CmdPing) Exec(ctx *cogs.Context) error {
	if _, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Pong!"); err != nil {
		return err
	}
	return nil
}

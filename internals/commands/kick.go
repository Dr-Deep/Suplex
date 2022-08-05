package commands

import (
	"github.com/tripledoomer/Suplex/internals/cogs"
)

type CmdKick struct{}

func (c *CmdKick) Invokes() []string {
	return []string{"kick", "k"}
}

func (c *CmdKick) Description() string {
	return "kick a member"
}

func (c *CmdKick) ModRequired() bool {
	return true
}

func (c *CmdKick) Exec(ctx *cogs.Context) error {
	// kick
	return nil
}

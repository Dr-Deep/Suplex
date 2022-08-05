package commands

import (
	"github.com/tripledoomer/Suplex/internals/cogs"
)

type CmdSay struct{}

func (c *CmdSay) Invokes() []string {
	return []string{"ping", "p"}
}

func (c *CmdSay) Description() string {
	return "say command"
}

func (c *CmdSay) ModRequired() bool {
	return false
}

func (c *CmdSay) Exec(ctx *cogs.Context) error {
	// say
	return nil
}

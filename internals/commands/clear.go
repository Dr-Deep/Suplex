package commands

import (
	"github.com/tripledoomer/Suplex/internals/cogs"
)

type CmdClear struct{}

func (c *CmdClear) Invokes() []string {
	return []string{"clear", "c"}
}

func (c *CmdClear) Description() string {
	return "delete a amount of messages in a channel"
}

func (c *CmdClear) ModRequired() bool {
	return true
}

func (c *CmdClear) Exec(ctx *cogs.Context) error {
	// clear
	return nil
}

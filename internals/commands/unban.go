package commands

import (
	"github.com/tripledoomer/Suplex/internals/cogs"
)

type CmdUnban struct{}

func (c *CmdUnban) Invokes() []string {
	return []string{"unban", "ub"}
}

func (c *CmdUnban) Description() string {
	return "unban a member"
}

func (c *CmdUnban) ModRequired() bool {
	return true
}

func (c *CmdUnban) Exec(ctx *cogs.Context) error {
	// unban
	return nil
}

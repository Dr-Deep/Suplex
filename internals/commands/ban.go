package commands

import (
	"github.com/tripledoomer/Suplex/internals/cogs"
)

type CmdBan struct{}

func (c *CmdBan) Invokes() []string {
	return []string{"ban", "b"}
}

func (c *CmdBan) Description() string {
	return "ban a member"
}

func (c *CmdBan) ModRequired() bool {
	return true
}

func (c *CmdBan) Exec(ctx *cogs.Context) error {
	// ban
	return nil
}

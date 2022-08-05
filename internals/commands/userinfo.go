package commands

import (
	"github.com/tripledoomer/Suplex/internals/cogs"
)

type CmdUserinfo struct{}

func (c *CmdUserinfo) Invokes() []string {
	return []string{"userinfo", "user", "u"}
}

func (c *CmdUserinfo) Description() string {
	return "gives some information about a member"
}

func (c *CmdUserinfo) ModRequired() bool {
	return false
}

func (c *CmdUserinfo) Exec(ctx *cogs.Context) error {
	// userinfo
	return nil
}

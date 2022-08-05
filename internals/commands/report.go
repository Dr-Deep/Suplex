package commands

import (
	"github.com/tripledoomer/Suplex/internals/cogs"
)

type CmdReport struct{}

func (c *CmdReport) Invokes() []string {
	return []string{"report", "r"}
}

func (c *CmdReport) Description() string {
	return "Report a member"
}

func (c *CmdReport) ModRequired() bool {
	return false
}

func (c *CmdReport) Exec(ctx *cogs.Context) error {
	// report
	return nil
}

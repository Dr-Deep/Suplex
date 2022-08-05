package commands

import (
	"github.com/tripledoomer/Suplex/internals/cogs"
)

type CmdHelp struct{}

func (c *CmdHelp) Invokes() []string {
	return []string{"help", "h"}
}

func (c *CmdHelp) Description() string {
	return "display help menu"
}

func (c *CmdHelp) ModRequired() bool {
	return false
}

func (c *CmdHelp) Exec(ctx *cogs.Context) error {
	// help list
	return nil
}

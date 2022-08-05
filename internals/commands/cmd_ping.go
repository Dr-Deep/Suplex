package commands

type CmdPing struct{}

func (c *CmdPing) Invokes() []string {
	return []string{"ping", "p"}
}

func (c *CmdPing) Description() string {
	return "Ping command"
}

func (c *CmdPing) AdminRequired() bool {
	return true
}

func (c *CmdPing) Exec(ctx *Context) error {
	if _, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Pong!"); err != nil {
		return err
	}
	return nil
}

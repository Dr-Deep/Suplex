package command

import (
	"suplex/internal"
)

type TestCommand struct {
	*internal.Suplex
}

func NewTestCommand(self *internal.Suplex) *internal.Command {
	var cmd = &TestCommand{self}
	return &internal.Command{
		Invoke:      "test",
		Description: "test command",
		Exec:        cmd.Exec,
	}
}

func (c *TestCommand) Exec(ctx *internal.CommandContext) {
	if _, err := c.Session.ChannelMessageSend(
		ctx.ChannelID,
		"test command",
	); err != nil {
		c.Logger.Error("(*TestCommand).Exec()", err.Error())
	}
}

/*
emb := &discordgo.MessageEmbed{
		Title: "Help Center",
		Description: "If you generally need help with the usage or setup of shinpuru, take a look into the " +
			"[**Wiki**](https://github.com/zekroTJA/shinpuru/wiki). There you can find a lot of useful resources " +
			"around shinpuru's features.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Command Help",
				Value: cmdHelp,
			},
			{
				Name: "Contact",
				Value: "If you need help with shinpuru, feel free to ask on my [dev discord](https://dc.zekro.de). " +
					"You can also contact me directly, either via email (`contact@zekro.de`) or via " +
					"[twitter](https://twitter.com/zekrotja).",
			},
		},
	}

	err = ctx.RespondEmbed(emb)
*/

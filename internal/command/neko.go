package command

/*
const (
	NekoURL = "http://saturn.0x4ba.de:50666/neko"
)

type NekoCommand struct {
	*internal.Suplex
}

func NewNekoCommand(self *internal.Suplex) *internal.Command {
	var (
		cmd    = &NekoCommand{self}
		isNSFW bool
	)

	return &internal.Command{
		AppCmd: discordgo.ApplicationCommand{
			Name:        "neko",
			Description: "",
			NSFW:        &isNSFW,
		},
		Exec: cmd.Exec,
	}
}

func (c *NekoCommand) Exec(ctx *discordgo.InteractionCreate) {
	//embed -> channel
}

func (c *NekoCommand) GetNekoPicture() ([]byte, error) {
	var body []byte
	resp, err := http.Get(NekoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
*/

package command

import (
	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"
)

/*
* als Suplex mit "typing..."
```
2. Multi-User-Session-Management
Um ein LLM so zu benutzen, dass es verschiedene User „getrennt“ behandeln kann, braucht man Sitzungen:
Pro Discord-User ein eigener Kontext (History, Persona, Memory).
Speicherung in DB
Bot erkennt den User über Message.Author.ID und holt den zugehörigen Kontext.
So können mehrere Leute gleichzeitig mit dem gleichen LLM über denselben Bot reden,
ohne dass die Kontexte sich vermischen.

In Thread oder command
```
https://github.com/go-deepseek/deepseek


/ask-llm: nachricht/frage + kontext der letzden nachrichten
*/

type AskLLMCommand struct {
	*internal.SuplexBot
}

func NewAskLLMCommand(bot *internal.SuplexBot) *internal.Command {
	var cmd = &AskLLMCommand{SuplexBot: bot}

	return &internal.Command{
		Exec: cmd.Exec,
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "ask-llm",
			Description: "Ask the LLM a question",
			//? Type              ApplicationCommandType
			//? Options           []*ApplicationCommandOption
			//? IntegrationTypes *[]ApplicationIntegrationType
		},
	}
}

func (cmd *AskLLMCommand) Exec(s *discordgo.Session, ev *discordgo.Interaction) {
	// TODO dafür muss DB stehen
}

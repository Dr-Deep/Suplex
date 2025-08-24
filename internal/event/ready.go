package event

import (
	"fmt"

	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"
)

type EventReadyHandler struct {
	*internal.SuplexBot
}

func NewReady(bot *internal.SuplexBot) *EventReadyHandler {
	return &EventReadyHandler{SuplexBot: bot}
}

func (bot *EventReadyHandler) Exec(s *discordgo.Session, ev *discordgo.Ready) {
	bot.Logger.Info(
		"READY",
		fmt.Sprintf("Username: '%s' ID: '%s'", ev.User.Username, ev.User.ID),
	)

	// Register Slash Commands
	//? aus ireiner liste
}

/*
func onReady(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("Eingeloggt als %s#%s\n", r.User.Username, r.User.Discriminator)

	// Status-Loop starten
	go func() {
		for {
			_ = s.UpdateStreamingStatus(0, "in Entwicklung...", "https://www.twitch.tv/deepcrack101")
			time.Sleep(7 * time.Second)

			ping := fmt.Sprintf("Latenz: %dms", s.HeartbeatLatency().Milliseconds())
			_ = s.UpdateStreamingStatus(0, ping, "https://www.twitch.tv/deepcrack101")
			time.Sleep(7 * time.Second)
		}
	}()
}
*/

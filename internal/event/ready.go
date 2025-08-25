package event

import (
	"fmt"
	"time"

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

	// Status-Loop
	go func() {
		for {
			_ = s.UpdateGameStatus(0, "mit seinem Schwanz")
			time.Sleep(7 * time.Second)

			ping := fmt.Sprintf("Latenz: %dms", s.HeartbeatLatency().Milliseconds())
			_ = s.UpdateStreamingStatus(0, ping, "https://hbsdsrv.1337.cx/")
			time.Sleep(7 * time.Second)
		}
	}()
}

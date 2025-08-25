/*
  - AutoVC

manages voice channels
*/
package event

import (
	"github.com/Dr-Deep/Suplex.git/internal"
)

type AutoVC_Handler struct {
	*internal.SuplexBot
}

/* //?
* VoiceServerUpdate
* VoiceStateUpdate
 */

func NewAutoVCHandler(bot *internal.SuplexBot) *AutoVC_Handler {
	return &AutoVC_Handler{SuplexBot: bot}
}

//func  (bot *AutoVC_Handler) Exec(s *discordgo.Session, ev *discordgo.)

/* discordgo.Session
   // Should voice connections reconnect on a session reconnect.
   ShouldReconnectVoiceOnSessionError bool


	// Stores a mapping of guild id's to VoiceConnections
	VoiceConnections map[string]*VoiceConnection
*/

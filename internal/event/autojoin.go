/*
 * Anti Guild Leave System
 */
package event

import (
	"github.com/Dr-Deep/Suplex.git/internal"
	web "github.com/Dr-Deep/Suplex.git/web/src"
	"github.com/bwmarrin/discordgo"
)

type AutojoinHandler struct {
	oauth2 *web.DiscordOAuth2Client
	*internal.SuplexBot
}

func NewAutojoin_Handler(bot *internal.SuplexBot) *AutojoinHandler {
	return &AutojoinHandler{SuplexBot: bot}
}

func (bot *AutojoinHandler) Exec_GuildMemberAdd(s *discordgo.Session, ev *discordgo.GuildMemberAdd) {
	// no bots
	if ev.User.Bot {
		return
	}

	// OAuth2 ?
	if bot.oauth2 == nil {
		bot.oauth2 = web.NewDiscordOAuth2(
			bot.Cfg.Discord_Settings.Token,
			bot.Cfg.Discord_Settings.Client_ID,
			bot.Cfg.Discord_Settings.Client_Secret,
			bot.Cfg.Discord_Settings.Guild_ID,
			"REDIR_URL",
			[]string{"identify", "guilds.join"},
			nil, // OnSuccess web.OnSuccessFunc
			nil, //  OnError web.OnErrorFunc
		)
	} else {
		// nur die OnSuccess web.OnSuccessFunc ;  OnError web.OnErrorFunc
	}

	// notify user per DM
	s.ChannelMessageSendEmbed(
		ev.User.ID,
		//! auf oauth2 verify link
		&discordgo.MessageEmbed{},
	)

	// on success: role
	// wenn nd prison

	/*
	 * per DM, discord oauth2 link für API perms
	 * dann member rolle
	 */

	//
}

func (bot *AutojoinHandler) Exec_GuildMemberRemove(s *discordgo.Session, ev *discordgo.GuildMemberRemove) {
	/*
	 * gucken in DB nach OAuth2 daten für userID + refresh token?
	 */
}

/*

	//! store *DiscordOAuth2_Resp_Token in DB
// Get UserID
	user, err := oauth2.GetUser(access_token.Access_Token)
	if err != nil {
		//?
	}
	//! store *discordgo.User in DB

	// Add User to Guild
	if err := oauth2.AddGuildMember(user.ID); err != nil {
		//?
	}
*/

/*
Having the user's access token allows your application to make certain requests to the API on their behalf, restricted to whatever scopes were requested. expires_in is how long, in seconds, until the returned access token expires, allowing you to anticipate the expiration and refresh the token. To refresh, make another POST request to the token URL with the following parameters:

    grant_type - must be set to refresh_token
    refresh_token - the user's refresh token

*/

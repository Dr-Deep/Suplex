/*
* Welcome Banner for GuildMemberAdd
https://github.com/natrixdev/discord_welcome_bot_banners/blob/main/banners/banners.md
*/
package event

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"

	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/bwmarrin/discordgo"
	"github.com/fogleman/gg"
)

const (
	// Dimensions
	bannerWidth  = 600
	bannerHeight = 200

	// Assets
	bannerBackgroundImagePath = "/background.png"
	bannerFontFacePoints      = 32
	bannerFontFacePath        = "/BebasNeue-Regular.ttf" //? replit font
)

type WelcomeUserBannerHandler struct {
	bannerImg  image.Image
	avatarSize int

	memberCount int
	*internal.SuplexBot
}

func NewWelcomeUserBannerHandler(bot *internal.SuplexBot) *WelcomeUserBannerHandler {
	var h = &WelcomeUserBannerHandler{
		avatarSize: 128,
		SuplexBot:  bot,
	}

	// Load Background Asset in Memory
	bgImg, err := gg.LoadImage(bot.Cfg.AssetsDir + bannerBackgroundImagePath)
	if err != nil {
		bot.Logger.Error("Background Asset Error", err.Error())
	}

	// Load Banner Image in Memory
	bannerImg := gg.NewContext(bannerWidth, bannerHeight)
	bannerImg.DrawImage(bgImg, 0, 0)

	bannerImg.DrawCircle(75, bannerHeight/2, float64(h.avatarSize)/2)
	bannerImg.Clip()

	// Load FontFace Asset in Memory
	if err := bannerImg.LoadFontFace(bot.Cfg.AssetsDir+bannerFontFacePath, bannerFontFacePoints); err != nil {
		bot.Logger.Error("FontFace Asset Error", err.Error())
	}

	h.bannerImg = bannerImg.Image()

	return h
}

func (bot *WelcomeUserBannerHandler) Exec(s *discordgo.Session, ev *discordgo.GuildMemberAdd) {
	// no bots
	if ev.User.Bot {
		return
	}

	// Fetch Home-Guild MemberCount
	guild, err := s.Guild(bot.Cfg.Discord_Settings.Guild_ID)
	if err != nil {
		bot.Logger.Error()
		return
	}
	bot.memberCount = guild.MemberCount + 1

	// Banner
	var (
		filename = fmt.Sprintf("%s.png", ev.User.ID)
		buf      = bytes.NewBuffer([]byte{})
	)
	if err := bot.drawWelcomeUserBanner(ev.User, buf); err != nil {
		bot.Logger.Error("Draw Welcome Banner Error", err.Error())
		return
	}

	// Send Banner Embed
	_, err = s.ChannelMessageSendComplex(
		"WELCOME CHANNEL ID", //!
		&discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{{
				Image: &discordgo.MessageEmbedImage{
					URL: "attachment://" + filename,
				},
			}},
			Files: []*discordgo.File{{
				Name: filename,
				//ContentType: "", //? format eig png
				Reader: buf,
			}},
		},
	)
	if err != nil {
		bot.Logger.Error("Send Welcome Banner Embed Error", err.Error())
		return
	}
}

func (bot *WelcomeUserBannerHandler) drawWelcomeUserBanner(user *discordgo.User, buf io.ReadWriter) error {
	// Load Avatar
	resp, err := http.Get(user.AvatarURL("")) //? size
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode Avatar
	avatarImg, _, err := image.Decode(resp.Body)
	if err != nil {
		return err
	}

	// reuse memory ?
	var (
		bannerImg = gg.NewContextForImage(bot.bannerImg)
	)

	// Draw Avatar
	bannerImg.DrawImageAnchored(
		avatarImg,
		75, bannerHeight/2,
		0.5, 0.5,
	)
	bannerImg.ResetClip()

	// Draw Text
	bannerImg.SetRGB(1, 1, 1)
	bannerImg.DrawStringAnchored(
		fmt.Sprintf("Welcome %s", user.Username),
		320, 80,
		0.5, 0.5,
	)
	bannerImg.DrawStringAnchored(
		fmt.Sprintf(
			"You're the %d Member!",
			bot.memberCount,
		),
		320, 130,
		0.5, 0.5,
	)

	// Encode Banner Image as PNG
	png.Encode(buf, bannerImg.Image())

	return nil
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tripledoomer/Suplex/internals/cogs"
	"github.com/tripledoomer/Suplex/internals/commands"
	"github.com/tripledoomer/Suplex/internals/config"
	"github.com/tripledoomer/Suplex/internals/events"
	"github.com/tripledoomer/Suplex/internals/logging"

	"github.com/bwmarrin/discordgo"
)

const (
	configFile = "config.json"
)

func init() {
	startTime := time.Now()
	logging.Logf(logging.Info, "[init] %s", startTime.String())

	// load config
	config.ParseConfigFromJSONFile(configFile)

	// TODO: log in db/file?

	// log level
	if config.Cfg.Logging.Level <= 4 {
		logging.SetLogLevel(config.Cfg.Logging.Level)
	} else {
		logging.Logf(logging.Warn, "invalid log level")
	}

	// TODO: database

	logging.Logf(logging.Info, "[init] took %s", time.Since(startTime))
}

func main() {
	startTime := time.Now()
	s, Serr := discordgo.New(config.Cfg.Token)
	if Serr != nil {
		logging.Logf(logging.Fatal, Serr.Error())
	}

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	registerEvents(s)
	registerCommands(s, config.Cfg)

	if err := s.Open(); err != nil {
		logging.Logf(logging.Fatal, err.Error())
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	s.Close()
	logging.Logf(logging.Info, "[main] took %s", time.Since(startTime))

}

func registerEvents(s *discordgo.Session) {
	joinLeaveHandler := events.NewJoinLeaveHandler()
	s.AddHandler(joinLeaveHandler.HandlerJoin)
	s.AddHandler(joinLeaveHandler.HandlerLeave)

	s.AddHandler(events.NewMessageHandler().Handler)

	s.AddHandler(events.NewReadyHandler().Handler)

}

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	cmdHandler := cogs.NewCommandHandler(cfg.Prefix)
	cmdHandler.OnError = func(err error, ctx *cogs.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, fmt.Sprintf("[*/OnError] ERR: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterMiddleware(&cogs.MwPermissions{})
	s.AddHandler(cmdHandler.HandleMessage)
}

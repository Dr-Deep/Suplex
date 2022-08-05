package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tripledoomer/Suplex/internals/commands"
	"github.com/tripledoomer/Suplex/internals/config"
	"github.com/tripledoomer/Suplex/internals/events"

	"github.com/bwmarrin/discordgo"
)

const (
	configFile = "config.json"
)

func init() {
	startTime := time.Now()
	log.Println("[init] Starting...")

	cfgErr := config.ParseConfigFromJSONFile(configFile)
	if cfgErr != nil {
		log.Panic(cfgErr)
	}

	log.Printf("[init] took %s\n", time.Since(startTime))
}

func main() {
	startTime := time.Now()
	log.Println("[main] Suplex is Launching...")
	s, Serr := discordgo.New(config.Cfg.Token)
	if Serr != nil {
		log.Panic(Serr)
	}

	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	registerEvents(s)
	registerCommands(s, config.Cfg)

	if err := s.Open(); err != nil {
		log.Panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	s.Close()
	log.Printf("[main] took %s\n", time.Since(startTime))
	log.Println("Exiting...")

}

func registerEvents(s *discordgo.Session) {
	joinLeaveHandler := events.NewJoinLeaveHandler()
	s.AddHandler(joinLeaveHandler.HandlerJoin)
	s.AddHandler(joinLeaveHandler.HandlerLeave)

	s.AddHandler(events.NewMessageHandler().Handler)

	s.AddHandler(events.NewReadyHandler().Handler)

}

func registerCommands(s *discordgo.Session, cfg *config.Config) {
	cmdHandler := commands.NewCommandHandler(cfg.Prefix)
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, fmt.Sprintf("[*/OnError] ERR: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})
	s.AddHandler(cmdHandler.HandleMessage)
}

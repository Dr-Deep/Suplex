package main

import (
	"os"
	"os/signal"
	"suplex/internal"
	"suplex/internal/config"
	"suplex/internal/event"
	"suplex/internal/log"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const (
	configFilePath = "./config.json"
)

var (
	cfg    *config.Configuration
	logger *log.Logger
	//db
)

func init() {
	var err error

	// Config
	cfg, err = config.ParseConfigFromJSONFile(configFilePath)
	if err != nil {
		panic(err)
	}

	// Logger
	logger, err = log.NewLogger(
		cfg.Logging.Method,
		cfg.Logging.File,
		cfg.Logging.Level,
	)
	if err != nil {
		panic(err)
	}

	//?db

}

func main() {

	suplex, err := internal.NewSuplex(cfg, logger)
	if err != nil {
		panic(err)
	}

	//?
	suplex.Session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	/*
	 * Register Events
	 */
	suplex.Session.AddHandler(event.NewReadyHandler(suplex).Exec)
	suplex.Session.AddHandler(event.NewJoinHandler(suplex).Exec)
	suplex.Session.AddHandler(event.NewLeaveHandler(suplex).Exec)

	// Start Bot
	go suplex.Start()

	//
	var sc = make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	// Stop Bot
	suplex.Stop()
}

/*
 * Suplex - Gewalt | Macht | Ordnung
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/Dr-Deep/Suplex.git/internal/command"
	"github.com/Dr-Deep/Suplex.git/internal/config"
	"github.com/Dr-Deep/Suplex.git/internal/database"
	"github.com/Dr-Deep/Suplex.git/internal/event"
	"github.com/Dr-Deep/logging-go"
)

var (
	suplex *internal.SuplexBot
	logger *logging.Logger
	cfg    *config.Configuration
	db     *database.Database

	// Flags
	logFilePath = flag.String(
		"logfile",
		"",
		"file to redirect logs to",
	)
	logLevel = flag.String(
		"loglevel",
		"info",
		"log level ('debug', 'info', 'error', 'fatal', 'none')",
	)
	configFilePath = flag.String(
		"config",
		"./config.yml",
		"configuration file",
	)
	databaseFilePath = flag.String(
		"database",
		"./suplex.db",
		"database file",
	)
)

func initSuplex() *internal.SuplexBot {
	_suplex, err := internal.NewSuplexBot(
		logger,
		cfg,
		db,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	/*
	 * Register Event-Handlers
	 */
	eventHandlers := []any{
		_suplex.CommandHandler.Exec,
		event.NewReady(_suplex).Exec,
		event.NewWelcomeUserBannerHandler(_suplex).Exec,
		event.NewAutojoin_Handler(_suplex).Exec_GuildMemberAdd,
		event.NewAutojoin_Handler(_suplex).Exec_GuildMemberRemove,
	}
	for _, h := range eventHandlers {
		_suplex.Session.AddHandler(h)
	}

	/*
	 * Register Command-Handlers
	 */
	commandHandlers := []*internal.Command{
		command.NewHelpCommand(_suplex),
	}
	for _, h := range commandHandlers {
		_suplex.CommandHandler.Register(h)
	}

	return _suplex
}

func setup() {
	_logger, err := internal.InitLogger(*logFilePath, *logLevel)
	if err != nil {
		panic(err)
	}

	_cfg, err := internal.InitConfig(*configFilePath)
	if err != nil {
		panic(err)
	}

	_db, err := internal.InitDatabase(*databaseFilePath)
	if err != nil {
		panic(err)
	}

	logger = _logger
	cfg = _cfg
	db = _db
	suplex = initSuplex()
}

func main() {
	flag.Parse()
	setup()

	// Suplex Bot
	err := suplex.Launch()
	if err != nil {
		logger.Fatal("launch error", err.Error())
	}

	// Close resources we opened here
	if db != nil {
		if err := db.Close(); err != nil {
			logger.Error("Database closed with error", err.Error())
		}
	}

	if logger != nil {
		if err := logger.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Logger closed with error: %s\n", err.Error())
		}
	}

	os.Exit(0)
}

/*
 * Suplex - Gewalt | Macht | Ordnung
 */

package main

import (
	"flag"

	"github.com/Dr-Deep/Suplex.git/internal"
	"github.com/Dr-Deep/Suplex.git/internal/config"
	"github.com/Dr-Deep/Suplex.git/internal/database"
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
	_suplex := internal.NewSuplexBot(
		logger,
		cfg,
		db,
	)

	/*
	* Register Events
	 */
	_suplex.Session.AddHandler(
		_suplex.CommandHandler,
	)

	/*
	 * Register Commands
	 */

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
	defer logger.Close()

	// Suplex Bot
	err := suplex.Launch()
	if err != nil {
		logger.Fatal("launch error", err.Error())
	}
}

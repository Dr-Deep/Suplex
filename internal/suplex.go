package internal

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Dr-Deep/Suplex.git/internal/config"
	"github.com/Dr-Deep/Suplex.git/internal/database"
	"github.com/Dr-Deep/logging-go"
	"github.com/bwmarrin/discordgo"
)

type SuplexBot struct {
	sync.Mutex

	// cmd handler
	// event handler?
	// middleware?
	CommandHandler *CommandHandler

	//
	Session *discordgo.Session

	//
	Logger *logging.Logger
	Cfg    *config.Configuration
	DB     *database.Database

	// Signals
	interuptSignals chan os.Signal
	reloadSignals   chan os.Signal
}

func NewSuplexBot(logger *logging.Logger, cfg *config.Configuration, db *database.Database) (*SuplexBot, error) {
	// New Discord Session
	session, err := discordgo.New(cfg.Discord_Settings.Token)
	if err != nil {
		return nil, err
	}

	var bot = &SuplexBot{
		Session: session,
		Logger:  logger,
		Cfg:     cfg,
		DB:      db,

		interuptSignals: make(chan os.Signal, 1),
		reloadSignals:   make(chan os.Signal, 1),
	}
	bot.CommandHandler = NewCommandHandler(bot)

	return bot, nil
}

func (bot *SuplexBot) Launch() error {
	bot.Lock()

	bot.Logger.Info("launching...")

	// Signals
	{
		signal.Notify(
			bot.interuptSignals,
			syscall.SIGINT,
			syscall.SIGTERM,
		)
		/*
			signal.Notify(
				bot.reloadSignals,
				syscall.SIGUSR1,
				syscall.SIGUSR2,
			)
		*/
	}

	// Discord Session
	{
		// Extended Logging
		if bot.Logger.Level == logging.LogDebug {
			bot.Session.Debug = true
			bot.Session.LogLevel = discordgo.LogDebug
		}

		/*
			+ discord oauth2 user-agent
			// The user agent used for REST APIs
			UserAgent string
		*/

		// Should the session reconnect the websocket on errors.
		bot.Session.ShouldReconnectOnError = true

		// Whether or not to call event handlers synchronously.
		bot.Session.SyncEvents = false

		/*
			QUIC client in future?
			// The http client used for REST requests
			Client *http.Client
		*/

		/*
			type Session struct {
			    // Should the session retry requests when rate limited.
			    ShouldRetryOnRateLimit bool

			    // Should state tracking be enabled.
			    // State tracking is the best way for getting the users
			    // active guilds and the members of the guilds.
			    StateEnabled bool

			    // Managed state object, updated internally with events when
			    // StateEnabled is true.
			    State *State

			    // used to deal with rate limits
			    Ratelimiter *RateLimiter
			}
		*/

		bot.Session.Identify.Intents |= discordgo.IntentsAll

		// Open Websocket
		if err := bot.Session.Open(); err != nil {
			return err
		}
	}

	// add slash cmds ?

	bot.Unlock()
	bot.run()

	return nil
}

func (bot *SuplexBot) Reload() {
	bot.Lock()
	bot.Logger.Info("reloading...")

	//
	_logger, err := InitLogger(_logFilePath, _logLevel)
	if err != nil {
		panic(err)
	}

	_cfg, err := InitConfig(_configFilePath)
	if err != nil {
		panic(err)
	}

	_db, err := InitDatabase(_databaseFilePath)
	if err != nil {
	}

	//
	bot.Logger = _logger
	bot.Cfg = _cfg
	bot.DB = _db

	bot.Logger.Info("reloaded!")
	bot.Unlock()
}

func (bot *SuplexBot) Shutdown() {
	bot.Lock()
	bot.Logger.Info("shutdown...")

	// remove slash cmds

	// Discord Session
	if err := bot.Session.Close(); err != nil {
		bot.Logger.Error("Session closed with error", err.Error())
	}

	// Signals
	signal.Stop(bot.interuptSignals)
	close(bot.interuptSignals)
	close(bot.reloadSignals)

	// running things?
	//
}

func (bot *SuplexBot) run() {
	defer bot.handlePanic()

	for {
		select {
		case <-bot.interuptSignals:
			bot.Logger.Info("catched SIGINT/SIGTERM")
			bot.Shutdown()
			return

		case <-bot.reloadSignals:
			bot.Reload()
		}
	}

}

func (bot *SuplexBot) handlePanic() {
	if r := recover(); r != nil {
		bot.Logger.Error("PANIC", fmt.Sprintf("%#v", r))
		bot.Logger.Info("trying to restart...")
		bot.Shutdown()
	}
}

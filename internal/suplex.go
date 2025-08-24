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
	// cmd handler
	// event handler?
	// middleware?

	//
	Session *discordgo.Session

	//
	Logger *logging.Logger
	Cfg    *config.Configuration
	DB     *database.Database

	// Signals
	interuptSignals chan os.Signal
	reloadSignals   chan os.Signal

	sync.Mutex
}

func NewSuplexBot(logger *logging.Logger, cfg *config.Configuration, db *database.Database) *SuplexBot {
	var bot = &SuplexBot{
		//

		Logger: nil,
		Cfg:    nil,
		DB:     nil,

		interuptSignals: make(chan os.Signal, 1),
		reloadSignals:   make(chan os.Signal, 1),
	}

	return bot
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
		session, err := discordgo.New(bot.Cfg.Discord_Settings.Token)
		if err != nil {
			return err
		}

		/*
			//?? Intents
			session.Identify.Intents |= discordgo.IntentAutoModerationExecution
			session.Identify.Intents |= discordgo.IntentMessageContent
			dg.Identify.Intents = discordgo.IntentsGuildMessages


			// Intents setzen (sonst funktionieren Join-Events nicht)
			dg.Identify.Intents = discordgo.IntentsGuilds |
			discordgo.IntentsGuildMessages |
			discordgo.IntentsGuildMembers
		*/

		// Open Websocket
		if err := session.Open(); err != nil {
			return err
		}

		bot.Session = session
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
	bot.Session.Close()

	// Signals
	signal.Stop(bot.interuptSignals)
	close(bot.interuptSignals)
	close(bot.reloadSignals)

	// running things?
	//
}

func (bot *SuplexBot) run() {
	//defer handlepanic ?

	// log

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

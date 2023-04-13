package internal

import (
	"suplex/internal/config"
	"suplex/internal/log"

	"github.com/bwmarrin/discordgo"
)

type Suplex struct {
	// Session
	Session *discordgo.Session

	// Database
	//Database *database.DB

	// Config
	Config *config.Configuration

	// Logger
	Logger *log.Logger
}

func NewSuplex(cfg *config.Configuration, logger *log.Logger) (*Suplex, error) {
	var suplex = &Suplex{
		//Database: db,
		Config: cfg,
		Logger: logger,
	}

	session, err := discordgo.New(suplex.Config.Token)
	if err != nil {
		return nil, err
	}

	suplex.Session = session

	return suplex, err
}

func (bot *Suplex) Start() {
	if err := bot.Session.Open(); err != nil {
		panic(err)
	}
}

func (bot *Suplex) Stop() {}

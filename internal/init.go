package internal

import (
	"os"

	"github.com/Dr-Deep/Suplex.git/internal/config"
	"github.com/Dr-Deep/Suplex.git/internal/database"
	"github.com/Dr-Deep/logging-go"
)

const (
	rwForOwnerOnlyPerm = 0o600
)

var (
	_logFilePath      string
	_logLevel         string
	_configFilePath   string
	_databaseFilePath string
)

func InitLogger(logFilePath string, logLevel string) (*logging.Logger, error) {
	_logFilePath = logFilePath
	_logLevel = logLevel
	var logger *logging.Logger

	if logFilePath != "" {
		// #nosec G304 -- Zugriff nur auf bekannte Log- und Config-Dateien
		logFile, err := os.OpenFile(
			logFilePath,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			rwForOwnerOnlyPerm,
		)
		if err != nil {
			return nil, err
		}

		logger = logging.NewLogger(logFile)
	} else {
		logger = logging.NewLogger(os.Stdout)
	}

	// Logging Level
	switch logLevel {
	case "debug":
		logger.Level = logging.LogDebug

	case "info":
		logger.Level = logging.LogInfo

	case "error":
		logger.Level = logging.LogError

	case "fatal":
		logger.Level = logging.LogFatal

	case "none":
		logger.Level = logging.Level(0)
	}

	return logger, nil
}

func InitConfig(configFilePath string) (*config.Configuration, error) {
	_configFilePath = configFilePath
	var cfg *config.Configuration

	// #nosec G304 -- Zugriff nur auf bekannte Log- und Config-Dateien
	cfgFile, err := os.OpenFile(
		configFilePath,
		os.O_RDONLY,
		rwForOwnerOnlyPerm,
	)
	if err != nil {
		return nil, err
	}

	cfg, err = config.UnmarshalConfigFile(cfgFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func InitDatabase(databaseFilePath string) (*database.Database, error) {
	_databaseFilePath = databaseFilePath

	db, err := database.NewDatabase(databaseFilePath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

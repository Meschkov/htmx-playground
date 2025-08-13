package app

import (
	"log/slog"
	"os"
)

type App struct {
	Config *Config
}

// InitializeApp initializes the application with the provided configuration file path.
func InitializeApp(configFilePath string) (*App, error) {
	app := &App{}
	config, err := LoadConfig(configFilePath)
	if err != nil {
		return nil, err
	}
	app.Config = config

	var logLevel slog.Level
	switch config.Logger.Level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelDebug
	}

	var logger slog.Handler
	switch app.Config.Logger.Format {
	case "json":
		logger = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		})
	case "text":
		logger = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		})
	default:
		logger = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		})
	}

	slog.SetDefault(slog.New(logger))

	return &App{
		Config: config,
	}, nil
}

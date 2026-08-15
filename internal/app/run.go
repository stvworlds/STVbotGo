package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"STVbotGo/internal/config"
	"STVbotGo/internal/logger"

	"github.com/bwmarrin/discordgo"
)

func Run() error {
	env, err := config.Load()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	config.Logger(env)

	// config.Logger may set LOG_LEVEL=debug for dev when unset.
	logger.Init(logger.Config{
		MinLevel: logger.ParseLevel(os.Getenv("LOG_LEVEL")),
	})
	defer logger.Close()

	logger.Infof("starting bot environment=%s", env.Environment)

	sess, err := session(env.Token, env)
	if err != nil {
		return fmt.Errorf("create Discord session: %w", err)
	}

	if err = sess.Open(); err != nil {
		return fmt.Errorf("open Discord connection: %w", err)
	}
	defer func(sess *discordgo.Session) {
		if closeErr := sess.Close(); closeErr != nil {
			logger.WithError(closeErr, "close Discord session")
			return
		}
		logger.Info("Discord session closed")
	}(sess)

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopSignal

	logger.Info("shutdown signal received")
	return nil
}

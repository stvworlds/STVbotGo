package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"STVbotGo/internal/bot"
	"STVbotGo/internal/config"
	"STVbotGo/internal/logging"
)

func Run() error {
	logging.Configure()

	env, err := config.Load()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}

	sess, err := bot.NewSession(env.Token)
	if err != nil {
		return fmt.Errorf("create Discord session: %w", err)
	}

	if err = sess.Open(); err != nil {
		return fmt.Errorf("open Discord connection: %w", err)
	}
	defer sess.Close()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopSignal

	return nil
}

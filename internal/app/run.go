package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"STVbotGo/internal/config"

	"github.com/bwmarrin/discordgo"
)

func Run() error {
	env, err := config.Load()
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	config.Logger(env)

	sess, err := session(env.Token, env)
	if err != nil {
		return fmt.Errorf("create Discord session: %w", err)
	}

	if err = sess.Open(); err != nil {
		return fmt.Errorf("open Discord connection: %w", err)
	}
	defer func(sess *discordgo.Session) {
		err := sess.Close()
		if err != nil {
			fmt.Printf("Error with close Discord session: %s", err)
		}
		fmt.Printf("Closed Discord session: %s", err)
	}(sess)

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopSignal

	return nil
}

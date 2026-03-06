package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.SetPrefix("STVbotGo: ")
	log.Println("Starting...")

	sess, err := discordgo.New("Bot token")
	if err != nil {
		log.Fatalln("Error creating Discord session,", err)
	}

	sess.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is now running.  Press CTRL-C to exit.")
	})

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == sess.State.User.ID {
			return
		}

		if m.Content == "hello" {
			send, err := s.ChannelMessageSend(m.ChannelID, "Hello, World!")
			if err != nil {
				return
			}
			log.Println(send)
		}

	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatalln("Error opening connection,", err)
	}
	defer func(sess *discordgo.Session) {
		err := sess.Close()
		if err != nil {
			log.Fatalln("Error closing connection,", err)
		}
	}(sess)

	log.Println("Bot is now online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

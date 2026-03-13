package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	log.SetPrefix("STVbotGo: ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Starting...")

	log.Println("Loading .env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	log.Println("Loaded .env")
	token := os.Getenv("TOKEN")

	log.Println("Starting Discord...")
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("Error creating Discord session,", err)
		return
	}

	sess.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is now running. Press CTRL-C to exit.")
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

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

package handlers

import (
	"log"

	"STVbotGo/internal/config"
	"STVbotGo/internal/handlers/antiscam"
	"STVbotGo/internal/handlers/autothreadopen"
	"STVbotGo/internal/handlers/cats"
	"STVbotGo/internal/handlers/convert"
	"STVbotGo/internal/handlers/dog"
	"STVbotGo/internal/handlers/helloworld"
	"STVbotGo/internal/handlers/ipmessage"
	"STVbotGo/internal/handlers/joinmessage"
	"STVbotGo/internal/handlers/newscreate"
	"STVbotGo/internal/handlers/ping"
	"STVbotGo/internal/handlers/readymessage"

	"github.com/bwmarrin/discordgo"
)

func RegisterAll(s *discordgo.Session, cfg config.Env) {
	ping.Register(s)
	convert.Register(s)
	newscreate.Register(s, cfg)

	antiscam.Register(s, cfg)
	autothreadopen.Register(s, cfg)
	cats.Register(s, cfg)
	dog.Register(s, cfg)
	helloworld.Register(s, cfg)
	ipmessage.Register(s, cfg)
	joinmessage.Register(s, cfg)
	readymessage.Register(s, cfg)

	log.Println("Handlers are now loaded.")
}

package config

import (
	"log"
)

const prefix = "STVbotGo: "

func Logger(cfg Env) {
	log.SetPrefix(prefix)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

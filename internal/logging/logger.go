package logging

import "log"

const prefix = "STVbotGo: "

func Configure() {
	log.SetPrefix(prefix)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

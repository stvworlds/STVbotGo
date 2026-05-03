package utils

import "log"

const prefix = "STVbotGo: "

func Logger() {
	log.SetPrefix(prefix)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

package config

import (
	"log"
	"os"
	"strings"
)

const prefix = "STVbotGo: "

// Logger configures the standard library logger used by packages that still
// call log.Printf directly. Application-level logging should prefer internal/logger.
func Logger(cfg Env) {
	log.SetPrefix(prefix)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Mirror environment into LOG_LEVEL default for internal/logger when unset.
	if os.Getenv("LOG_LEVEL") == "" && strings.EqualFold(cfg.Environment, "dev") {
		_ = os.Setenv("LOG_LEVEL", "debug")
	}
}

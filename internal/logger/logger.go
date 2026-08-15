package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Level is a severity of a log record.
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func (l Level) emoji() string {
	switch l {
	case LevelDebug:
		return "🔍"
	case LevelInfo:
		return "ℹ️"
	case LevelWarn:
		return "⚠️"
	case LevelError:
		return "❌"
	default:
		return "•"
	}
}

// Config controls logger behaviour.
type Config struct {
	// WebhookURL is a Discord webhook for remote logs. Empty disables webhook output.
	WebhookURL string
	// Username is the Discord webhook display name.
	Username string
	// MinLevel is the minimum level written to console and webhook.
	MinLevel Level
	// DisableConsole skips stdout/stderr logging when true.
	DisableConsole bool
	// DisableCaller omits file:line from messages when true.
	DisableCaller bool
	// QueueSize is the async webhook buffer size (default 64).
	QueueSize int
}

type entry struct {
	level   Level
	message string
	caller  string
	at      time.Time
}

var (
	mu       sync.RWMutex
	cfg      = defaultConfig()
	queue    chan entry
	workerWG sync.WaitGroup
	started  bool
)

func defaultConfig() Config {
	return Config{
		Username:  "STVbot log",
		MinLevel:  LevelInfo,
		QueueSize: 64,
	}
}

// Init configures the package logger. Safe to call multiple times; re-init
// replaces config and (re)starts the webhook worker if needed.
func Init(c Config) {
	mu.Lock()
	defer mu.Unlock()

	if c.Username == "" {
		c.Username = "STVbot log"
	}
	if c.QueueSize <= 0 {
		c.QueueSize = 64
	}
	if c.WebhookURL == "" {
		c.WebhookURL = os.Getenv("LOG_WEBHOOK_URL")
	}

	cfg = c

	if cfg.WebhookURL == "" {
		return
	}

	if !started {
		queue = make(chan entry, cfg.QueueSize)
		workerWG.Add(1)
		go webhookWorker()
		started = true
	}
}

// Close flushes pending webhook logs and stops the worker.
// Call on application shutdown.
func Close() {
	mu.Lock()
	if !started {
		mu.Unlock()
		return
	}
	started = false
	q := queue
	mu.Unlock()

	close(q)
	workerWG.Wait()
}

// Debug logs a debug message.
func Debug(msg string) { logf(LevelDebug, "%s", msg) }

// Debugf logs a formatted debug message.
func Debugf(format string, args ...any) { logf(LevelDebug, format, args...) }

// Info logs an info message.
func Info(msg string) { logf(LevelInfo, "%s", msg) }

// Infof logs a formatted info message.
func Infof(format string, args ...any) { logf(LevelInfo, format, args...) }

// Warn logs a warning message.
func Warn(msg string) { logf(LevelWarn, "%s", msg) }

// Warnf logs a formatted warning message.
func Warnf(format string, args ...any) { logf(LevelWarn, format, args...) }

// Error logs an error message.
func Error(msg string) { logf(LevelError, "%s", msg) }

// Errorf logs a formatted error message.
func Errorf(format string, args ...any) { logf(LevelError, format, args...) }

// WithError logs err at error level with a message prefix.
func WithError(err error, msg string) {
	if err == nil {
		return
	}
	logf(LevelError, "%s: %v", msg, err)
}

func logf(level Level, format string, args ...any) {
	mu.RLock()
	c := cfg
	q := queue
	isStarted := started
	mu.RUnlock()

	if level < c.MinLevel {
		return
	}

	msg := fmt.Sprintf(format, args...)
	caller := ""
	if !c.DisableCaller {
		caller = callerInfo(3)
	}

	e := entry{
		level:   level,
		message: msg,
		caller:  caller,
		at:      time.Now(),
	}

	if !c.DisableConsole {
		writeConsole(e)
	}

	if c.WebhookURL == "" || !isStarted || q == nil {
		return
	}

	select {
	case q <- e:
	default:
		// Drop remote log rather than block hot paths; still keep console output.
		log.Printf("logger: webhook queue full, dropped %s log", level)
	}
}

func writeConsole(e entry) {
	parts := []string{e.level.String()}
	if e.caller != "" {
		parts = append(parts, e.caller)
	}
	parts = append(parts, e.message)
	line := strings.Join(parts, " | ")

	switch e.level {
	case LevelError, LevelWarn:
		_, _ = fmt.Fprintln(os.Stderr, line)
	default:
		_, _ = fmt.Fprintln(os.Stdout, line)
	}
}

func webhookWorker() {
	defer workerWG.Done()
	for e := range queue {
		sendWebhook(e)
	}
}

func sendWebhook(e entry) {
	mu.RLock()
	url := cfg.WebhookURL
	username := cfg.Username
	mu.RUnlock()

	if url == "" {
		return
	}

	body := formatWebhookBody(e)
	// Discord content hard limit is 2000 runes; keep a safety margin.
	body = truncate(body, 1900)

	payload, err := json.Marshal(map[string]string{
		"username": username,
		"content":  body,
	})
	if err != nil {
		log.Printf("logger: failed to encode webhook payload: %v", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		log.Printf("logger: failed to send webhook: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		log.Printf("logger: webhook status %d: %s", resp.StatusCode, strings.TrimSpace(string(respBody)))
	}
}

func formatWebhookBody(e entry) string {
	var b strings.Builder
	b.WriteString(e.level.emoji())
	b.WriteByte(' ')
	b.WriteString("**")
	b.WriteString(e.level.String())
	b.WriteString("**")
	b.WriteString(" · `")
	b.WriteString(e.at.Format(time.RFC3339))
	b.WriteString("`")
	if e.caller != "" {
		b.WriteString(" · `")
		b.WriteString(e.caller)
		b.WriteString("`")
	}
	b.WriteByte('\n')
	b.WriteString("```\n")
	b.WriteString(e.message)
	if !strings.HasSuffix(e.message, "\n") {
		b.WriteByte('\n')
	}
	b.WriteString("```")
	return b.String()
}

func callerInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

func truncate(s string, max int) string {
	if max <= 0 || len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}

// ParseLevel converts a human-readable level name into Level.
// Accepts: debug, info, warn/warning, error. Unknown values default to info.
func ParseLevel(s string) Level {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "debug":
		return LevelDebug
	case "warn", "warning":
		return LevelWarn
	case "error":
		return LevelError
	default:
		return LevelInfo
	}
}

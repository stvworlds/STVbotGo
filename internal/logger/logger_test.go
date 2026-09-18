package logger

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestParseLevel(t *testing.T) {
	cases := map[string]Level{
		"debug":   LevelDebug,
		"DEBUG":   LevelDebug,
		"info":    LevelInfo,
		"":        LevelInfo,
		"unknown": LevelInfo,
		"warn":    LevelWarn,
		"warning": LevelWarn,
		"error":   LevelError,
	}
	for in, want := range cases {
		if got := ParseLevel(in); got != want {
			t.Fatalf("ParseLevel(%q)=%v want %v", in, got, want)
		}
	}
}

func TestTruncate(t *testing.T) {
	if got := truncate("hello", 10); got != "hello" {
		t.Fatalf("unexpected: %q", got)
	}
	if got := truncate("hello world", 8); got != "hello..." {
		t.Fatalf("unexpected: %q", got)
	}
}

func TestFormatWebhookBody(t *testing.T) {
	body := formatWebhookBody(entry{
		level:   LevelError,
		message: "boom",
		caller:  "logger.go:1",
		at:      time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC),
	})
	for _, part := range []string{"ERROR", "boom", "logger.go:1", "```"} {
		if !strings.Contains(body, part) {
			t.Fatalf("body missing %q: %s", part, body)
		}
	}
}

func TestConsoleOnlyLog(t *testing.T) {
	Init(Config{
		MinLevel:       LevelDebug,
		DisableCaller:  true,
		DisableConsole: true,
	})
	// Must not panic without webhook worker.
	Debug("d")
	Info("i")
	Warn("w")
	Error("e")
	Debugf("%s", "d")
	Infof("%s", "i")
	Warnf("%s", "w")
	Errorf("%s", "e")
	WithError(assertErr{}, "failed")
	Close()
}

type assertErr struct{}

func (assertErr) Error() string { return "x" }

func TestInitSendsWebhookMessage(t *testing.T) {
	var received []string
	var mu sync.Mutex
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		mu.Lock()
		received = append(received, string(body))
		mu.Unlock()
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	Init(Config{
		WebhookURL:     ts.URL,
		MinLevel:       LevelInfo,
		DisableConsole: true,
	})
	Close()

	mu.Lock()
	defer mu.Unlock()
	if len(received) == 0 {
		t.Fatal("expected webhook message on Init, got none")
	}
	if !strings.Contains(received[0], "logger initialized") {
		t.Fatalf("expected message to contain 'logger initialized', got %s", received[0])
	}
}

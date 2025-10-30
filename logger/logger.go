package logger

import (
	"log/slog"
	"os"
	"sync"
)

var once sync.Once

// Init sets up a default slog logger. It’s safe to call multiple times.
func Init() {
	once.Do(func() {
		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelInfo,
			AddSource: false,
		})
		slog.SetDefault(slog.New(handler))
	})
}

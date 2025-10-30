package golang

import (
	"context"
	"log/slog"
	"time"

	"github.com/vinaycharlie01/godsapkg/execx"
)

// RunTests runs Go tests with given arguments
func RunTests(args ...string) error {
	slog.Info("🧪 Running Go Tests...")
	defaultArgs := []string{"test", "./..."}
	start := time.Now()
	if err := execx.Run(context.Background(), "go", true, append(defaultArgs, args...)...); err != nil {
		return err
	}
	slog.Info("✅ Tests passed", "duration", time.Since(start))
	return nil
}

// RunLint runs golangci-lint with given arguments
func RunLint(args ...string) error {
	slog.Info("🔍 Running Go Linter...")
	defaultArgs := []string{"run", "--timeout=5m"}
	start := time.Now()
	if err := execx.Run(context.Background(), "golangci-lint", true, append(defaultArgs, args...)...); err != nil {
		return err
	}
	slog.Info("✅ Lint passed", "duration", time.Since(start))
	return nil
}

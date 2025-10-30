package golang

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/vinaycharlie01/godsapkg/execx"
)

// RunTests runs Go tests with given arguments
func RunTests(args ...string) error {
	slog.Info("🧪 Running Go Tests...")
	defaultArgs := []string{"test", "./..."}
	start := time.Now()
	if err := execx.Run(context.Background(), "go", false, append(defaultArgs, args...)...); err != nil {
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
	if err := execx.Run(context.Background(), "golangci-lint", false, append(defaultArgs, args...)...); err != nil {
		return err
	}
	slog.Info("✅ Lint passed", "duration", time.Since(start))
	return nil
}

func RunInstall(pkgs []string, args ...string) error {
	if len(pkgs) == 0 {
		return fmt.Errorf("no package specified for installation")
	}

	slog.Info("📦 Installing Go packages individually...", "packages", pkgs)

	start := time.Now()
	for _, pkg := range pkgs {
		cmdArgs := append([]string{"install", pkg}, args...)
		if err := execx.Run(context.Background(), "go", false, cmdArgs...); err != nil {
			return fmt.Errorf("failed to install %s: %w", pkg, err)
		}
	}

	slog.Info("✅ Installation complete", "duration", time.Since(start))
	return nil
}

// RunModTasks runs `go mod tidy` and `go mod verify` sequentially
func RunModTasks() error {
	slog.Info("📦 Running Go module maintenance (tidy & verify)...")

	start := time.Now()

	commands := [][]string{
		{"mod", "tidy"},
		{"mod", "verify"},
	}

	for _, args := range commands {
		slog.Info("🔧 Executing", "command", fmt.Sprintf("go %s", strings.Join(args, " ")))
		if err := execx.Run(context.Background(), "go", false, args...); err != nil {
			return fmt.Errorf("failed to run 'go %s': %w", strings.Join(args, " "), err)
		}
	}
	slog.Info("✅ Module maintenance completed successfully", "duration", time.Since(start))
	return nil
}

// RunTests runs Go tests with given arguments
func Run() error {
	slog.Info("🧪 Running Go Mod Tidy...")
	defaultArgs := []string{"mod", "tidy"}
	start := time.Now()
	if err := execx.Run(context.Background(), "go", false, defaultArgs...); err != nil {
		return err
	}
	slog.Info("✅ Tests passed", "duration", time.Since(start))
	return nil
}

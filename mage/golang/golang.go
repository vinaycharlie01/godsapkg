package golang

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/vinaycharlie01/godsapkg/golang"
)

type Go mg.Namespace

// func init() {
// 	logger.Init()
// }

// Lint runs golangci-lint
func (Go) Lint(ctx context.Context) error {
	return fmt.Errorf("lint failed: %w", golang.RunLint())
}

// Test runs Go unit tests
func (Go) Test(ctx context.Context) error {
	start := time.Now()
	slog.Info("🧪 Running Go Tests...")
	if err := golang.RunTests("-v", "-cover"); err != nil {
		return fmt.Errorf("tests failed: %w", err)
	}
	slog.Info("✅ Tests passed", "duration", time.Since(start))
	return nil
}

// Clean removes build artifacts (optional)
func (Go) Clean() error {
	slog.Info("🧹 Cleaning up workspace...")
	return sh.RunV("go", "clean", "-testcache")
}

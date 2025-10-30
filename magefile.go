//go:build mage
// +build mage

package main

import (
	"context"
	"log/slog"

	"github.com/magefile/mage/mg"
	"github.com/vinaycharlie01/godsapkg/golang"
	"github.com/vinaycharlie01/godsapkg/logger"

	// mage:import
	_ "github.com/vinaycharlie01/godsapkg/mage/golang"
)

func init() {
	// ✅ Initialize slog globally here once
	logger.Init()
	slog.Info("🔧 Logger initialized for Mage tasks")
}

// Run all checks: lint + test
func CI(ctx context.Context) error {

	mg.SerialDeps(golang.RunLint, golang.RunTests)
	slog.Info("✅ CI pipeline completed successfully")
	return nil
}

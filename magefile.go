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
	// Initialize slog globally here once
	logger.Init()
	slog.Info("🔧 Logger initialized for Mage tasks")
}

func Setup(ctx context.Context) error {
	//Correct mg.Deps syntax — wrapped function properly
	pkgs := []string{
		"github.com/golangci/golangci-lint/cmd/golangci-lint@latest",
		"github.com/magefile/mage@latest",
	}

	mg.Deps(func() error {
		return golang.RunInstall(pkgs)
	})
	return nil
}

func CI(ctx context.Context) error {
	// Step 1: install required Go tools
	// Step 2: run tasks in sequence
	mg.SerialDeps(golang.RunModTasks, golang.RunLint, golang.RunTests)

	slog.Info("CI pipeline completed successfully")
	return nil
}

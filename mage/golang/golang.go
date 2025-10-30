package golang

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

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
	// slog.Info("🧪 Running Go Tests...")
	return fmt.Errorf("tests failed: %w", golang.RunTests("-v", "-cover"))

}

// Clean removes build artifacts (optional)
func (Go) Clean() error {
	slog.Info("🧹 Cleaning up workspace...")
	return sh.RunV("go", "clean", "-testcache")
}

func (Go) Installpkg(ctx context.Context, pkg string) error {

	if len(pkg) == 0 {
		return fmt.Errorf("no package")
		// pkg = []string{"./..."} // default: current module
	}
	pkgs := strings.Split(pkg, ",")

	if err := golang.RunInstall(pkgs); err != nil {
		return err
	}
	return nil
}

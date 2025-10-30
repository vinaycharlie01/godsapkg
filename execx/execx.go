package execx

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
)

// Run executes a command and streams its output.
// If streamToLog is true, output is sent to slog; otherwise, to terminal.
func Run(ctx context.Context, command string, streamToLog bool, args ...string) error {
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Stdin = os.Stdin

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command %q: %w", command, err)
	}

	if streamToLog {
		go streamToSlog(ctx, stdout, slog.LevelInfo)
		go streamToSlog(ctx, stderr, slog.LevelError)
	} else {
		go func() {
			_, _ = io.Copy(os.Stdout, stdout)
		}()
		go func() {
			_, _ = io.Copy(os.Stderr, stderr)
		}()
	}

	// Wait for the command to finish execution
	if err := cmd.Wait(); err != nil {
		// if context was canceled, wrap cleanly
		if ctx.Err() != nil {
			return fmt.Errorf("command %q canceled: %w", command, ctx.Err())
		}
		return fmt.Errorf("command %q failed: %w", command, err)
	}

	return nil
}

// streamToSlog reads command output and logs it to slog with the given level.
func streamToSlog(ctx context.Context, r io.Reader, level slog.Level) {
	scanner := bufio.NewScanner(r)
	const maxCapacity = 1024 * 1024 // 1 MB max line size
	buf := make([]byte, 64*1024)    // 64 KB initial buffer
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			slog.WarnContext(ctx, "stream canceled", "reason", ctx.Err())
			return
		default:
			slog.Log(ctx, level, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		slog.ErrorContext(ctx, "failed to read stream", "err", err)
	}
}

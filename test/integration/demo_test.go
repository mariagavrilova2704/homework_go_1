package integration_test

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

const commandTimeout = 10 * time.Second

func TestDemoBinaryOutput(t *testing.T) {
	t.Parallel()

	output := buildAndRunDemo(t)
	want := strings.Join([]string{
		"Привет, Мария! Добро пожаловать в Go.",
		"true",
		"300",
		"Пройдено 3 из 12 уроков",
		"true",
	}, "\n") + "\n"

	if output != want {
		t.Fatalf("demo output = %q, want %q", output, want)
	}
}

func TestDemoBinaryPrintsFiveLines(t *testing.T) {
	t.Parallel()

	output := strings.TrimSuffix(buildAndRunDemo(t), "\n")
	lines := strings.Split(output, "\n")

	if len(lines) != 5 {
		t.Fatalf("demo printed %d lines, want 5: %q", len(lines), output)
	}
}

func buildAndRunDemo(t *testing.T) string {
	t.Helper()

	binaryPath := buildDemoBinary(t)
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	t.Cleanup(cancel)

	cmd := exec.CommandContext(ctx, binaryPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("run demo binary: %v\noutput:\n%s", err, output)
	}

	return strings.ReplaceAll(string(output), "\r\n", "\n")
}

func buildDemoBinary(t *testing.T) string {
	t.Helper()

	repoRoot := repositoryRoot(t)
	binaryPath := filepath.Join(t.TempDir(), "demo")
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	t.Cleanup(cancel)

	cmd := exec.CommandContext(ctx, "go", "build", "-o", binaryPath, "./cmd/demo")
	cmd.Dir = repoRoot
	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("build demo binary: %v\noutput:\n%s", err, output)
	}

	return binaryPath
}

func repositoryRoot(t *testing.T) string {
	t.Helper()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot detect current test file path")
	}

	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
}

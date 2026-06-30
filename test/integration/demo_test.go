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
		"Курс: Go backend",
		"Урок 1: терминал, Git и первый Go-проект",
		"github.com/rinat-course/homework1",
		"go run ./cmd/demo",
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

func TestDemoBinaryCustomArgs(t *testing.T) {
	t.Parallel()

	output := buildAndRunDemo(t,
		"Алексей",
		"Основы Go",
		"package и fmt",
		"student",
		"homework1",
		"./cmd/demo",
	)
	want := strings.Join([]string{
		"Привет, Алексей! Добро пожаловать в Go.",
		"Курс: Основы Go",
		"Урок 1: package и fmt",
		"github.com/student/homework1",
		"go run ./cmd/demo",
	}, "\n") + "\n"

	if output != want {
		t.Fatalf("demo output with custom args = %q, want %q", output, want)
	}
}

func TestDemoBinaryRejectsTooManyArgs(t *testing.T) {
	t.Parallel()

	binaryPath := buildDemo(t)
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, binaryPath, "1", "2", "3", "4", "5", "6", "7")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("demo was expected to fail, output: %q", string(output))
	}

	text := string(output)
	if !strings.Contains(text, "too many arguments") {
		t.Fatalf("demo error output = %q, want 'too many arguments'", text)
	}
	if !strings.Contains(text, "Usage:") {
		t.Fatalf("demo error output = %q, want usage", text)
	}
}

func buildAndRunDemo(t *testing.T, args ...string) string {
	t.Helper()

	binaryPath := buildDemo(t)
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, binaryPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("run demo binary: %v, output: %s", err, output)
	}

	return string(output)
}

func buildDemo(t *testing.T) string {
	t.Helper()

	repoRoot := findRepoRoot(t)
	binaryPath := filepath.Join(t.TempDir(), binaryName())

	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "build", "-o", binaryPath, "./cmd/demo")
	cmd.Dir = repoRoot
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("build demo binary: %v, output: %s", err, output)
	}

	return binaryPath
}

func findRepoRoot(t *testing.T) string {
	t.Helper()

	current, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}

	for {
		if fileExists(filepath.Join(current, "go.mod")) {
			return current
		}

		parent := filepath.Dir(current)
		if parent == current {
			t.Fatal("go.mod was not found")
		}
		current = parent
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func binaryName() string {
	if runtime.GOOS == "windows" {
		return "main.exe"
	}

	return "main"
}

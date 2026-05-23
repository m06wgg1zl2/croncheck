package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// buildBinary compiles the binary into a temp file and returns its path.
// Skips the test if the build fails (e.g. in environments without go toolchain).
func buildBinary(t *testing.T) string {
	t.Helper()
	tmp := t.TempDir()
	out := tmp + "/croncheck"
	cmd := exec.Command("go", "build", "-o", out, ".")
	cmd.Dir = "."
	if err := cmd.Run(); err != nil {
		t.Skipf("could not build binary: %v", err)
	}
	return out
}

func TestCLI_ValidExpression(t *testing.T) {
	bin := buildBinary(t)
	cmd := exec.Command(bin, "-n", "3", "* * * * *")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := string(out)
	if !strings.Contains(output, "* * * * *") {
		t.Errorf("expected expression in output, got:\n%s", output)
	}
	if !strings.Contains(output, "UTC") {
		t.Errorf("expected UTC timezone in output, got:\n%s", output)
	}
}

func TestCLI_InvalidExpression(t *testing.T) {
	bin := buildBinary(t)
	cmd := exec.Command(bin, "not a cron")
	out, err := cmd.Output()
	if err == nil {
		t.Fatalf("expected non-zero exit, got output:\n%s", out)
	}
	if exitErr, ok := err.(*exec.ExitError); ok {
		if exitErr.ExitCode() != 1 {
			t.Errorf("expected exit code 1, got %d", exitErr.ExitCode())
		}
		if !strings.Contains(string(exitErr.Stderr), "invalid cron expression") {
			t.Errorf("expected error message in stderr, got: %s", string(exitErr.Stderr))
		}
	}
}

func TestCLI_NoArgs(t *testing.T) {
	bin := buildBinary(t)
	cmd := exec.Command(bin)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected non-zero exit when no args provided")
	}
}

func TestCLI_DescribeFlag(t *testing.T) {
	bin := buildBinary(t)
	cmd := exec.Command(bin, "-describe", "-n", "2", "0 9 * * 1-5")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := string(out)
	if len(output) == 0 {
		t.Error("expected non-empty output with -describe flag")
	}
}

func TestCLI_InvalidTimezone(t *testing.T) {
	bin := buildBinary(t)
	cmd := exec.Command(bin, "-tz", "Not/ATimezone", "* * * * *")
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected non-zero exit for invalid timezone")
	}
}

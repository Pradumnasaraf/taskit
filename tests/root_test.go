package test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestRootCmd tests the root command (taskit)
func TestRootCmd(t *testing.T) {

	expectedOutput := "A CLI tool to manage your tasks"
	cmd := exec.Command("taskit")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:32])
	assert.Equal(t, got, expectedOutput)
}

// TestRootCmdHelpFlag tests the root command (taskit) with the help flag
func TestRootCmdHelpFlag(t *testing.T) {

	expectedOutput := "A CLI tool to manage your tasks"
	cmd := exec.Command("taskit", "--help")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:32])
	assert.Equal(t, got, expectedOutput)
}

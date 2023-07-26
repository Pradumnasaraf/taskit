package test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestListCmd(t *testing.T) {

	expectedOutput := `╔══`
	cmd := exec.Command("taskit", "list")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output))
	assert.Matches(t, got, expectedOutput)

}

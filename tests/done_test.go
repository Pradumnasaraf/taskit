package test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestDoneCmd tests the done command
func TestDoneCmd(t *testing.T) {

	expectedOutput := "Task marked as done successfully"
	cmd := exec.Command("taskit", "done", "1")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output))
	assert.Equal(t, got, expectedOutput)

	// Validate the json output
	payload := ReadTasks()

	// Validate the json output 3rd task
	assert.Equal(t, payload[0]["Done"], true)

}
